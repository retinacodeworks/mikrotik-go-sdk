package firewall

import "github.com/go-resty/resty/v2"

type Firewall struct {
	Client *resty.Client
}

func (f *Firewall) Nats() Nats {
	return NatsImpl{Client: f.Client}
}

func (f *Firewall) Filters() Filters {
	return FiltersImpl{Client: f.Client}
}

func (f *Firewall) Mangles() Mangles {
	return ManglesImpl{Client: f.Client}
}

func (f *Firewall) AddressLists() AddressLists {
	return AddressListsImpl{Client: f.Client}
}
