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

func SendHttp[T any](apiUrl string, data string, method string, headers map[string]string) (*Res[T], error) {
	// API接口返回值
	var res Res[T]
	url := apiUrl
	jsonStr := []byte(data)
	var req *http.Request
	var err error
	if method == "GET" || method == "DELETE" {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			// 根据实际情况选择合适的错误处理方式
			log.Println("Failed to create HTTP request:", err)
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
		if err != nil {
			// 根据实际情况选择合适的错误处理方式
			log.Println("Failed to create HTTP request:", err)
			return nil, err
		}
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// 文件上传
func UpLoadFile(uploadUrl string, filePath string, contentMD5 string, contentType string) string {
	fmt.Printf("开始上传uploadUrl %s\n", uploadUrl)
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

func SendCommHttp[T any](apiUrl string, dataJsonStr string, method string) (*Res[T], error) {
	log.Println("请求参数JSON字符串:" + dataJsonStr)
	httpUrl := InstaneEsignInitConfig().host + apiUrl
	log.Println("发送地址: " + httpUrl)
	md5Str := DohashMd5(dataJsonStr)
	message := AppendSignDataString(method, "*/*", md5Str, "application/json; charset=UTF-8", "", "", apiUrl)
	reqSignature := DoSignatureBase64(message, InstaneEsignInitConfig().projectScert)
	// 初始化接口返回值
	initResult, err := SendHttp[T](httpUrl, dataJsonStr, method, buildCommHeader(md5Str, reqSignature))
	return initResult, err
}

func buildCommHeader(contentMD5 string, reqSignature string) (header map[string]string) {
	headers := map[string]string{}
	headers["X-Tsign-Open-App-Id"] = InstaneEsignInitConfig().projectId
	headers["X-Tsign-Open-Ca-Timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	headers["Accept"] = "*/*"
	headers["X-Tsign-Open-Ca-Signature"] = reqSignature
	headers["Content-MD5"] = contentMD5
	headers["Content-Type"] = "application/json; charset=UTF-8"
	headers["X-Tsign-Open-Auth-Mode"] = "Signature"
	return headers
}
