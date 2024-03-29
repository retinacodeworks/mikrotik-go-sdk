package ipv4

import (
	"github.com/go-resty/resty/v2"
	"github.com/retinacodeworks/mikrotik-go-sdk/service/ipv4/firewall"
)

type Ipv4 struct {
	Client *resty.Client
}

func (i *Ipv4) Addresses() Addresses {
	return AddressesImpl{Client: i.Client}
}

func (i *Ipv4) Firewall() firewall.Firewall {
	return firewall.Firewall{Client: i.Client}
}
