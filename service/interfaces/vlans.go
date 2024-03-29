package interfaces

import "github.com/go-resty/resty/v2"

type Vlans interface {
	ListVlans() (*ListVlanOutput, error)
	CreateVlan(req *CreateVlanInput) (*CreateVlanOutput, error)
}

type ListVlanOutput struct {
	Interfaces []Interface
}

type CreateVlanInput struct {
}

type CreateVlanOutput struct {
}

type VlansImpl struct {
	Client *resty.Client
}

func (v VlansImpl) ListVlans() (*ListVlanOutput, error) {
	var resp ListVlanOutput
	_, err := v.Client.R().
		SetResult(resp).
		SetQueryParams(map[string]string{
			"type": "vlan",
		}).
		Get("/interface/print")
	return &resp, err
}

func (v VlansImpl) CreateVlan(req *CreateVlanInput) (*CreateVlanOutput, error) {
	return nil, nil
}
