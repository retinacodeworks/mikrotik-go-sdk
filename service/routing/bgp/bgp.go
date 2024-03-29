package bgp

import "github.com/go-resty/resty/v2"

type BGP struct {
	Client *resty.Client
}

func (b *BGP) Connections() Connections {
	return ConnectionsImpl{Client: b.Client}
}

func (b *BGP) Templates() Templates {
	return TemplatesImpl{Client: b.Client}
}

func (b *BGP) Vpns() VPNs {
	return VPNsImpl{Client: b.Client}
}
