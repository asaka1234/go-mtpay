package go_mtpay

type MTPayInitParams struct {
	MerchantId string `json:"merchantId" mapstructure:"merchantId" config:"merchantId" yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey" yaml:"accessKey"`
	SecretKey  string `json:"secretKey" mapstructure:"secretKey" config:"secretKey" yaml:"secretKey"`

	DepositUrl  string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl" yaml:"depositUrl"`
	WithdrawUrl string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl" yaml:"withdrawUrl"`

	DepositCallbackUrl  string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl" yaml:"depositBackUrl"`
	WithdrawCallbackUrl string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl" yaml:"withdrawBackUrl"`
}

// ---------------------------------------------
type MTPayDepositReq struct {
	Client          MTPayClient `json:"client" mapstructure:"client"`
	DepositCurrency string      `json:"depositCurrency" mapstructure:"depositCurrency"` //存款的货币类型（例如 MTC/CNY/HKD）
	FiatCurrency    string      `json:"fiatCurrency" mapstructure:"fiatCurrency"`       //法定货币代码（例如 CNY/HKD）
	DepositAmount   float64     `json:"depositAmount" mapstructure:"depositAmount"`     //存款金额
	MerchantOrderNo string      `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //唯一的商户订单号
	//以下三个让sdk来设置
	//WebhookURL    string             `json:"webhookUrl" mapstructure:"webhookUrl"` //通知的 Webhook URL
	//Language      string             `json:"language" mapstructure:"language"`     //en,zh_CN,zh_TW 默认是en
	//PaymentMethod MTPayPaymentMethod `json:"paymentMethod" mapstructure:"paymentMethod"`
}

type MTPayClient struct {
	RealName string `json:"realName" mapstructure:"realName"` // //客户的真实姓名
	//RegisteredAt int64  `json:"registeredAt"`
}

type MTPayPaymentMethod struct {
	//写死7
	Method int `json:"method" mapstructure:"method"` //位枚举,银行卡 → 1,微信支付 → 2,支付宝 → 4. 每个值表示一种支付方式，可以通过位运算组合多个值
}

//-------------------------------

type MTPayDepositResponse struct {
	Data       MTPayDepositResponseData `json:"data" mapstructure:"data"`
	IsSuccess  bool                     `json:"isSuccess" mapstructure:"isSuccess"`   //true/false
	StatusCode string                   `json:"statusCode" mapstructure:"statusCode"` //SUCCESS成功
	Message    string                   `json:"message" mapstructure:"message"`
	Version    string                   `json:"version" mapstructure:"version"`
}

type MTPayDepositResponseData struct {
	CheckoutURL string `json:"checkoutUrl" mapstructure:"checkoutUrl"` //支付url
	RequestCode string `json:"requestCode" mapstructure:"requestCode"` //请求编号 (可以把这个认为是psp的orderId)
	IsAvailable bool   `json:"isAvailable" mapstructure:"isAvailable"` //存款是否可用
	ExpiresAt   int64  `json:"expiresAt" mapstructure:"expiresAt"`     //到期时间戳（毫秒）
	CreatedAt   int64  `json:"createdAt" mapstructure:"createdAt"`     //创建时间戳（毫秒）
}

//==============================================

type MTPayWithdrawReq struct {
	Client           MTPayClient `json:"client" mapstructure:"client"`
	WithdrawCurrency string      `json:"withdrawCurrency" mapstructure:"withdrawCurrency"`
	FiatCurrency     string      `json:"fiatCurrency" mapstructure:"fiatCurrency"`
	WithdrawAmount   float64     `json:"withdrawAmount" mapstructure:"withdrawAmount"`
	BankCard         BankCard    `json:"bankCard" mapstructure:"bankCard"`               //客户用于接收资金的银行卡信息
	MerchantOrderNo  string      `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商户唯一订单号
	//以下让sdk设置
	//RequiresReview bool   `json:"requiresReview" mapstructure:"requiresReview"` //此笔提现是否需要人工审核(默认否)
	//WebhookURL     string `json:"webhookUrl" mapstructure:"webhookUrl"`         //通知用的 Webhook 地址
}

type BankCard struct {
	CardNumber     string `json:"cardNumber" mapstructure:"cardNumber"`         //客户银行卡号
	BankName       string `json:"bankName" mapstructure:"bankName"`             //客户银行名称
	BankBranchName string `json:"bankBranchName" mapstructure:"bankBranchName"` //客户银行分支机构名称
}

type MTPayWithdrawResp struct {
	Data       ResponseData `json:"data" mapstructure:"data"`
	IsSuccess  bool         `json:"isSuccess" mapstructure:"isSuccess"`
	StatusCode string       `json:"statusCode" mapstructure:"statusCode"`
	Message    string       `json:"message" mapstructure:"message"`
	Version    string       `json:"version" mapstructure:"version"`
}

type ResponseData struct {
	RequestCode string `json:"requestCode" mapstructure:"requestCode"` //请求追踪订单号 (psp的订单号)
	CreatedAt   int64  `json:"createdAt" mapstructure:"createdAt"`     //创建时间戳（毫秒
}

//===============callback===============================

type MTPayWebhookBackReq struct {
	Signature string    `json:"signature" mapstructure:"signature"` //签名，要验证（签名验证+时间戳检查）
	Timestamp int64     `json:"timestamp" mapstructure:"timestamp"`
	Data      TradeData `json:"data" mapstructure:"data"`
}

type TradeData struct {
	TradeType         string  `json:"tradeType" mapstructure:"tradeType"`             //交易类型：Deposit（存款）或 Withdraw（提现）
	RequestCode       string  `json:"requestCode" mapstructure:"requestCode"`         //psp系统生成的唯一请求编号
	MerchantOrderNo   string  `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商户自定义的唯一订单编号
	InternalOrderNo   string  `json:"internalOrderNo" mapstructure:"internalOrderNo"` //psp内部的内部订单编号
	Source            string  `json:"source" mapstructure:"source"`                   //请求来源：MerchantBackend、API_V1、API_V2
	ClientName        string  `json:"clientName" mapstructure:"clientName"`           //客户姓名
	RequestAmount     float64 `json:"requestAmount" mapstructure:"requestAmount"`     //请求的交易金额
	RequestCurrency   string  `json:"requestCurrency" mapstructure:"requestCurrency"` //易币种（如 MTC、CNY、HKD）
	Message           string  `json:"message" mapstructure:"message"`
	PaymentMethod     string  `json:"paymentMethod" mapstructure:"paymentMethod"` //支付方式（如适用）：BankCard、WeChatPay 等
	RequestStatus     string  `json:"requestStatus" mapstructure:"requestStatus"` // 当前交易状态: Finished,Failed,Cancelled,InProgress
	UnitPrice         float64 `json:"unitPrice" mapstructure:"unitPrice"`
	TransactionAmount float64 `json:"transactionAmount" mapstructure:"transactionAmount"`
	ReceivedAmount    float64 `json:"receivedAmount" mapstructure:"receivedAmount"`
	TransactionFee    float64 `json:"transactionFee" mapstructure:"transactionFee"`
	PaymentAmount     float64 `json:"paymentAmount" mapstructure:"paymentAmount"`
	FiatCurrency      string  `json:"fiatCurrency" mapstructure:"fiatCurrency"`
}

// 给callback的response
type MTPayWebhookBackResp struct {
	Success bool `json:"success" mapstructure:"success"` // true是成功
}
