package bgp

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ListBGPConnectionsInput struct{}
type ListBGPConnectionsOutput []Connection

type GetBGPConnectionInput struct {
	Id string
}
type GetBGPConnectionOutput Connection

type CreateBGPConnectionInput Connection
type CreateBGPConnectionOutput Connection

type UpdateBGPConnectionInput Connection
type UpdateBGPConnectionOutput Connection

type DeleteBGPConnectionInput struct {
	Id string
}
type DeleteBGPConnectionOutput struct{}

type Connections interface {
	ListBGPConnections(input *ListBGPConnectionsInput) (*ListBGPConnectionsOutput, error)
	GetBGPConnection(input *GetBGPConnectionInput) (*GetBGPConnectionOutput, error)
	CreateBGPConnection(input *CreateBGPConnectionInput) (*CreateBGPConnectionOutput, error)
	UpdateBGPConnection(input *UpdateBGPConnectionInput) (*UpdateBGPConnectionOutput, error)
	DeleteBGPConnection(input *DeleteBGPConnectionInput) (*DeleteBGPConnectionOutput, error)
}

type Connection struct {
	Id              string `json:".id,omitempty"`
	As              string `json:"as,omitempty"`
	Name            string `json:"name"`
	AddressFamilies string `json:"address-families,omitempty"`
	Connect         string `json:"connect,omitempty"`
	Listen          string `json:"listen,omitempty"`
	LocalAddress    string `json:"local.address,omitempty"`
	LocalRole       string `json:"local.role,omitempty"`
	LocalPort       string `json:"local.port,omitempty"`
	LocalTtl        string `json:"local.ttl,omitempty"`
	RemoteAddress   string `json:"remote.address,omitempty"`
	RemotePort      string `json:"remote.port,omitempty"`
	RemoteAs        string `json:"remote.as,omitempty"`
	RemoteAllowedAs string `json:"remote.allowed-at,omitempty"`
	RemoteTtl       string `json:"remote.ttl,omitempty"`
	RoutingTable    string `json:"routing-table,omitempty"`
	TcpMd5Key       string `json:"tcp-md5-key,omitempty"`
	Templates       string `json:"templates,omitempty"`
	Disabled        string `json:"disabled,omitempty"`
	Inactive        string `json:"inactive,omitempty"`
}

type ConnectionsImpl struct {
	Client *resty.Client
}

func (c ConnectionsImpl) ListBGPConnections(input *ListBGPConnectionsInput) (*ListBGPConnectionsOutput, error) {
	var resp ListBGPConnectionsOutput
	v, _ := query.Values(input)
	_, err := c.Client.R().
		SetResult(&resp).
		SetQueryString(v.Encode()).
		Get("/routing/bgp/connection")
	return &resp, err
}

func (c ConnectionsImpl) GetBGPConnection(input *GetBGPConnectionInput) (*GetBGPConnectionOutput, error) {
	var resp GetBGPConnectionOutput
	_, err := c.Client.R().
		SetResult(&resp).
		Get(fmt.Sprintf("/routing/bridge/connection/%s", input.Id))
	return &resp, err
}

func (c ConnectionsImpl) CreateBGPConnection(input *CreateBGPConnectionInput) (*CreateBGPConnectionOutput, error) {
	var res CreateBGPConnectionOutput
	_, err := c.Client.R().
		SetResult(&res).
		SetBody(input).
		Put("/routing/bridge/connection")
	return &res, err
}

func (c ConnectionsImpl) UpdateBGPConnection(input *UpdateBGPConnectionInput) (*UpdateBGPConnectionOutput, error) {
	var res UpdateBGPConnectionOutput
	_, err := c.Client.R().
		SetResult(&res).
		SetBody(input).
		Patch(fmt.Sprintf("/routing/bridge/connection/%s", input.Id))
	return &res, err
}

func (c ConnectionsImpl) DeleteBGPConnection(input *DeleteBGPConnectionInput) (*DeleteBGPConnectionOutput, error) {
	_, err := c.Client.R().
		Delete(fmt.Sprintf("/routing/bridge/connection/%s", input.Id))
	return nil, err
}
