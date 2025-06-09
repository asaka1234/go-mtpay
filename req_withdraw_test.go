package go_mtpay

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &MTPayInitParams{MERCHANT_ID, ACCESS_KEY, SECRET_KEY, DEPOST_URL, WITHDRAW_URL, DEPOST_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() MTPayWithdrawReq {
	return MTPayWithdrawReq{
		Client: MTPayClient{
			RealName: "cy", //客户信息
		},
		WithdrawCurrency: "CNY",
		FiatCurrency:     "CNY",
		WithdrawAmount:   200.00,
		MerchantOrderNo:  "793793793",
		BankCard: BankCard{
			CardNumber:     "123",
			BankName:       "工行",
			BankBranchName: "解放路支行",
		},
	}
}
