package go_mtpay

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, MTPayInitParams{MERCHANT_ID, ACCESS_KEY, SECRET_KEY, DEPOST_URL, WITHDRAW_URL, DEPOST_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() MTPayDepositReq {
	return MTPayDepositReq{
		Client: MTPayClient{
			RealName: "cy", //客户信息
		},
		DepositCurrency: "CNY",
		FiatCurrency:    "CNY",
		DepositAmount:   200.00,
		MerchantOrderNo: "793793793",
	}
}
