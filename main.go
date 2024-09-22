package main

import (
	"encoding/json"
	"fmt"

	Model "XuanYuanAPI-Golang/model"
	Service "XuanYuanAPI-Golang/service"
	Tools "XuanYuanAPI-Golang/utils"
)

func main() {
	// 配置参数(作为全局执行一次即可)
	Service.EsignInitService()

	// 获取文件上传地址-start
	filePath := "table.docx"
	contentMd5, size := Tools.CountFileMd5(filePath)
	var getFileUploadUrlInfo Model.GetFileUploadUrlInfo
	getFileUploadUrlInfo.ContentMd5 = contentMd5
	getFileUploadUrlInfo.ContentType = "application/octet-stream"
	getFileUploadUrlInfo.ConvertToPDF = true
	getFileUploadUrlInfo.FileName = "table.docx"
	getFileUploadUrlInfo.FileSize = size
	var getFileUploadUrlInfoJson string
	if data, err := json.Marshal(getFileUploadUrlInfo); err == nil {
		getFileUploadUrlInfoJson = string(data)
	}
	initResult := Service.GetFileUploadUrl(getFileUploadUrlInfoJson)
	// getFileUploadUrlData := Tools.BytetoJson(initResult)["data"]
	// getFileUploadUrlDataMap := getFileUploadUrlData.(map[string]interface{})
	// fileId := getFileUploadUrlDataMap["fileId"]
	// uploadUrl := getFileUploadUrlDataMap["fileUploadUrl"]
	fmt.Println(initResult.Data.FileId)
	fmt.Println(initResult.Data.FileUploadUrl)

	// 上传文件-start
	result := Tools.UpLoadFile(initResult.Data.FileUploadUrl, filePath, contentMd5, "application/octet-stream")
	fmt.Println(result)
	// 上传文件-end
}
