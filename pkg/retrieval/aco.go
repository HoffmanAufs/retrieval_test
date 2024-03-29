package retrieval

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gauss-project/aurorafs/pkg/boson"
)

const (
	defaultRate int64 = 2_000_000/8
)

type route struct {
	linkNode boson.Address
	targetNode boson.Address
}

func (r *route) ToString() string{
	return fmt.Sprintf("%v,%v", r.linkNode.String(), r.targetNode.String())
}

type DownloadDetail struct{
	startMs int64
	endMs 	int64
	size	int64
}

type routeMetric struct{
	downloadCount	int64
	downloadDetail	*DownloadDetail
}

type acoServer struct{
	routeMetric map[string]*routeMetric
	toZeroElapsed int64
	mutex	sync.Mutex
}

func newAcoServer() *acoServer{
	return &acoServer{
		routeMetric: make(map[string]*routeMetric),
		toZeroElapsed: 20*60,		// 1200s
		mutex: sync.Mutex{},
	}
}

func (s *acoServer) OnDownloadStart(route route){
	routeKey := route.ToString()

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, exist := s.routeMetric[routeKey]; exist{
		s.routeMetric[routeKey].downloadCount += 1
	}else{
		s.routeMetric[routeKey] = &routeMetric{
			downloadCount: 1,
			downloadDetail: &DownloadDetail{
				0,0,0,
			},
		}
	}
}

func (s *acoServer) onDownloadEnd(route route){
	routeKey := route.ToString()

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, exist := s.routeMetric[routeKey]; exist{
		if s.routeMetric[routeKey].downloadCount > 0{
			s.routeMetric[routeKey].downloadCount -= 1
		}
	}
}

func (s *acoServer) OnDownloadFinish(route route, downloadDetail *DownloadDetail){
	if downloadDetail == nil{
		s.onDownloadEnd(route)
		return
	}else{
		s.onDownloadTaskFinish(route, downloadDetail.startMs, downloadDetail.endMs, downloadDetail.size)
	}
}

func (s *acoServer) onDownloadTaskFinish(route route, startMs int64, endMs int64, size int64){
	routeKey := route.ToString()
	var retStartMs, retEndMs int64

	if startMs < endMs{
		retStartMs, retEndMs = startMs, endMs
	}else if startMs == endMs{
		retStartMs, retEndMs = startMs, endMs+10
	}else{
		retStartMs, retEndMs = endMs, startMs
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	// new route, just record msg
	if _, exist := s.routeMetric[routeKey]; !exist{
		s.routeMetric[routeKey] = &routeMetric{
			downloadCount: 0,
			downloadDetail: &DownloadDetail{
				startMs: retStartMs,
				endMs: retEndMs,
				size: size,
			},
		}
	// route exists, need update routeMetric
	}else{
		recordStartMs, recordEndMs := s.routeMetric[routeKey].downloadDetail.startMs, s.routeMetric[routeKey].downloadDetail.endMs

		if endMs < recordStartMs{
			return
		}else if recordEndMs < startMs{
			s.routeMetric[routeKey].downloadDetail = &DownloadDetail{
				startMs: startMs,
				endMs: endMs,
				size: size,
			}
		}else{
			if startMs < recordStartMs{
				s.routeMetric[routeKey].downloadDetail.startMs = startMs
			}
			if endMs > recordEndMs{
				s.routeMetric[routeKey].downloadDetail.endMs = endMs
			}
			s.routeMetric[routeKey].downloadDetail.size += size
		}
	}
}

func (s* acoServer) GetRouteAcoIndex(routeList []route) ([] int){
	routeCount := len(routeList)
	// get the score for each route
	routeScoreList := s.getSelectRouteListScore(routeList)

	// decide the order of the route 
	routeIndexList := make([]int, 0)

	totalScore := int64(0)
	for _, v := range routeScoreList{
		totalScore += v
	}

	rand.Seed(time.Now().Unix())
	selectRouteCount := 0
	for {
		selectRouteIndex, curScore, curSum := 0, int64(0), int64(0)

		randNum := (rand.Int63()%(totalScore))+1
		for k, v := range routeScoreList{
			curScore = v
			if curScore == 0{
				continue
			}
			nextSum := curSum + curScore
			if curSum < randNum && randNum <= nextSum{
				selectRouteIndex = k
				break
			} 
			curSum = nextSum
		}

		routeIndexList = append(routeIndexList, selectRouteIndex)

		routeScoreList[selectRouteIndex] = 0
		totalScore -= curScore
		selectRouteCount += 1
		if selectRouteCount >= routeCount{
			break
		}
	}
	return routeIndexList
}

func (s *acoServer) getSelectRouteListScore(routeList []route)([]int64){
	routeCount := len(routeList)
	routeScoreList := make([]int64, routeCount)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	for k, v := range routeList{
		curRoute := v
		curRouteScore := s.getCurRouteScore(curRoute)
		routeIndex := k
		routeScoreList[routeIndex] = curRouteScore
	}
	return routeScoreList
}

func (s *acoServer) getCurRouteScore(route route) int64{
	routeKey := route.ToString()

	curRouteState, exist := s.routeMetric[routeKey]

	if !exist{
		return defaultRate
	}

	curUnixTs := time.Now().Unix()
	elapsed := curUnixTs - (curRouteState.downloadDetail.endMs/1000)
	if elapsed < 0{
		elapsed = 0
	}

	if elapsed >= s.toZeroElapsed{
		return defaultRate/(curRouteState.downloadCount+1)
	}

	if curRouteState.downloadDetail.endMs == 0{
		return defaultRate/(curRouteState.downloadCount+1)
	}

	downloadDuration := float64(curRouteState.downloadDetail.endMs-curRouteState.downloadDetail.startMs)/1000.
	downloadSize := float64(curRouteState.downloadDetail.size)

	downloadRate := int64(downloadSize/downloadDuration)
	weightedDownloadRate := downloadRate/(curRouteState.downloadCount+1)

	reserveScale := 1.0 - (float64(elapsed)/float64(s.toZeroElapsed))

	scoreAtCurrent := int64(float64(weightedDownloadRate - defaultRate)*reserveScale) + defaultRate

	return scoreAtCurrent
}
