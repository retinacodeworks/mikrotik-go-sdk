package firewall

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type AddressList struct {
	Address string `json:"address"`
	Dynamic string `json:"dynamic,omitempty"`
	List    string `json:"list"`
	Timeout string `json:"timeout,omitempty"`
}

type ListAddressListsInput struct{}
type ListAddressListsOutput struct{}

type GetAddressListInput struct{}
type GetAddressListOutput struct{}

type CreateAddressListInput struct{}
type CreateAddressListOutput struct{}

type UpdateAddressListInput struct{}
type UpdateAddressListOutput struct{}

type DeleteAddressListInput struct{}
type DeleteAddressListOutput struct{}

type AddressLists interface {
	ListAddressLists(input *ListAddressListsInput) (*ListAddressListsOutput, error)
	GetAddressList(input *GetAddressListInput) (*GetAddressListOutput, error)
	CreateAddressList(input *CreateAddressListInput) (*CreateAddressListOutput, error)
	UpdateAddressList(input *UpdateAddressListInput) (*UpdateAddressListOutput, error)
	DeleteAddressList(input *DeleteAddressListInput) (*DeleteAddressListOutput, error)
}

type AddressListsImpl struct {
	Client *resty.Client
}

func (n AddressListsImpl) ListAddressLists(input *ListAddressListsInput) (*ListAddressListsOutput, error) {
	return nil, errors.New("not implemented")
}

func (n AddressListsImpl) GetAddressList(input *GetAddressListInput) (*GetAddressListOutput, error) {
	return nil, errors.New("not implemented")
}

func (n AddressListsImpl) CreateAddressList(input *CreateAddressListInput) (*CreateAddressListOutput, error) {
	return nil, errors.New("not implemented")
}

func (n AddressListsImpl) UpdateAddressList(input *UpdateAddressListInput) (*UpdateAddressListOutput, error) {
	return nil, errors.New("not implemented")
}

func (n AddressListsImpl) DeleteAddressList(input *DeleteAddressListInput) (*DeleteAddressListOutput, error) {
	return nil, errors.New("not implemented")
}
