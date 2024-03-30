package routing

import (
	"github.com/go-resty/resty/v2"
	"github.com/retinacodeworks/mikrotik-go-sdk/service/routing/bgp"
)

type Routing struct {
	Client *resty.Client
}

func (r Routing) BGP() bgp.BGP {
	return bgp.BGP{Client: r.Client}
}
