//go:build anotherwebrtc

// +build anotherwebrtc

package compat

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/config"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"

	noise "github.com/libp2p/go-libp2p/p2p/security/noise"
	tls "github.com/libp2p/go-libp2p/p2p/security/tls"
	webrtc "github.com/libp2p/go-libp2p/p2p/transport/webrtc"
)

type PeerAddrInfo = peer.AddrInfo

func NewLibp2(ctx context.Context, secureChannel string, opts ...config.Option) (host.Host, error) {
	security := getSecurityByName(secureChannel)
	return libp2p.New(
		append(opts, security, libp2p.Transport(webrtc.New))...,
	)
}

func getSecurityByName(secureChannel string) libp2p.Option {
	switch secureChannel {
	case "noise":
		return libp2p.Security(noise.ID, noise.New)
	case "tls":
		return libp2p.Security(tls.ID, tls.New)
	}
	panic(fmt.Sprintf("unknown secure channel: %s", secureChannel))
}
