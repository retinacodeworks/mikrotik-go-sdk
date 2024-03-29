package interfaces

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ListInterfacesInput struct {
	Type string `url:"type"`
}

type ListInterfacesOutput struct {
	Interfaces []Interface
}

type GetInterfaceInput struct {
	Id string
}

type GetInterfaceOutput struct {
	Interface Interface
}

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
	ListInterfaces(input *ListInterfacesInput) (*ListInterfacesOutput, error)
	GetInterface(input *GetInterfaceInput) (*GetInterfaceOutput, error)
	CreateInterface(input *CreateInterfaceInput) (*CreateInterfaceOutput, error)
	UpdateInterface(input *UpdateInterfaceInput) (*UpdateInterfaceOutput, error)
	DeleteInterface(input *DeleteInterfaceInput) (*DeleteInterfaceOutput, error)
	Vlans() Vlans
}

type Interface struct {
	Id          string `json:".id,omitempty"`
	ActualMtu   int8   `json:"actual-mtu,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
	FpRxByte    int    `json:"fp-rx-byte,omitempty"`
	FpRxPacket  int    `json:"fp-rx-packet,omitempty"`
	FpTxByte    int    `json:"fp-tx-byte,omitempty"`
	FpTxPacket  int    `json:"fp-tx-packet,omitempty"`
	LinkDowns   int    `json:"link-downs,omitempty"`
	Mtu         int    `json:"mtu,omitempty"`
	Name        string `json:"name,omitempty"`
	Running     bool   `json:"running,omitempty"`
	RxByte      int    `json:"rx-byte,omitempty"`
	RxDrop      int    `json:"rx-drop,omitempty"`
	RxError     int    `json:"rx-error,omitempty"`
	RxPacket    int    `json:"rx-packet,omitempty"`
	TxByte      int    `json:"tx-byte,omitempty"`
	TxDrop      int    `json:"tx-drop,omitempty"`
	TxError     int    `json:"tx-error,omitempty"`
	TxPacket    int    `json:"tx-packet,omitempty"`
	TxQueueDrop int    `json:"tx-queue-drop,omitempty"`
	Type        string `json:"type,omitempty"`
}

type InterfacesImpl struct {
	Client *resty.Client
}

func (i InterfacesImpl) Vlans() Vlans {
	return VlansImpl{Client: i.Client}
}

func (i InterfacesImpl) ListInterfaces(input *ListInterfacesInput) (*ListInterfacesOutput, error) {
	var resp ListInterfacesOutput
	v, _ := query.Values(input)
	_, err := i.Client.R().
		SetResult(resp).
		SetQueryString(v.Encode()).
		Get("/interface")
	return &resp, err
}

func (i InterfacesImpl) GetInterface(input *GetInterfaceInput) (*GetInterfaceOutput, error) {
	var resp GetInterfaceOutput
	_, err := i.Client.R().
		SetResult(resp).
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
