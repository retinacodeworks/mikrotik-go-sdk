package routeros

import (
	"crypto/tls"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/retinacodeworks/mikrotik-go-sdk/pkg/interfaces"
	"github.com/retinacodeworks/mikrotik-go-sdk/pkg/ipv4"
	"github.com/retinacodeworks/mikrotik-go-sdk/pkg/routing"
	"net/http"
)

type Options struct {
	Address string
	Headers map[string]string
	Tls     *tls.Config
	Debug   bool
}

type RouterOS struct {
	Client *resty.Client
}

func New(opts *Options) *RouterOS {
	client := resty.New()
	client.SetBaseURL(opts.Address)
	client.SetDebug(opts.Debug)
	//client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
	//	// Debug flag
	//	fmt.Println(request.URL)
	//	return nil
	//})
	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		switch response.StatusCode() {
		case http.StatusInternalServerError:
			return errors.New("an unexpected error occured")
		case http.StatusBadRequest:
			return errors.New(response.String())
		case http.StatusNotFound:
			return errors.New(response.String())
		default:
			return nil
		}
	})
	client.SetHeaders(opts.Headers)
	client.SetTLSClientConfig(opts.Tls)

	return &RouterOS{
		Client: client,
	}
}

type VersionInfoResponse struct {
	BoardName string `json:"board-name"`
	Version   string `json:"version"`
}

func (r *RouterOS) GetVersionInfo() (VersionInfoResponse, error) {
	var resp VersionInfoResponse
	_, err := r.Client.R().
		SetResult(&resp).
		Get("/rest/system/resource")
	return resp, err
}

func (r *RouterOS) Ipv4() ipv4.Ipv4 {
	return ipv4.Ipv4{Client: r.Client}
}

func (r *RouterOS) Routing() routing.Routing {
	return routing.Routing{Client: r.Client}
}

func (r *RouterOS) Interfaces() interfaces.Interfaces {
	return interfaces.InterfacesImpl{Client: r.Client}
}
