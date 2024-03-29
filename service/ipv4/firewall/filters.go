package firewall

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type Filter struct {
	Action                  string `json:"action"`
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
	HwOffload               string `json:"hw-offload,omitempty"`
	IcmpOptions             string `json:"icmp-options,omitempty"`
	InBridgePort            string `json:"in-bridge-port,omitempty"`
	InBridgePortList        string `json:"in-bridge-port-list,omitempty"`
	InInterface             string `json:"in-interface,omitempty"`
	InInterfaceList         string `json:"in-interface-list,omitempty"`
	IngressPriority         int8   `json:"ingress-priority,omitempty"`
	IpsecPolicy             string `json:"ipsec-policy,omitempty"`
	Ipv4Options             string `json:"ipv4-options,omitempty"`
	JumpTarget              string `json:"jump-target,omitempty"`
	Layer7Protocol          string `json:"layer7-protocol,omitempty"`
	Limit                   string `json:"limit,omitempty"`
	Log                     string `json:"log,omitempty"`
	LogPrefix               string `json:"log-prefix,omitempty"`
	Nth                     string `json:"nth,omitempty"`
	OutBridgePort           string `json:"out-bridge-port,omitempty"`
	OutBridgePortList       string `json:"out-bridge-port-list,omitempty"`
	OutInterface            string `json:"out-interface,omitempty"`
	OutInterfaceList        string `json:"out-interface-list,omitempty"`
	PacketMark              string `json:"packet-mark,omitempty"`
	PacketSize              string `json:"packet-size,omitempty"`
	PerConnectionClassifier string `json:"per-connection-classifier,omitempty"`
	Port                    string `json:"port,omitempty"`
	Priority                int8   `json:"priority,omitempty"`
	Protocol                string `json:"protocol,omitempty"`
	Psd                     string `json:"psd,omitempty"`
	Random                  int8   `json:"random,omitempty"`
	RejectWith              string `json:"reject-with,omitempty"`
	RoutingTable            string `json:"routing-table,omitempty"`
	RoutingMark             string `json:"routing-mark,omitempty"`
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

type ListFilterRulesInput struct{}
type ListFilterRulesOutput struct{}

type GetFilterRuleInput struct{}
type GetFilterRuleOutput struct{}

type CreateFilterRuleInput struct{}
type CreateFilterRuleOutput struct{}

type UpdateFilterRuleInput struct{}
type UpdateFilterRuleOutput struct{}

type DeleteFilterRuleInput struct{}
type DeleteFilterRuleOutput struct{}

type Filters interface {
	ListFilterRules(input *ListFilterRulesInput) (*ListFilterRulesOutput, error)
	GetFilterRule(input *GetFilterRuleInput) (*GetFilterRuleOutput, error)
	CreateFilterRule(input *CreateFilterRuleInput) (*CreateFilterRuleOutput, error)
	UpdateFilterRule(input *UpdateFilterRuleInput) (*UpdateFilterRuleOutput, error)
	DeleteFilterRule(input *DeleteFilterRuleInput) (*DeleteFilterRuleOutput, error)
}

type FiltersImpl struct {
	Client *resty.Client
}

func (n FiltersImpl) ListFilterRules(input *ListFilterRulesInput) (*ListFilterRulesOutput, error) {
	return nil, errors.New("not implemented")
}

func (n FiltersImpl) GetFilterRule(input *GetFilterRuleInput) (*GetFilterRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n FiltersImpl) CreateFilterRule(input *CreateFilterRuleInput) (*CreateFilterRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n FiltersImpl) UpdateFilterRule(input *UpdateFilterRuleInput) (*UpdateFilterRuleOutput, error) {
	return nil, errors.New("not implemented")
}

func (n FiltersImpl) DeleteFilterRule(input *DeleteFilterRuleInput) (*DeleteFilterRuleOutput, error) {
	return nil, errors.New("not implemented")
}
