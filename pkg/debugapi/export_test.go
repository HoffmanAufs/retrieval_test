// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debugapi

type (
	StatusResponse                    = statusResponse
	PingpongResponse                  = pingpongResponse
	PeerConnectResponse               = peerConnectResponse
	PeersResponse                     = peersResponse
	AddressesResponse                 = addressesResponse
	WelcomeMessageRequest             = welcomeMessageRequest
	WelcomeMessageResponse            = welcomeMessageResponse
	BalancesResponse                  = balancesResponse
	BalanceResponse                   = balanceResponse
	SettlementResponse                = settlementResponse
	SettlementsResponse               = settlementsResponse
)

var (
	ErrCantBalance         = errCantBalance
	ErrCantBalances        = errCantBalances
	ErrNoBalance           = errNoBalance
	ErrCantSettlementsPeer = errCantSettlementsPeer
	ErrCantSettlements     = errCantSettlements
	ErrInvalidAddress      = errInvalidAddress
)
