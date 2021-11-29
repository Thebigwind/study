package wechat

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Reply struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default() //实例化一个gin

	router.POST("/getnum", getNum)

	fmt.Println("服务启动...端口为9300")
	router.Run("127.0.0.1:9300") //监听9300端口
}

func getNum(c *gin.Context) {
	var req struct {
		EncryptedData string
		Iv            string
		Code          string
	}
	err := c.Bind(&req)
	if err != nil {
		c.Error(err)
	}
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=" + req.Code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	if err != nil {
		c.Error(err)

	}
	defer resp.Body.Close()
	s, _ := ioutil.ReadAll(resp.Body)
	res := make(map[string]string)
	json.Unmarshal(s, &res)
	key, _ := base64.StdEncoding.DecodeString(res["session_key"])
	iv, _ := base64.StdEncoding.DecodeString(req.Iv)
	ciphertext, _ := base64.StdEncoding.DecodeString(req.EncryptedData)
	plaintext := make([]byte, len(ciphertext))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext = PKCS7UnPadding(plaintext)
	fmt.Println("return:", string(plaintext))
	c.JSON(http.StatusOK, Reply{http.StatusOK, string(plaintext)})
}

// 发送post请求
func post(url, data string) (string, error) {
	reader := bytes.NewReader([]byte(data))

	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", err
	}
	defer request.Body.Close() //程序在使用完回复后必须关闭回复的主体
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//必须设定该参数,POST参数才能正常提交，意思是以json串提交数据

	client := http.Client{}
	resp, err := client.Do(request) //Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		return "", err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBytes), nil
}

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
