package pay

import (
	"fmt"
	"gameserver/utils"
	"github.com/objcoding/wxpay"
	"github.com/sirupsen/logrus"
)

// 统一下单
func UnifiedOrder()  {
	account := wxpay.NewAccount("wxa0a2cd667244d857", "1523222491", "c6e0f9b0c46bb7fa1d8bfa602ee3a81a",false) // sandbox环境请传true
	client := wxpay.NewClient(account)

	params := make(wxpay.Params)
	params.SetString("body", "test").
		SetString("out_trade_no", utils.UniqueId()).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://notify.objcoding.com/notify").
		SetString("nonce_str",utils.GetRandomString(32)).
		SetString("trade_type", "JSAPI").
		SetInt64("total_fee", 1)
	sign := client.Sign(params)
	fmt.Println(sign)

	// client.SetSignType(HMACSHA256)

	// 设置支付账户
	// client.SetAccount(account)
	order, err := client.UnifiedOrder(params)
	logrus.Info("error:",err)
	if err != nil {
		logrus.Info("error:",err)
	}
	code,ok := order["return_code"]
	if !ok || code != "SUCCESS" {
		logrus.Error(fmt.Sprintf("UnifiedOrder order:%v",order))
	}
	fmt.Println(order,err)
}

// 查询订单
func QueryOrder()  {
	params := make(wxpay.Params)

	params.SetString("out_trade_no", "3568785")
	account := wxpay.NewAccount("appid", "mchid", "apiKey",false)
	client := wxpay.NewClient(account)
	client.OrderQuery(params)
}