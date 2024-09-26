package Tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Res[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// ParseResponse 是一个假定存在的函数，用于将 byte 切片转换为目标类型 Res[T]。
func ParseResponse[T any](body []byte) (Res[T], error) {
	var res Res[T]
	if err := json.Unmarshal(body, &res); err != nil {
		return res, fmt.Errorf("无法解析响应体: %v", err)
	}
	return res, nil
}

// SendHttp 发送 HTTP 请求并返回结果。
// T 表示返回数据的类型，可以是任何实现了 json.Unmarshaler 接口的类型。
func SendHttp[T any](apiUrl string, data string, method string, headers map[string]string) (res Res[T], err error) {
	url := apiUrl
	jsonStr := []byte(data)
	var req *http.Request

	if method == "GET" || method == "DELETE" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	}
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// 尝试将响应体解析为目标类型 Res[T]
	res, err = ParseResponse[T](body)
	return
}

// 文件上传
func UpLoadFile(uploadUrl string, filePath string, contentMD5 string, contentType string) string {
	// 创建一个缓冲区对象,后面的要上传的body都存在这个缓冲区里
	bodyBuf := &bytes.Buffer{}
	// 要上传的文件
	// 创建第一个需要上传的文件,filepath.Base获取文件的名称
	// 打开文件
	fd1, _ := os.Open(filePath)
	defer fd1.Close()
	// 把第一个文件流写入到缓冲区里去
	_, _ = io.Copy(bodyBuf, fd1)
	// 获取请求Content-Type类型,后面有用
	// contentType := bodyWriter.FormDataContentType()
	// 创建一个http客户端请求对象
	client := &http.Client{}
	// 创建一个post请求
	req, _ := http.NewRequest("PUT", uploadUrl, nil)
	// 设置请求头
	req.Header.Set("Content-MD5", contentMD5)
	// 这里的Content-Type值就是上面contentType的值
	req.Header.Set("Content-Type", contentType)
	// 转换类型
	req.Body = io.NopCloser(bodyBuf)
	// 发送数据
	data, _ := client.Do(req)
	// 读取请求返回的数据
	bytes, _ := io.ReadAll(data.Body)
	defer data.Body.Close()
	// 返回数据
	return string(bytes)
}

func SendCommHttp[T any](apiUrl string, dataJsonStr string, method string) (Res[T], error) {
	log.Println("请求参数JSON字符串:" + dataJsonStr)
	httpUrl := InstaneEsignInitConfig().Host() + apiUrl
	log.Println("发送地址: " + httpUrl)
	md5Str := DohashMd5(dataJsonStr)
	message := AppendSignDataString(method, "*/*", md5Str, "application/json; charset=UTF-8", "", "", apiUrl)
	reqSignature := DoSignatureBase64(message, InstaneEsignInitConfig().ProjectScert())
	// 初始化接口返回值
	res, err := SendHttp[T](httpUrl, dataJsonStr, method, buildCommHeader(md5Str, reqSignature))
	return res, err
}

func buildCommHeader(contentMD5 string, reqSignature string) (header map[string]string) {
	headers := map[string]string{}
	headers["X-Tsign-Open-App-Id"] = InstaneEsignInitConfig().ProjectId()
	headers["X-Tsign-Open-Ca-Timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	headers["Accept"] = "*/*"
	headers["X-Tsign-Open-Ca-Signature"] = reqSignature
	headers["Content-MD5"] = contentMD5
	headers["Content-Type"] = "application/json; charset=UTF-8"
	headers["X-Tsign-Open-Auth-Mode"] = "Signature"
	return headers
}
