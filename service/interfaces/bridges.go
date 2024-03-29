package interfaces

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type Bridge struct {
	Id                string `json:".id,omitempty"`
	ActualMtu         string `json:"actual-mtu,omitempty"`
	AgeingTime        string `json:"ageing-time,omitempty"`
	Arp               string `json:"arp,omitempty"`
	ArpTimeout        string `json:"arp-timeout,omitempty"`
	AutoMac           string `json:"auto-mac,omitempty"`
	DhcpSnooping      string `json:"dhcp-snooping,omitempty"`
	Disabled          string `json:"disabled,omitempty"`
	FastForward       string `json:"fast-forward,omitempty"`
	ForwardDelay      string `json:"forward-delay,omitempty"`
	IgmpSnooping      string `json:"igmp-snooping,omitempty"`
	L2Mtu             string `json:"l2mtu,omitempty"`
	MacAddress        string `json:"mac-address,omitempty"`
	MaxMessageAge     string `json:"max-message-age,omitempty"`
	Mtu               string `json:"mtu,omitempty"`
	Name              string `json:"name,omitempty"`
	Priority          string `json:"priority,omitempty"`
	ProtocolMode      string `json:"protocol-mode,omitempty"`
	Running           string `json:"running,omitempty"`
	TransmitHoldCount string `json:"transmit-hold-count,omitempty"`
	VlanFiltering     string `json:"vlan-filtering,omitempty"`
}

type ListBridgesInput struct{}

type ListBridgesOutput []Bridge

type GetBridgeInput struct {
	Id string
}

type GetBridgeOutput Bridge

type CreateBridgeInput struct {
	Id string
}

type CreateBridgeOutput struct {
	Bridge Bridge
}

type UpdateBridgeInput struct {
	Id string
}

type UpdateBridgeOutput struct {
	Bridge Bridge
}

type DeleteBridgeInput struct {
	Id string
}
type DeleteBridgeOutput struct {
	Success bool
}

type Bridges interface {
	ListBridges(input *ListBridgesInput) (*ListBridgesOutput, error)
	GetBridge(input *GetBridgeInput) (*GetBridgeOutput, error)
	CreateBridge(input *CreateBridgeInput) (*CreateBridgeOutput, error)
	UpdateBridge(input *UpdateBridgeInput) (*UpdateBridgeOutput, error)
	DeleteBridge(input *DeleteBridgeInput) (*DeleteBridgeOutput, error)
}

type BridgesImpl struct {
	Client *resty.Client
}

func (b BridgesImpl) ListBridges(input *ListBridgesInput) (*ListBridgesOutput, error) {
	var resp ListBridgesOutput
	v, _ := query.Values(input)
	_, err := b.Client.R().
		SetResult(&resp).
		SetQueryString(v.Encode()).
		Get("/interface/bridge")
	return &resp, err
}

func (b BridgesImpl) GetBridge(input *GetBridgeInput) (*GetBridgeOutput, error) {
	var resp GetBridgeOutput
	_, err := b.Client.R().
		SetResult(&resp).
		Get(fmt.Sprintf("/interface/bridge/%s", input.Id))
	return &resp, err
}

func (b BridgesImpl) CreateBridge(input *CreateBridgeInput) (*CreateBridgeOutput, error) {
	return nil, errors.New("not implemented")
}

func (b BridgesImpl) UpdateBridge(input *UpdateBridgeInput) (*UpdateBridgeOutput, error) {
	return nil, errors.New("not implemented")
}

func (b BridgesImpl) DeleteBridge(input *DeleteBridgeInput) (*DeleteBridgeOutput, error) {
	return nil, errors.New("not implemented")
}
