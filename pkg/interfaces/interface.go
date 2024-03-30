package interfaces

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ListInterfacesInput struct {
	Type string `url:"type,omitempty"`
}

type ListInterfacesOutput []Interface

type GetInterfaceInput struct {
	Id string
}

type GetInterfaceOutput Interface

type CreateInterfaceInput struct {
	Disabled bool   `json:"disabled,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
}

type CreateInterfaceOutput struct {
	Interface Interface
}

type UpdateInterfaceInput struct {
	Disabled bool   `json:"disabled,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
}

type UpdateInterfaceOutput struct {
	Interface Interface
}

type DeleteInterfaceInput struct {
	Id string
}

type DeleteInterfaceOutput struct {
	Success bool
}

type Interfaces interface {
	ListInterfaces(input *ListInterfacesInput) (ListInterfacesOutput, error)
	GetInterface(input *GetInterfaceInput) (*GetInterfaceOutput, error)
	CreateInterface(input *CreateInterfaceInput) (*CreateInterfaceOutput, error)
	UpdateInterface(input *UpdateInterfaceInput) (*UpdateInterfaceOutput, error)
	DeleteInterface(input *DeleteInterfaceInput) (*DeleteInterfaceOutput, error)
	Vlans() Vlans
	Bridges() Bridges
}

type Interface struct {
	Id          string `json:".id,omitempty"`
	ActualMtu   string `json:"actual-mtu,omitempty"`
	Disabled    string `json:"disabled,omitempty"`
	FpRxByte    string `json:"fp-rx-byte,omitempty"`
	FpRxPacket  string `json:"fp-rx-packet,omitempty"`
	FpTxByte    string `json:"fp-tx-byte,omitempty"`
	FpTxPacket  string `json:"fp-tx-packet,omitempty"`
	LinkDowns   string `json:"link-downs,omitempty"`
	Mtu         string `json:"mtu,omitempty"`
	Name        string `json:"name,omitempty"`
	Running     string `json:"running,omitempty"`
	RxByte      string `json:"rx-byte,omitempty"`
	RxDrop      string `json:"rx-drop,omitempty"`
	RxError     string `json:"rx-error,omitempty"`
	RxPacket    string `json:"rx-packet,omitempty"`
	TxByte      string `json:"tx-byte,omitempty"`
	TxDrop      string `json:"tx-drop,omitempty"`
	TxError     string `json:"tx-error,omitempty"`
	TxPacket    string `json:"tx-packet,omitempty"`
	TxQueueDrop string `json:"tx-queue-drop,omitempty"`
	Type        string `json:"type,omitempty"`
}

type InterfacesImpl struct {
	Client *resty.Client
}

func (i InterfacesImpl) Vlans() Vlans {
	return VlansImpl{Client: i.Client}
}

func (i InterfacesImpl) Bridges() Bridges {
	return BridgesImpl{Client: i.Client}
}

func (i InterfacesImpl) ListInterfaces(input *ListInterfacesInput) (ListInterfacesOutput, error) {
	var resp ListInterfacesOutput
	v, _ := query.Values(input)
	_, err := i.Client.R().
		SetResult(&resp).
		SetQueryString(v.Encode()).
		Get("/interface")
	return resp, err
}

func (i InterfacesImpl) GetInterface(input *GetInterfaceInput) (*GetInterfaceOutput, error) {
	var resp GetInterfaceOutput
	_, err := i.Client.R().
		SetResult(&resp).
		Get(fmt.Sprintf("/interface/%s", input.Id))
	return &resp, err
}

func (i InterfacesImpl) CreateInterface(input *CreateInterfaceInput) (*CreateInterfaceOutput, error) {
	return nil, errors.New("not implemented")
}

func (i InterfacesImpl) UpdateInterface(input *UpdateInterfaceInput) (*UpdateInterfaceOutput, error) {
	return nil, errors.New("not implemented")
}

func (i InterfacesImpl) DeleteInterface(input *DeleteInterfaceInput) (*DeleteInterfaceOutput, error) {
	return nil, errors.New("not implemented")
}
