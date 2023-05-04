package http

import (
	"crypto/tls"
	"kaspi-qr/internal/adapters/notifier"
	"net/http"
	"time"

	"github.com/rendau/dop/dopErrs"

	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/client/httpc/httpclient"

	"github.com/rendau/dop/adapters/logger"
)

const (
	RequestTimeout = 10 * time.Second
)

type St struct {
	lg logger.Lite
}

func New(lg logger.Lite) *St {
	return &St{
		lg: lg,
	}
}

func (o *St) NotifyOrderStatusChange(uri string, obj *notifier.OrderStatusChangeReqSt) error {
	if uri == "" {
		o.lg.Infow("NotifyOrderStatusChange", "uri", uri, "obj", obj)
		return nil
	}

	httpClient := httpclient.New(o.lg, &httpc.OptionsSt{
		Client: &http.Client{
			Timeout:   RequestTimeout,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
		LogPrefix: "NotifyOrderStatusChange: ",
	})

	_, err := httpClient.Send(&httpc.OptionsSt{
		Uri:    uri,
		Method: "POST",
		//LogFlags:  httpc.LogResponse,

		ReqObj: obj,
	})
	if err != nil {
		return dopErrs.ServiceNA
	}

	return nil
}
