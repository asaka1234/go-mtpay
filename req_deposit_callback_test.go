package go_mtpay

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDepositCallback(t *testing.T) {

	jsonStr := `{
  "signature": "0582700D1CC8282B3EBFA7CFDF35854C30763E1C4541560DBCB67116F9EAC9E1",
  "timestamp": 1749111542621,
  "data": {
    "tradeType": "Deposit",
    "requestCode": "ef695e11e2c34573a08be7fba0b8f925",
    "merchantOrderNo": "202506051117290894",
    "internalOrderNo": "AT-D-4FXE5PQU3",
    "source": "API_V2",
    "clientName": "admin",
    "requestAmount": 1103,
    "requestCurrency": "CNY",
    "message": "",
    "paymentMethod": "BankCard",
    "requestStatus": "InProgress",
    "unitPrice": 7.28,
    "transactionAmount": 151.51,
    "receivedAmount": 146.21,
    "transactionFee": 5.3,
    "paymentAmount": 1103,
    "fiatCurrency": "CNY"
  },
  "latency": 13.473,
  "status_code": 404
}`

	var resp MTPayWebhookBackReq

	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		fmt.Printf("JSON 解析失败: %v", err)
		return
	}

	// 打印解析结果
	fmt.Printf("解析结果: %+v\n", resp)

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, MTPayInitParams{MERCHANT_ID, ACCESS_KEY, SECRET_KEY, DEPOST_URL, WITHDRAW_URL, DEPOST_BACK_URL, WITHDRAW_BACK_URL})
	err = cli.DepositCallback(resp, process)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

}

func process(req MTPayWebhookBackReq) error {
	return nil
}
