package go_mtpay

import (
	"github.com/asaka1234/go-mtpay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params MTPayInitParams

	ryClient   *resty.Client
	debugModel bool //是否调试模式
	logger     utils.Logger
}

func NewClient(logger utils.Logger, params MTPayInitParams) *Client {
	return &Client{
		Params: params,

		ryClient:   resty.New(), //client实例
		debugModel: false,
		logger:     logger,
	}
}

func (cli *Client) SetDebugModel(debugModel bool) {
	cli.debugModel = debugModel
}
