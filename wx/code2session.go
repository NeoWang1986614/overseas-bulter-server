package wx

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	Error "overseas-bulter-server/error"
)

var (
	getUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type WxSession struct {
	OpenId 		string
	Session_key 	string
}

func  Code2Session(code string, appId string, appSecret string)  *WxSession{
	fmt.Println(code)
	url := fmt.Sprintf(getUrl, appId, appSecret, code)
	fmt.Println(url)
	resp, err:= http.Get(url);
	Error.CheckErr(err)
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    Error.CheckErr(err)
	fmt.Println(string(body))
	session := &WxSession{}
	json.Unmarshal(body, session)
	// fmt.Println(session.OpenId)
	// fmt.Println(session.SessionKey)
	return session
}