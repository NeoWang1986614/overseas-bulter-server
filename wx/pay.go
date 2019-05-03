package wx

import(
	"fmt"
	// "encoding/base64"
	// "encoding/json"
	// "crypto/aes"
	// "crypto/cipher"
	// entity "overseas-bulter-server/entity"
	// Error "overseas-bulter-server/error"
	wxPay "github.com/objcoding/wxpay"
)

func CreateAccount() {
	account1 := wxPay.NewAccount("appid", "mchid", "apiKey", false)
	fmt.Println("account1 = ", account1)
}