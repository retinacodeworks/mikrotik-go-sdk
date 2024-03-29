package bgp

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type ListBGPConnectionsInput struct{}
type ListBGPConnectionsOutput struct{}

type GetBGPConnectionInput struct{}
type GetBGPConnectionOutput struct{}

type CreateBGPConnectionInput struct{}
type CreateBGPConnectionOutput struct{}

type UpdateBGPConnectionInput struct{}
type UpdateBGPConnectionOutput struct{}

type DeleteBGPConnectionInput struct{}
type DeleteBGPConnectionOutput struct{}

type Connections interface {
	ListBGPConnections(input *ListBGPConnectionsInput) (*ListBGPConnectionsOutput, error)
	GetBGPConnection(input *GetBGPConnectionInput) (*GetBGPConnectionOutput, error)
	CreateBGPConnection(input *CreateBGPConnectionInput) (*CreateBGPConnectionOutput, error)
	UpdateBGPConnection(input *UpdateBGPConnectionInput) (*UpdateBGPConnectionOutput, error)
	DeleteBGPConnection(input *DeleteBGPConnectionInput) (*DeleteBGPConnectionOutput, error)
}

type Connection struct {
	Name      string          `json:"name"`
	Connect   string          `json:"connect,omitempty"`
	Listen    string          `json:"listen,omitempty"`
	Local     ConnectionLocal `json:"local,omitempty"`
	TcpMd5Key string          `json:"tcp-md5-key,omitempty"`
	Templates string          `json:"templates,omitempty"`
}

type ConnectionLocal struct {
	Address string `json:"address,omitempty"`
	Port    int16  `json:"port,omitempty"`
	Role    string `json:"role,omitempty"`
	Ttl     int8   `json:"ttl,omitempty"`
}

type ConnectionRemote struct {
	Address   string `json:"address,omitempty"`
	Port      int16  `json:"port,omitempty"`
	As        int32  `json:"as,omitempty"`
	AllowedAs string `json:"allowed-as,omitempty"`
	Ttl       int8   `json:"ttl,omitempty"`
}

type ConnectionsImpl struct {
	Client *resty.Client
}

func (c ConnectionsImpl) ListBGPConnections(input *ListBGPConnectionsInput) (*ListBGPConnectionsOutput, error) {
	return nil, errors.New("not implemented")
}

func (c ConnectionsImpl) GetBGPConnection(input *GetBGPConnectionInput) (*GetBGPConnectionOutput, error) {
	return nil, errors.New("not implemented")
}

func (c ConnectionsImpl) CreateBGPConnection(input *CreateBGPConnectionInput) (*CreateBGPConnectionOutput, error) {
	return nil, errors.New("not implemented")
}

func (c ConnectionsImpl) UpdateBGPConnection(input *UpdateBGPConnectionInput) (*UpdateBGPConnectionOutput, error) {
	return nil, errors.New("not implemented")
}

func (c ConnectionsImpl) DeleteBGPConnection(input *DeleteBGPConnectionInput) (*DeleteBGPConnectionOutput, error) {
	return nil, errors.New("not implemented")
}
