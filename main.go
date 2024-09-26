package main

import (
	"fmt"
	"log"

	Model "github.com/pauljohn21/esgin-go-sdk/model"
	Service "github.com/pauljohn21/esgin-go-sdk/service"
	Tools "github.com/pauljohn21/esgin-go-sdk/utils"
)

func main() {
	// 配置参数(作为全局执行一次即可)
	EsignInit()

	// 获取文件上传地址-start
	filePath := "./file/demo.docx"
	contentMd5, size := Tools.CountFileMd5(filePath)
	fileUploadUrlInfo := Model.FileUploadUrlInfo{
		ContentMd5:   contentMd5,
		ContentType:  "application/octet-stream",
		ConvertToPDF: true,
		FileName:     "测试0924.docx",
		FileSize:     size,
	}

	initResult := Service.GetFileUploadUrl(fileUploadUrlInfo)

	log.Println("文件id: " + initResult.Data.FileId)

	// 上传文件-start
	result := Tools.UpLoadFile(initResult.Data.FileUploadUrl, filePath, contentMd5, "application/octet-stream")
	fmt.Println(result)
	// 上传文件-end
}

const (
	host         = "https://smlopenapi.esign.cn"
	projectId    = "7439012567"
	projectScert = "c0d5ea7e936a53d515938cd5dabcba37"
)

func EsignInit() {
	config := Tools.InstaneEsignInitConfig()
	config.SetHost(host)
	config.SetProjectId(projectId)
	config.SetProjectScert(projectScert)
}
