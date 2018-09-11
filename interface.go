package ifconnmgr

import (
	"context"
	"time"

	inet "github.com/dms3-p2p/go-p2p-net"
	peer "github.com/dms3-p2p/go-p2p-peer"
)

type ConnManager interface {
	TagPeer(peer.ID, string, int)
	UntagPeer(peer.ID, string)
	GetTagInfo(peer.ID) *TagInfo
	TrimOpenConns(context.Context)
	Notifee() inet.Notifiee
}

type TagInfo struct {
	FirstSeen time.Time
	Value     int
	Tags      map[string]int
	Conns     map[string]time.Time
}
