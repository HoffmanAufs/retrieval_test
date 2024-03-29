package routetab

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gauss-project/aurorafs/pkg/boson"
	"github.com/gauss-project/aurorafs/pkg/logging"
	"github.com/gauss-project/aurorafs/pkg/p2p"
	"github.com/gogf/gf/os/gmlock"
)

var (
	PendingTimeout  = time.Second * 10
	pendingInterval = time.Second
)

type pendCallResItem struct {
	src        boson.Address
	createTime time.Time
	resCh      chan struct{}
}

type pendingCallResArray []pendCallResItem
type pendCallResTab struct {
	items       map[common.Hash]pendingCallResArray
	mu          *gmlock.Locker
	lockTimeout time.Duration
	logger      logging.Logger
	addr        boson.Address
	metrics     metrics
}

func newPendCallResTab(addr boson.Address, logger logging.Logger, met metrics) *pendCallResTab {
	return &pendCallResTab{
		items:       make(map[common.Hash]pendingCallResArray),
		mu:          gmlock.New(),
		lockTimeout: time.Second,
		logger:      logger,
		addr:        addr,
		metrics:     met,
	}
}

func (pend *pendCallResTab) Add(target, src boson.Address, resCh chan struct{}) error {
	key := target.String()
	pending := pendCallResItem{
		src:        src,
		createTime: time.Now(),
		resCh:      resCh,
	}
	mKey := common.BytesToHash(target.Bytes())

	pend.mu.Lock(key)
	defer pend.mu.Unlock(key)

	res, ok := pend.items[mKey]
	if ok {
		now := pendingCallResArray{}
		for _, v := range res {
			if v.src.Equal(src) {
				continue
			}
			now = append(now, v)
		}
		pend.items[mKey] = append(now, pending)
	} else {
		pend.items[mKey] = pendingCallResArray{pending}
	}
	return nil
}

func (pend *pendCallResTab) Forward(ctx context.Context, s *Service, target boson.Address, routes []RouteItem) error {
	key := target.String()

	pend.mu.Lock(key)
	defer pend.mu.Unlock(key)

	mKey := common.BytesToHash(target.Bytes())
	res := pend.items[mKey]
	for _, v := range res {
		if !v.src.Equal(pend.addr) {
			// forward
			s.doResp(ctx, p2p.Peer{Address: v.src}, target, routes)
		} else if v.resCh != nil {
			// sync return
			v.resCh <- struct{}{}
		}
	}
	delete(pend.items, mKey)
	return nil
}

func (pend *pendCallResTab) Gc(expire time.Duration) {
	for destKey, item := range pend.items {
		pend.mu.TryLockFunc(destKey.String(), func() {
			var now pendingCallResArray
			for i := len(item) - 1; i >= 0; i-- {
				if time.Since(item[i].createTime).Seconds() >= expire.Seconds() {
					if i == len(item)-1 {
						delete(pend.items, destKey)
					} else {
						now = item[i:]
					}
					break
				}
			}
			if len(now) > 0 {
				pend.items[destKey] = now
			}
		})
	}
}
