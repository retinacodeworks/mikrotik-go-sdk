package firewall

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type Nat struct {
	Action                  string `json:"action"`
	AddressList             string `json:"address-list,omitempty"`
	AddressListTimeout      string `json:"address-list-timeout,omitempty"`
	Chain                   string `json:"chain"`
	Comment                 string `json:"comment,omitempty"`
	ConnectionBytes         string `json:"connection-bytes,omitempty"`
	ConnectionLimit         string `json:"connection-limit,omitempty"`
	ConnectionMark          string `json:"connection-mark,omitempty"`
	ConnectionRate          string `json:"connection-rate,omitempty"`
	ConnectionType          string `json:"connection-type,omitempty"`
	Content                 string `json:"content,omitempty"`
	Dscp                    int8   `json:"dscp,omitempty"`
	DstAddress              string `json:"dst-address,omitempty"`
	DstAddressList          string `json:"dst-address-list,omitempty"`
	DstAddressType          string `json:"dst-address-type,omitempty"`
	DstLimit                string `json:"dst-limit,omitempty"`
	DstPort                 string `json:"dst-port,omitempty"`
	Fragment                string `json:"fragment,omitempty"`
	Hotspot                 string `json:"hotspot,omitempty"`
	IcmpOptions             string `json:"icmp-options,omitempty"`
	InBridgePort            string `json:"in-bridge-port,omitempty"`
	InInterface             string `json:"in-interface,omitempty"`
	IngressPriority         int8   `json:"ingress-priority,omitempty"`
	IpsecPolicy             string `json:"ipsec-policy,omitempty"`
	Ipv4Options             string `json:"ipv4-options,omitempty"`
	JumpTarget              string `json:"jump-target,omitempty"`
	Layer7Protocol          string `json:"layer7-protocol,omitempty"`
	Limit                   string `json:"limit,omitempty"`
	Log                     string `json:"log,omitempty"`
	LogPrefix               string `json:"log-prefix,omitempty"`
	OutBridgePort           string `json:"out-bridge-port,omitempty"`
	OutInterface            string `json:"out-interface,omitempty"`
	PacketMark              string `json:"packet-mark,omitempty"`
	PacketSize              string `json:"packet-size,omitempty"`
	PerConnectionClassifier string `json:"per-connection-classifier,omitempty"`
	Port                    string `json:"port,omitempty"`
	Protocol                string `json:"protocol,omitempty"`
	Psd                     string `json:"psd,omitempty"`
	Random                  int8   `json:"random,omitempty"`
	RoutingMark             string `json:"routing-mark,omitempty"`
	SameNotByDst            string `json:"same-not-by-dst,omitempty"`
	SrcAddress              string `json:"src-address,omitempty"`
	SrcAddressList          string `json:"src-address-list,omitempty"`
	SrcAddressType          string `json:"src-address-type,omitempty"`
	SrcPort                 string `json:"src-port,omitempty"`
	SrcMacAddress           string `json:"src-mac-address,omitempty"`
	TcpMss                  string `json:"tcp-mss,omitempty"`
	Time                    string `json:"time,omitempty"`
	ToAddresses             string `json:"to-addresses,omitempty"`
	ToPorts                 string `json:"to-ports,omitempty"`
	Ttl                     int8   `json:"ttl,omitempty"`
}

type ListNatRulesInput struct{}
type ListNatRulesOutput struct{}

type GetNatRuleInput struct{}
type GetNatRuleOutput struct{}

type CreateNatRuleInput struct{}
type CreateNatRuleOutput struct{}

type UpdateNatRuleInput struct{}
type UpdateNatRuleOutput struct{}

type DeleteNatRuleInput struct{}
type DeleteNatRuleOutput struct{}

type Nats interface {
	ListNatRules(input *ListNatRulesInput) (*ListNatRulesOutput, error)
	GetNatRule(input *GetNatRuleInput) (*GetNatRuleOutput, error)
	CreateNatRule(input *CreateNatRuleInput) (*CreateNatRuleOutput, error)
	UpdateNatRule(input *UpdateNatRuleInput) (*UpdateNatRuleOutput, error)
	DeleteNatRule(input *DeleteNatRuleInput) (*DeleteNatRuleOutput, error)
}

type NatsImpl struct {
	Client *resty.Client
}

func (n NatsImpl) ListNatRules(input *ListNatRulesInput) (*ListNatRulesOutput, error) {
	return nil, errors.New("not implemented")
}

func (n NatsImpl) GetNatRule(input *GetNatRuleInput) (*GetNatRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n NatsImpl) CreateNatRule(input *CreateNatRuleInput) (*CreateNatRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n NatsImpl) UpdateNatRule(input *UpdateNatRuleInput) (*UpdateNatRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n NatsImpl) DeleteNatRule(input *DeleteNatRuleInput) (*DeleteNatRuleOutput, error) {
	return nil, errors.New("not implemented")
}
