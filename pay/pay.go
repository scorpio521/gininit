package pay

import (
	"fmt"
	"github.com/iGoogle-ink/gopay"
	"math/rand"
	"strconv"
	"time"
	"gininit/tool"
)

func WeixinPay(subject string, out_trade_no string, total_fee int) (payData map[string]interface{}, err error) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	appId := ""
	mchID := ""
	apiKey := ""
	//初始化微信客户端
	client := gopay.NewWeChatClient(appId, mchID, apiKey, true)
	number := gopay.GetRandomString(32)
	//初始化参数Map
	body := make(gopay.BodyMap)
	body.Set("nonce_str", number)
	body.Set("body", subject)
	body.Set("out_trade_no", out_trade_no)
	body.Set("total_fee", total_fee)
	//body.Set("spbill_create_ip", ip)
	body.Set("notify_url", "http://www.gopay.ink")
	body.Set("trade_type", gopay.TradeType_App)
	body.Set("device_info", "WEB")
	body.Set("sign_type", gopay.SignType_MD5)
	//发起下单请求
	wxRsp, err := client.UnifiedOrder(body)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	paySign := gopay.GetAppPaySign(appId, mchID, wxRsp.NonceStr, wxRsp.PrepayId, gopay.SignType_MD5, timeStamp, apiKey)
	payData = make(map[string]interface{})
	payData["appid"] = appId
	payData["partnerid"] = mchID
	payData["prepayid"] = wxRsp.PrepayId
	payData["noncestr"] = wxRsp.NonceStr
	payData["timestamp"] = timeStamp
	payData["package"] = "Sign=WXPay"
	payData["sign"] = paySign
	err = nil
	fmt.Println(payData)
	return

	//APP支付需要的参数信息
	//fmt.Println("paySign：", paySign)
	//fmt.Println("ReturnCode：", wxRsp.ReturnCode)
	//fmt.Println("ReturnMsg：", wxRsp.ReturnMsg)
	//fmt.Println("Appid：", wxRsp.Appid)
	//fmt.Println("MchId：", wxRsp.MchId)
	//fmt.Println("DeviceInfo：", wxRsp.DeviceInfo)
	//fmt.Println("NonceStr：", wxRsp.NonceStr)
	//fmt.Println("Sign：", wxRsp.Sign)
	//fmt.Println("ResultCode：", wxRsp.ResultCode)
	//fmt.Println("ErrCode：", wxRsp.ErrCode)
	//fmt.Println("ErrCodeDes：", wxRsp.ErrCodeDes)
	//fmt.Println("PrepayId：", wxRsp.PrepayId)
	//fmt.Println("TradeType：", wxRsp.TradeType)
	//fmt.Println("CodeUrl:", wxRsp.CodeUrl)
	//fmt.Println("MwebUrl:", wxRsp.MwebUrl)
}
func WeixinPayCheck(out_trade_no string) (dataReturn map[string]interface{}, err error) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	appId := "wx31571b0e6a803035"
	mchID := "1535802401"
	apiKey := "d95ffc620424ae4a584aaac539b2f208"
	//初始化微信客户端
	client := gopay.NewWeChatClient(appId, mchID, apiKey, true)
	//初始化参数Map
	body := make(gopay.BodyMap)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("device_info", "WEB")
	body.Set("sign_type", gopay.SignType_MD5)
	body.Set("out_trade_no", out_trade_no)

	//发查询订单
	wxRsp, err := client.QueryOrder(body)
	dataReturn = tool.StructToMapViaJson(wxRsp)

	fmt.Println("dataReturn", dataReturn)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = nil
	return
}
func WeixinPayRefund(out_trade_no, out_refund_no string, refund_fee int) (dataReturn map[string]interface{}, err error) {
	//初始化微信客户端
	//    appId：应用ID
	//    mchID：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	appId := "wx31571b0e6a803035"
	mchID := "1535802401"
	apiKey := "d95ffc620424ae4a584aaac539b2f208"
	//初始化微信客户端
	client := gopay.NewWeChatClient(appId, mchID, apiKey, true)
	//初始化参数Map
	body := make(gopay.BodyMap)
	body.Set("nonce_str", gopay.GetRandomString(32))
	body.Set("sign_type", gopay.SignType_MD5)
	body.Set("out_refund_no", out_refund_no)
	body.Set("out_trade_no", out_trade_no)
	//body.Set("total_fee", total_fee)
	body.Set("refund_fee", refund_fee)

	//发起退款
	wxRsp, err := client.Refund(body, "/Users/xx/Documents/wxpay/apiclient_cert.pem", "/Users/xx/Documents/wxpay/apiclient_key.pem", "")
	dataReturn = tool.StructToMapViaJson(wxRsp)

	fmt.Println("dataReturn", dataReturn)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = nil
	return
}
func ALipay(subject, out_trade_no string, total_amount float64) (payData map[string]interface{}, err error) {

	privateKey := ""
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := gopay.NewAliPayClient("", privateKey, true)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		//SetAppAuthToken("").
		//SetReturnUrl("https://www.gopay.ink").
		SetNotifyUrl("https://www.gopay.ink")
	//请求参数
	body := make(gopay.BodyMap)
	body.Set("subject", subject)
	body.Set("out_trade_no", out_trade_no)
	body.Set("total_amount", total_amount)
	//手机APP支付参数请求
	payParam, err := client.AliPayTradeAppPay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payParam:", payParam)
	payData = make(map[string]interface{})
	err = nil
	payData["orderString"] = payParam

	return
}
func ALipayRefund(out_trade_no string, refund_amount float64, refund_reason string) (data interface{}, err error) {

	privateKey := ""//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := gopay.NewAliPayClient("", privateKey, true)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2")

	fmt.Println(refund_amount)
	//请求参数
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", out_trade_no)
	body.Set("refund_amount", refund_amount)
	body.Set("refund_reason", refund_reason)
	//手机APP支付参数请求
	payParam, err := client.AliPayTradeRefund(body)
	//fmt.Println("payParam:", payParam)
	dataReturn := tool.StructToMapViaJson(payParam)
	fmt.Println("dataReturn", dataReturn)

	data = dataReturn["alipay_trade_refund_response"]

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	return
}

func ALipayCheck(out_trade_no string) (data interface{}, err error) {

	privateKey := ""//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := gopay.NewAliPayClient("", privateKey, true)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2")
		//SetAppAuthToken("").
		//SetReturnUrl("https://www.gopay.ink").
		//SetNotifyUrl("https://www.gopay.ink")
	//请求参数
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", out_trade_no)
	//手机APP支付参数请求
	payParam, err := client.AliPayTradeQuery(body)
	//fmt.Println("payParam:", payParam)
	dataReturn := tool.StructToMapViaJson(payParam)

	data = dataReturn["alipay_trade_query_response"]

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	return
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func GetRandomStringInt(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
