package bgp

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type VPN struct{}

type ListVPNsInput struct{}
type ListVPNsOutput struct{}

type GetVPNInput struct{}
type GetVPNOutput struct{}

type CreateVPNInput struct{}
type CreateVPNOutput struct{}

type UpdateVPNInput struct{}
type UpdateVPNOutput struct{}

type DeleteVPNInput struct{}
type DeleteVPNOutput struct{}

type VPNs interface {
	ListVPNs(input *ListVPNsInput) (*ListVPNsOutput, error)
	GetVPN(input *GetVPNInput) (*GetVPNOutput, error)
	CreateVPN(input *CreateVPNInput) (*CreateVPNOutput, error)
	UpdateVPN(input *UpdateVPNInput) (*UpdateVPNOutput, error)
	DeleteVPN(input *DeleteVPNInput) (*DeleteVPNOutput, error)
}

type VPNsImpl struct {
	Client *resty.Client
}

func (v VPNsImpl) ListVPNs(input *ListVPNsInput) (*ListVPNsOutput, error) {
	return nil, errors.New("not implemented")
}

func (v VPNsImpl) GetVPN(input *GetVPNInput) (*GetVPNOutput, error) {
	return nil, errors.New("not implemented")
}

func (v VPNsImpl) CreateVPN(input *CreateVPNInput) (*CreateVPNOutput, error) {
	return nil, errors.New("not implemented")
}

func (v VPNsImpl) UpdateVPN(input *UpdateVPNInput) (*UpdateVPNOutput, error) {
	return nil, errors.New("not implemented")
}

func (v VPNsImpl) DeleteVPN(input *DeleteVPNInput) (*DeleteVPNOutput, error) {
	return nil, errors.New("not implemented")
}
