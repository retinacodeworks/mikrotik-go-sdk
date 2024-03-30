package firewall

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type Mangle struct {
	Action                  string `json:"action"`
	AddressList             string `json:"address-list,omitempty"`
	AddressListTimeout      string `json:"address-list-timeout,omitempty"`
	Chain                   string `json:"chain"`
	Comment                 string `json:"comment,omitempty"`
	ConnectionBytes         string `json:"connection-bytes,omitempty"`
	ConnectionLimit         string `json:"connection-limit,omitempty"`
	ConnectionMark          string `json:"connection-mark,omitempty"`
	ConnectionNatState      string `json:"connection-nat-state,omitempty"`
	ConnectionRate          string `json:"connection-rate,omitempty"`
	ConnectionState         string `json:"connection-state,omitempty"`
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
	InBridgePortList        string `json:"in-bridge-port-list,omitempty"`
	InInterface             string `json:"in-interface,omitempty"`
	IngressPriority         int8   `json:"ingress-priority,omitempty"`
	IpsecPolicy             string `json:"ipsec-policy,omitempty"`
	Ipv4Options             string `json:"ipv4-options,omitempty"`
	JumpTarget              string `json:"jump-target,omitempty"`
	Layer7Protocol          string `json:"layer7-protocol,omitempty"`
	Limit                   string `json:"limit,omitempty"`
	Log                     string `json:"log,omitempty"`
	LogPrefix               string `json:"log-prefix,omitempty"`
	NewDscp                 string `json:"new-dscp,omitempty"`
	NewMss                  int8   `json:"new-mss,omitempty"`
	NewPacketMark           string `json:"new-packet-mark,omitempty"`
	NewPriority             string `json:"new-priority,omitempty"`
	NewRoutingMark          string `json:"new-routing-mark,omitempty"`
	NewTtl                  string `json:"new-ttl,omitempty"`
	Nth                     string `json:"nth,omitempty"`
	OutBridgePort           string `json:"out-bridge-port,omitempty"`
	OutInterface            string `json:"out-interface,omitempty"`
	PacketMark              string `json:"packet-mark,omitempty"`
	PacketSize              string `json:"packet-size,omitempty"`
	Passthrough             string `json:"passthrough,omitempty"`
	PerConnectionClassifier string `json:"per-connection-classifier,omitempty"`
	Port                    string `json:"port,omitempty"`
	Priority                int8   `json:"priority,omitempty"`
	Protocol                string `json:"protocol,omitempty"`
	Psd                     string `json:"psd,omitempty"`
	Random                  int8   `json:"random,omitempty"`
	RoutingMark             string `json:"routing-mark,omitempty"`
	RouteDst                string `json:"route-dst,omitempty"`
	SrcAddress              string `json:"src-address,omitempty"`
	SrcAddressList          string `json:"src-address-list,omitempty"`
	SrcAddressType          string `json:"src-address-type,omitempty"`
	SrcPort                 string `json:"src-port,omitempty"`
	SrcMacAddress           string `json:"src-mac-address,omitempty"`
	TcpFlags                string `json:"tcp-flags,omitempty"`
	TcpMss                  string `json:"tcp-mss,omitempty"`
	Time                    string `json:"time,omitempty"`
	TlsHost                 string `json:"tls-host,omitempty"`
	Ttl                     int8   `json:"ttl,omitempty"`
}

type ListMangleRulesInput struct{}
type ListMangleRulesOutput struct{}

type GetMangleRuleInput struct{}
type GetMangleRuleOutput struct{}

type CreateMangleRuleInput struct{}
type CreateMangleRuleOutput struct{}

type UpdateMangleRuleInput struct{}
type UpdateMangleRuleOutput struct{}

type DeleteMangleRuleInput struct{}
type DeleteMangleRuleOutput struct{}

type Mangles interface {
	ListMangleRules(input *ListMangleRulesInput) (*ListMangleRulesOutput, error)
	GetMangleRule(input *GetMangleRuleInput) (*GetMangleRuleOutput, error)
	CreateMangleRule(input *CreateMangleRuleInput) (*CreateMangleRuleOutput, error)
	UpdateMangleRule(input *UpdateMangleRuleInput) (*UpdateMangleRuleOutput, error)
	DeleteMangleRule(input *DeleteMangleRuleInput) (*DeleteMangleRuleOutput, error)
}

type ManglesImpl struct {
	Client *resty.Client
}

func (n ManglesImpl) ListMangleRules(input *ListMangleRulesInput) (*ListMangleRulesOutput, error) {
	return nil, errors.New("not implemented")
}

func (n ManglesImpl) GetMangleRule(input *GetMangleRuleInput) (*GetMangleRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n ManglesImpl) CreateMangleRule(input *CreateMangleRuleInput) (*CreateMangleRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n ManglesImpl) UpdateMangleRule(input *UpdateMangleRuleInput) (*UpdateMangleRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n ManglesImpl) DeleteMangleRule(input *DeleteMangleRuleInput) (*DeleteMangleRuleOutput, error) {
	return nil, errors.New("not implemented")
}
