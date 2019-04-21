package wx

import(
	"fmt"
	"encoding/base64"
	"encoding/json"
	"crypto/aes"
	"crypto/cipher"
	entity "overseas-bulter-server/entity"
	Error "overseas-bulter-server/error"
)



func DecodeShareInfo(sessionKey, encryptData, iv string) *entity.ShareInfoOpenData {
	decrypData,err := decryptWXOpenData(sessionKey, encryptData, iv)
	Error.CheckErr(err);

	ret := &entity.ShareInfoOpenData{}

	err = json.Unmarshal(decrypData, ret)
	Error.CheckErr(err)
	fmt.Print(ret)
	return ret;
}

func decryptWXOpenData(sessionKey, encryptData, iv string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	dataBytes, err := aesDecrypt(decodeBytes, sessionKeyBytes, ivBytes)

	fmt.Println(" aes open data=", string(dataBytes));

	return dataBytes, nil

}

func aesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x0e' || ch == '\x03' {
			origData[i] = ' '
		}
	}
	//{"phoneNumber":"15082726017","purePhoneNumber":"15082726017","countryCode":"86","watermark":{"timestamp":1539657521,"appid":"wx4c6c3ed14736228c"}}//<nil>
	return origData, nil
}