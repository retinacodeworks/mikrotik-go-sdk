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
type ListAddressesOutput []Address

type GetAddressInput struct {
	Id string
}
type GetAddressOutput Address

type CreateAddressInput Address
type CreateAddressOutput Address

type UpdateAddressInput Address
type UpdateAddressOutput Address

type DeleteAddressInput Address
type DeleteAddressOutput Address

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
	var res ListAddressesOutput
	_, err := a.Client.R().
		SetResult(&res).
		Get("/ip/address")
	return &res, err
}

func (a AddressesImpl) GetAddress(input *GetAddressInput) (*GetAddressOutput, error) {
	var res GetAddressOutput
	_, err := a.Client.R().
		SetResult(&res).
		Get(fmt.Sprintf("/ip/address/%s", input.Id))
	return &res, err
}

func (a AddressesImpl) CreateAddress(input *CreateAddressInput) (*CreateAddressOutput, error) {
	var res CreateAddressOutput
	_, err := a.Client.R().
		SetResult(&res).
		SetBody(Address{
			Address:   input.Address,
			Comment:   input.Comment,
			Interface: input.Interface,
			Network:   input.Network,
		}).
		Put("/ip/address")
	return &res, err
}

func (a AddressesImpl) UpdateAddress(input *UpdateAddressInput) (*UpdateAddressOutput, error) {
	var res UpdateAddressOutput
	_, err := a.Client.R().
		SetResult(&res).
		SetBody(Address{
			Address:   input.Address,
			Comment:   input.Comment,
			Interface: input.Interface,
			Network:   input.Network,
		}).
		Patch(fmt.Sprintf("/ip/address/%s", input.Id))
	return &res, err
}

func (a AddressesImpl) DeleteAddress(input *DeleteAddressInput) (*DeleteAddressOutput, error) {
	_, err := a.Client.R().
		Delete(fmt.Sprintf("/ip/address/%s", input.Id))
	return nil, err
}
