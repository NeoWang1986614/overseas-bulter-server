package wx

import(
	"fmt"
	// "encoding/base64"
	// "encoding/json"
	// "crypto/aes"
	// "crypto/cipher"
	"errors"
	"time"
	entity "overseas-bulter-server/entity"
	storage "overseas-bulter-server/storage"
	Error "overseas-bulter-server/error"
	wxpay "github.com/objcoding/wxpay"
)

const (
	appId = "wx69af257300e856ce"
	mchId = "1534099671"
	apiKey = "kflskdjfklsdjfksdfflj34234234234"
)

func queryWxOpenIdByUserId(userId string) (string, error){
	dbUser := storage.QueryUserByUid(userId)
	if(0 == len(dbUser.WxOpenId)){
		return "", errors.New("err: No WxOpenId Found for userId = " + userId)
	}
	return dbUser.WxOpenId, nil
}

func getWxPayClient() *wxpay.Client{
	account := wxpay.NewAccount(appId, mchId, apiKey, false)
	fmt.Println("account1 = ", account)
	account.SetCertData("./apiclient_cert.pem")

	// 新建微信支付客户端
	client := wxpay.NewClient(account)

	return client
}

func Prepayment(userId, orderId string, totalFee int64) (*entity.PrepaymentResult, error){

	wxOpenId, err := queryWxOpenIdByUserId(userId)
	Error.CheckErr(err)
	fmt.Println("query wxOpenId = ", wxOpenId)

	client := getWxPayClient();

	existParams, isNeedNewPayment := queryWxPayment(client, orderId);

	params := make(wxpay.Params)
	if(isNeedNewPayment){
		fmt.Println("订单重新生成")
		params.SetString("body", "test").
		SetString("out_trade_no", orderId).
		SetInt64("total_fee", totalFee).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("openid", wxOpenId).
		SetString("notify_url", "https://bulter.mroom.com.cn:8008/overseas-bulter/v1/wxpay/notify").
		SetString("trade_type", "JSAPI")
		fmt.Println("prepayment params = ", params)
		newParams, err := client.UnifiedOrder(params)
		Error.CheckErr(err)
		fmt.Println("params = ", newParams);
		storage.UpdateOrderWxPrepayIdByUid(newParams["prepay_id"], orderId)
		params = newParams
	}else{
		fmt.Println("订单使用未支付订单")
		dbOrder := storage.QueryOrder(orderId)
		if(nil == dbOrder){
			panic("not found order when prepayment")
		}
		existParams.SetString("prepay_id", dbOrder.WxPrepayId)
		params = existParams;
	}

	if("SUCCESS" != params["result_code"]){
		return nil, errors.New("预支付下单错误!")
	}

	retObject := &entity.PrepaymentResult{
		NonceStr: params["nonce_str"],
		PrepayId: params["prepay_id"],
		SignType: "MD5",
		Sign: params["sign"],
		Timestamp: time.Now().Unix(),
		AppId: appId}
	

	return retObject, nil
	
}

func queryWxPayment(c *wxpay.Client, orderId string) (wxpay.Params, bool) {
//支付订单查询
	fmt.Println("queryWxPayment out_trade_no =  ", orderId)
	params := make(wxpay.Params)
	params.SetString("out_trade_no", orderId)
	p, err := c.OrderQuery(params)
	Error.CheckErr(err)
	fmt.Println("p = ", p);
	if("ORDERNOTEXIST" == p["err_code"]){
		return p, true
	}
	if("NOTPAY" == p["trade_state"]){
		return p, false
	}
	panic("order has paied，重复支付已支付")
}

func CloseWxPayment(orderId string) {
	c := getWxPayClient();
	fmt.Println("close wx payment out_trade_no =  ", orderId)
	params := make(wxpay.Params)
	params.SetString("out_trade_no", orderId)
	p, err := c.CloseOrder(params)
	Error.CheckErr(err)
	fmt.Println("p = ", p);
	if("SUCCESS" == p["return_code"]){
		fmt.Println("关闭订单成功!")
	}else{
		fmt.Println("关闭订单失败!")
	}
}