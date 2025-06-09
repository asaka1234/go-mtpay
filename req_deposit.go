package go_mtpay

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-mtpay/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// https://docs.mtpay.biz/zh/api/merchant-api/v2/merchant-deposit-api/
func (cli *Client) Deposit(req MTPayDepositReq) (*MTPayDepositResponse, error) {

	rawURL := cli.Params.DepositUrl

	// 1. 拿到请求参数，转为map
	var signDataMap map[string]interface{}
	mapstructure.Decode(req, &signDataMap)
	signDataMap["webhookUrl"] = cli.Params.DepositCallbackUrl
	signDataMap["language"] = "zh_CN" //先写死
	signDataMap["paymentMethod"] = map[string]interface{}{
		"method": 7, //写死
	}

	// 2. 计算签名,补充参数
	timestamp := time.Now().UTC().UnixMilli()
	signStr, timestamp, _ := utils.GenSign(cli.Params.AccessKey, cli.Params.SecretKey, timestamp)

	fmt.Printf("sign: %s\n", signStr)

	var result MTPayDepositResponse

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(signDataMap).
		SetHeaders(getHeaders(cli.Params.AccessKey, signStr, timestamp)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
