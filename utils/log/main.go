package log

import (
	"github.com/Dev-Addict/gm-common/types/payloads"
	"github.com/go-resty/resty/v2"
)

var client = resty.New()

func Log(payload payloads.LogPayload) error {
	_, err := client.R().SetBody(payload).Post("http://logger-service:8080/log")
	if err != nil {
		return err
	}

	return nil
}
