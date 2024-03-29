package ipv4

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Address struct {
	Id        string `json:".id,omitempty"`
	Address   string `json:"address,omitempty"`
	Network   string `json:"network,omitempty"`
	Interface string `json:"interface,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

type ListAddressesInput struct{}
type ListAddressesOutput struct {
	Addresses []Address
}

type GetAddressInput struct {
	Id string
}

type GetAddressOutput struct {
	Address *Address
}

type CreateAddressInput struct {
	Address   string
	Network   string
	Comment   string
	Interface string
}

type CreateAddressOutput struct {
	Address *Address
}

type UpdateAddressInput struct {
	Id        string
	Address   string
	Network   string
	Comment   string
	Interface string
}
type UpdateAddressOutput struct {
	Address *Address
}

type DeleteAddressInput struct {
	Id string
}

type DeleteAddressOutput struct{}

type Addresses interface {
	ListAddresses(input *ListAddressesInput) (*ListAddressesOutput, error)
	GetAddress(input *GetAddressInput) (*GetAddressOutput, error)
	CreateAddress(input *CreateAddressInput) (*CreateAddressOutput, error)
	UpdateAddress(input *UpdateAddressInput) (*UpdateAddressOutput, error)
	DeleteAddress(input *DeleteAddressInput) (*DeleteAddressOutput, error)
}

type AddressesImpl struct {
	Client *resty.Client
}

func (a AddressesImpl) ListAddresses(input *ListAddressesInput) (*ListAddressesOutput, error) {
	var res []Address
	_, err := a.Client.R().
		SetResult(&res).
		Get("/ip/address")
	if err != nil {
		return nil, err
	}
	return &ListAddressesOutput{Addresses: res}, nil
}

func (a AddressesImpl) GetAddress(input *GetAddressInput) (*GetAddressOutput, error) {
	var res Address
	_, err := a.Client.R().
		SetResult(&res).
		Get(fmt.Sprintf("/ip/address/%s", input.Id))
	if err != nil {
		return nil, err
	}
	return &GetAddressOutput{Address: &res}, nil
}

//	curl -k -u admin: -X PUT https://10.155.101.214/rest/ip/address \
//	 --data '{"address": "192.168.111.111", "interface": "dummy"}' -H "content-type: application/json"
//
// {".id":"*A","actual-interface":"dummy","address":"192.168.111.111/32","disabled":"false",
// "dynamic":"false","interface":"dummy","invalid":"false","network":"192.168.111.111"}
func (a AddressesImpl) CreateAddress(input *CreateAddressInput) (*CreateAddressOutput, error) {
	var res Address
	_, err := a.Client.R().
		SetResult(&res).
		SetBody(Address{
			Address:   input.Address,
			Comment:   input.Comment,
			Interface: input.Interface,
			Network:   input.Network,
		}).
		Put("/ip/address")
	if err != nil {
		return nil, err
	}
	return &CreateAddressOutput{Address: &res}, nil
}

func (a AddressesImpl) UpdateAddress(input *UpdateAddressInput) (*UpdateAddressOutput, error) {
	var res Address
	_, err := a.Client.R().
		SetResult(&res).
		SetBody(Address{
			Address:   input.Address,
			Comment:   input.Comment,
			Interface: input.Interface,
			Network:   input.Network,
		}).
		Patch(fmt.Sprintf("/ip/address/%s", input.Id))
	if err != nil {
		return nil, err
	}
	return &UpdateAddressOutput{Address: &res}, nil
}

func (a AddressesImpl) DeleteAddress(input *DeleteAddressInput) (*DeleteAddressOutput, error) {
	_, err := a.Client.R().
		Delete(fmt.Sprintf("/ip/address/%s", input.Id))
	if err != nil {
		return nil, err
	}
	return nil, nil
}
