package go_mtpay

import (
	"fmt"
	"github.com/asaka1234/go-mtpay/utils"
)

// https://docs.mtpay.biz/zh/api/merchant-api/v2/webhook/
func (cli *Client) DepositCallback(req MTPayWebhookBackReq, processor func(MTPayWebhookBackReq) error) error {
	//验证签名
	sign := req.Signature //收到的签名
	timestamp := req.Timestamp
	verify := utils.VerifySign(cli.Params.AccessKey, cli.Params.SecretKey, timestamp, sign)

	if !verify {
		return fmt.Errorf("sign verify failed")
	}

	if req.Data.TradeType != "Deposit" {
		return fmt.Errorf("trade type error")
	}

	//开始处理
	return processor(req)
}
