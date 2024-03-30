package firewall

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type Filter struct {
	Id                      string `json:".id,omitempty"`
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
	Dscp                    string `json:"dscp,omitempty"`
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
	IngressPriority         string `json:"ingress-priority,omitempty"`
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
	Priority                string `json:"priority,omitempty"`
	Protocol                string `json:"protocol,omitempty"`
	Psd                     string `json:"psd,omitempty"`
	Random                  string `json:"random,omitempty"`
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
	Ttl                     string `json:"ttl,omitempty"`
}

type ListFilterRulesInput struct{}
type ListFilterRulesOutput []Filter

type GetFilterRuleInput struct {
	Id string
}
type GetFilterRuleOutput Filter

type CreateFilterRuleInput Filter
type CreateFilterRuleOutput Filter

type UpdateFilterRuleInput Filter
type UpdateFilterRuleOutput Filter

type DeleteFilterRuleInput struct {
	Id string
}

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
	var resp ListFilterRulesOutput
	v, _ := query.Values(input)
	_, err := n.Client.R().
		SetResult(&resp).
		SetQueryString(v.Encode()).
		Get("/ip/firewall/filter")
	return &resp, err
}

func (n FiltersImpl) GetFilterRule(input *GetFilterRuleInput) (*GetFilterRuleOutput, error) {
	var resp GetFilterRuleOutput
	_, err := n.Client.R().
		SetResult(&resp).
		Get(fmt.Sprintf("/ip/firewall/filter/%s", input.Id))
	return &resp, err
}

func (n FiltersImpl) CreateFilterRule(input *CreateFilterRuleInput) (*CreateFilterRuleOutput, error) {
	var res CreateFilterRuleOutput
	_, err := n.Client.R().
		SetResult(&res).
		SetBody(input).
		Put("/ip/firewall/filter")
	return &res, err
}

func (n FiltersImpl) UpdateFilterRule(input *UpdateFilterRuleInput) (*UpdateFilterRuleOutput, error) {
	var res UpdateFilterRuleOutput
	_, err := n.Client.R().
		SetResult(&res).
		SetBody(input).
		Patch(fmt.Sprintf("/ip/firewall/filter/%s", input.Id))
	return &res, err
}

func (n FiltersImpl) DeleteFilterRule(input *DeleteFilterRuleInput) (*DeleteFilterRuleOutput, error) {
	_, err := n.Client.R().
		Delete(fmt.Sprintf("/ip/firewall/filter/%s", input.Id))
	return nil, err
}
