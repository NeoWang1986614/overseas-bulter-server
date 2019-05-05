package wx

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"time"
	Error "overseas-bulter-server/error"
	storage "overseas-bulter-server/storage"
)

const (
	public_account_app_id = "wx2bcf56cbf31ff7b1"
	public_account_app_secret = "3ffff3d4d669083dfec39556180a8e67"
	get_access_token_url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

	get_public_account_material = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s"
	get_public_account_material_detail = "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s"
)

type WxError struct {
	ErrCode 	uint `json:"errcode"`
	ErrMsg 		uint `json:"errmsg"` 
}

type GetAccessTokenResult struct {
	AccessToken 		string	`json:"access_token"` 
	ExpiresIn 			uint  	`json:"expires_in"` 
}

type GetPublicMaterialRequest struct {
	Type 				string 	`json:"type"`
	Offset 				uint	`json:"offset"` 
	Count 				uint  	`json:"count"` 
}

type GetPublicMaterialDetailRequest struct {
	MediaId 			string 	`json:"media_id"`
}

type WxPublicAccountNewsItem struct {
	Title 					string 	`json:"title"`
	Author 					string	`json:"author"` 
	Digest 					string  `json:"digest"` 
	Content 				string  `json:"content"`
	NeedOpenComment 		uint  	`json:"need_open_comment"`
	OnlyFansCanComment 		uint  	`json:"only_fans_can_comment"` 
}
type WxPublicAccountContentNewsItem struct {
	NewsItem 				[]WxPublicAccountNewsItem 	`json:"news_item"`
	CreateTime				uint 						`json:"create_time"`
	UpdateTime				uint 						`json:"update_time"`
}
type WxPublicAccountMaterialItem struct {
	MediaId 				string 							`json:"media_id"`
	Content 				WxPublicAccountContentNewsItem 	`json:"content"`
	UpdateTime				uint 							`json:"update_time"`
}
type WxPublicAccountMaterial struct {
	Item 				 	[]WxPublicAccountMaterialItem 	`json:"item"`
	TotalCount 				uint 							`json:"total_count"`
	ItemCount				uint 							`json:"item_count"`
}

type WxPublicAccountMaterialDetailNewsItem struct {
	Title				string 	`json:"title"`
	Author				string 	`json:"author"`
	Digest				string 	`json:"digest"`
	Content				string 	`json:"content"`
	ContentSourceUrl	string 	`json:"content_source_url"`
	ThumbMediaId		string 	`json:"thumb_media_id"`
	ShowCoverPic		uint 	`json:"show_cover_pic"`
	Url					string 	`json:"url"`
	ThumbUrl			string 	`json:"thumb_url"`
	NeedOpenComment		uint 	`json:"need_open_comment"`
	OnlyFansCanComment	uint 	`json:"only_fans_can_comment"`
}

type WxPublicAccountMaterialDetail struct {
	NewsItem			[]WxPublicAccountMaterialDetailNewsItem `json:"news_item"`
	CreateTime			uint 	`json:"create_time"`
	UpdateTime			uint 	`json:"update_time"`
}

func checkValidAndGetAccessToken() string{
	var isNeedUpdateAccessTokenFromRemote = false;

	wechats := storage.QueryWechatAll()
	if(0 == len(wechats)){
		/*没有记录*/
		fmt.Println("-- 没有记录")
		isNeedUpdateAccessTokenFromRemote = true;
	}else if(isAccessTokenTimeout(&wechats[0])){
		/*已过期*/
		fmt.Println("-- 已过期")
		isNeedUpdateAccessTokenFromRemote = true;
	}
	
	if(isNeedUpdateAccessTokenFromRemote){
		ret := getAccessToken()
		if(0 != len(ret.AccessToken)){
			storage.AddWechat(ret.AccessToken, ret.ExpiresIn)
			return ret.AccessToken
		}
		fmt.Println("error: Access Token is empty")
		return ""
	}
	
	return wechats[0].AccessToken

}

func isAccessTokenTimeout(wechat *storage.DbWechat) bool{
	if(nil == wechat){
		panic("wechat is invalid")
	}

	updateTime, err := time.Parse("2006-01-02 15:04:05", wechat.UpdateTime) 
	Error.CheckErr(err)
	// fmt.Println("update time: ",updateTime)
	
	nowTime := time.Now().UTC().Add(time.Duration(8) * time.Hour)
	//

	// fmt.Println("now time: ", nowTime)
	durationHours := nowTime.Sub(updateTime).Hours()
	durationMinutes := nowTime.Sub(updateTime).Minutes()

	fmt.Println("duration hours = ", durationHours)
	fmt.Println("duration minutes = ", durationMinutes)
	
	return uint(durationHours) >= (wechat.ExpiresIn / 3600);
}

func getAccessToken() *GetAccessTokenResult{
	url := fmt.Sprintf(get_access_token_url, public_account_app_id, public_account_app_secret)
	fmt.Println("** get access token url: ",url)
	resp, err:= http.Get(url);
	Error.CheckErr(err)
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    Error.CheckErr(err)
	fmt.Println("** get access token response body: ",string(body))
	ret := &GetAccessTokenResult{}
	json.Unmarshal(body, ret)
	return ret
}

func GetPublicAccountMaterail(mType string, offset, count uint) *WxPublicAccountMaterial{

	url := fmt.Sprintf(get_public_account_material, checkValidAndGetAccessToken())
	fmt.Println(url)

	reqData := &GetPublicMaterialRequest{
		Type: mType,
		Offset: offset,
		Count: count}

	dataByte, err := json.Marshal(reqData)
	Error.CheckErr(err)
	fmt.Print(string(dataByte))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataByte))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // fmt.Println("response Status:", resp.Status)
    // fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
	
	ret := &WxPublicAccountMaterial{}
	wxErr := &WxError{}
	err = json.Unmarshal(body, &ret)
	Error.CheckErr(err)
	if(nil != err){
		err = json.Unmarshal(body, &wxErr)
	}
	
	return ret
}

func GetPublicAccountMaterailDetail(mediaId string) *WxPublicAccountMaterialDetail{

	url := fmt.Sprintf(get_public_account_material_detail, checkValidAndGetAccessToken())
	fmt.Println(url)

	reqData := &GetPublicMaterialDetailRequest{MediaId: mediaId}

	dataByte, err := json.Marshal(reqData)
	Error.CheckErr(err)
	fmt.Print(string(dataByte))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataByte))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

	ret := &WxPublicAccountMaterialDetail{}
	err = json.Unmarshal(body, &ret)
	
	return ret
}