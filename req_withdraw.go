package go_mtpay

import (
	"crypto/tls"
	"github.com/asaka1234/go-mtpay/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// https://docs.mtpay.biz/zh/api/merchant-api/v2/merchant-withdraw-bankcard-api/
func (cli *Client) Withdraw(req MTPayWithdrawReq) (*MTPayWithdrawResp, error) {

	rawURL := cli.Params.WithdrawUrl

	// 1. 拿到请求参数，转为map
	var signDataMap map[string]interface{}
	mapstructure.Decode(req, &signDataMap)
	signDataMap["requiresReview"] = false
	signDataMap["webhookUrl"] = cli.Params.WithdrawCallbackUrl

	// 2. 计算签名,补充参数
	timestamp := time.Now().UTC().UnixMilli()
	signStr, timestamp, _ := utils.GenSign(cli.Params.AccessKey, cli.Params.SecretKey, timestamp)

	var result MTPayWithdrawResp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(signDataMap).
		SetHeaders(getHeaders(cli.Params.AccessKey, signStr, timestamp)).
		SetDebug(cli.debugModel).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
