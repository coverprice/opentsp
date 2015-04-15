// Copyright 2014 The Sporting Exchange Limited. All rights reserved.
// Use of this source code is governed by a free license that can be
// found in the LICENSE file.

package collect

import (
	"opentsp.org/contrib/collect-netscaler/nitro"
)

func init() {
	registerStatFunc("protocoltcp", protocolTCP)
}

func protocolTCP(emit emitFn, r *nitro.ResponseStat) {
	x := r.ProtocolTCP

	emit("protocol.tcp.ActiveServerConn", *x.ActiveServerConn)
	emit("protocol.tcp.ClientConnOpenedRate", *x.ClientConnOpenedRate)
	emit("protocol.tcp.CltFinRate", *x.CltFinRate)
	emit("protocol.tcp.CurClientConn", *x.CurClientConn)
	emit("protocol.tcp.CurClientConnClosing", *x.CurClientConnClosing)
	emit("protocol.tcp.CurClientConnEstablished", *x.CurClientConnEstablished)
	emit("protocol.tcp.CurClientConnOpening", *x.CurClientConnOpening)
	emit("protocol.tcp.CurServerConn", *x.CurServerConn)
	emit("protocol.tcp.CurServerConnClosing", *x.CurServerConnClosing)
	emit("protocol.tcp.CurServerConnEstablished", *x.CurServerConnEstablished)
	emit("protocol.tcp.CurServerConnOpening", *x.CurServerConnOpening)
	emit("protocol.tcp.FinWaitClosedRate", *x.FinWaitClosedRate)
	emit("protocol.tcp.PcbZombieCallRate", *x.PcbZombieCallRate)
	emit("protocol.tcp.RxBytesRate", *x.RxBytesRate)
	emit("protocol.tcp.RxPktsRate", *x.RxPktsRate)
	emit("protocol.tcp.ServerConnOpenedRate", *x.ServerConnOpenedRate)
	emit("protocol.tcp.SpareConn", *x.SpareConn)
	emit("protocol.tcp.SurgeQueueLen", *x.SurgeQueueLen)
	emit("protocol.tcp.SvrFinRate", *x.SvrFinRate)
	emit("protocol.tcp.SynFlushRate", *x.SynFlushRate)
	emit("protocol.tcp.SynHeldRate", *x.SynHeldRate)
	emit("protocol.tcp.SynProbeRate", *x.SynProbeRate)
	emit("protocol.tcp.SynRate", *x.SynRate)
	emit("protocol.tcp.TotServerConnOpened", *x.TotServerConnOpened)
	emit("protocol.tcp.TxBytesRate", *x.TxBytesRate)
	emit("protocol.tcp.TxPktsRate", *x.TxPktsRate)
	emit("protocol.tcp.WaitToDataRate", *x.WaitToDataRate)
	emit("protocol.tcp.WaitToSynRate", *x.WaitToSynRate)
	emit("protocol.tcp.ZombieActiveHalfCloseCltConnFlushedRate", *x.ZombieActiveHalfCloseCltConnFlushedRate)
	emit("protocol.tcp.ZombieActivehalfCloseSvrConnFlushedRate", *x.ZombieActivehalfCloseSvrConnFlushedRate)
	emit("protocol.tcp.ZombieCltConnFlushedRate", *x.ZombieCltConnFlushedRate)
	emit("protocol.tcp.ZombieHalfOpenCltConnFlushedRate", *x.ZombieHalfOpenCltConnFlushedRate)
	emit("protocol.tcp.ZombieHalfOpenSvrConnFlushedRate", *x.ZombieHalfOpenSvrConnFlushedRate)
	emit("protocol.tcp.ZombiePassiveHalfClosecltConnFlushedRate", *x.ZombiePassiveHalfClosecltConnFlushedRate)
	emit("protocol.tcp.ZombiePassivehalfCloseSrvConnFlushedRate", *x.ZombiePassivehalfCloseSrvConnFlushedRate)
	emit("protocol.tcp.ZombieSvrConnFlushedRate", *x.ZombieSvrConnFlushedRate)

	emit("protocol.tcp.errors type=AnyPortFailRate", *x.ErrAnyPortFailRate)
	emit("protocol.tcp.errors type=BadChecksumRate", *x.ErrBadChecksumRate)
	emit("protocol.tcp.errors type=BadstateConnRate", *x.ErrBadstateConnRate)
	emit("protocol.tcp.errors type=CipAllocRate", *x.ErrCipAllocRate)
	emit("protocol.tcp.errors type=CltHoleRate", *x.ErrCltHoleRate)
	emit("protocol.tcp.errors type=CltOutOfOrderRate", *x.ErrCltOutOfOrderRate)
	emit("protocol.tcp.errors type=CltRetrasmitRate", *x.ErrCltRetrasmitRate)
	emit("protocol.tcp.errors type=CookiePktMssRejectRate", *x.ErrCookiePktMssRejectRate)
	emit("protocol.tcp.errors type=CookiePktSeqDropRate", *x.ErrCookiePktSeqDropRate)
	emit("protocol.tcp.errors type=CookiePktSeqRejectRate", *x.ErrCookiePktSeqRejectRate)
	emit("protocol.tcp.errors type=CookiePktSigRejectRate", *x.ErrCookiePktSigRejectRate)
	emit("protocol.tcp.errors type=DataAfterFinRate", *x.ErrDataAfterFinRate)
	emit("protocol.tcp.errors type=FastRetransmissionsRate", *x.ErrFastRetransmissionsRate)
	emit("protocol.tcp.errors type=FifthRetransmissionsRate", *x.ErrFifthRetransmissionsRate)
	emit("protocol.tcp.errors type=FinGiveUpRate", *x.ErrFinGiveUpRate)
	emit("protocol.tcp.errors type=FinRetryRate", *x.ErrFinRetryRate)
	emit("protocol.tcp.errors type=FirstRetransmissionsRate", *x.ErrFirstRetransmissionsRate)
	emit("protocol.tcp.errors type=FourthRetransmissionsRate", *x.ErrFourthRetransmissionsRate)
	emit("protocol.tcp.errors type=FullRetrasmitRate", *x.ErrFullRetrasmitRate)
	emit("protocol.tcp.errors type=IpPortFailRate", *x.ErrIpPortFailRate)
	emit("protocol.tcp.errors type=OutOfWindowPktsRate", *x.ErrOutOfWindowPktsRate)
	emit("protocol.tcp.errors type=PartialRetrasmitRate", *x.ErrPartialRetrasmitRate)
	emit("protocol.tcp.errors type=RetransmitGiveUpRate", *x.ErrRetransmitGiveUpRate)
	emit("protocol.tcp.errors type=RetransmitRate", *x.ErrRetransmitRate)
	emit("protocol.tcp.errors type=RstInTimeWaitRate", *x.ErrRstInTimeWaitRate)
	emit("protocol.tcp.errors type=RstNonEstRate", *x.ErrRstNonEstRate)
	emit("protocol.tcp.errors type=RstOutOfWindowRate", *x.ErrRstOutOfWindowRate)
	emit("protocol.tcp.errors type=RstRate", *x.ErrRstRate)
	emit("protocol.tcp.errors type=RstThresholdRate", *x.ErrRstThresholdRate)
	emit("protocol.tcp.errors type=SecondRetransmissionsRate", *x.ErrSecondRetransmissionsRate)
	emit("protocol.tcp.errors type=SentRstRate", *x.ErrSentRstRate)
	emit("protocol.tcp.errors type=SeventhRetransmissionsRate", *x.ErrSeventhRetransmissionsRate)
	emit("protocol.tcp.errors type=SixthRetransmissionsRate", *x.ErrSixthRetransmissionsRate)
	emit("protocol.tcp.errors type=StrayPktRate", *x.ErrStrayPktRate)
	emit("protocol.tcp.errors type=SvrHoleRate", *x.ErrSvrHoleRate)
	emit("protocol.tcp.errors type=SvrOutOfOrderRate", *x.ErrSvrOutOfOrderRate)
	emit("protocol.tcp.errors type=SvrRetrasmitRate", *x.ErrSvrRetrasmitRate)
	emit("protocol.tcp.errors type=SynDroppedCongestionRate", *x.ErrSynDroppedCongestionRate)
	emit("protocol.tcp.errors type=SynGiveupRate", *x.ErrSynGiveupRate)
	emit("protocol.tcp.errors type=SynInSynrcvdRate", *x.ErrSynInSynrcvdRate)
	emit("protocol.tcp.errors type=SynRetryRate", *x.ErrSynRetryRate)
	emit("protocol.tcp.errors type=SynSentBadackRate", *x.ErrSynSentBadackRate)
	emit("protocol.tcp.errors type=SyninestRate", *x.ErrSyninestRate)
	emit("protocol.tcp.errors type=ThirdRetransmissionsRate", *x.ErrThirdRetransmissionsRate)
}
