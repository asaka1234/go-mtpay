package go_mtpay

import (
	"github.com/asaka1234/go-mtpay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params MTPayInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params MTPayInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
