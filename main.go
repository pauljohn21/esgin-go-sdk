package main

import (
	"fmt"
	"log"

	Model "esgin-go-sdk/model"
	Service "esgin-go-sdk/service"
	Tools "esgin-go-sdk/utils"
)

func main() {
	// 配置参数(作为全局执行一次即可)
	Service.EsignInitService()

	// 获取文件上传地址-start
	filePath := "table.docx"
	contentMd5, size := Tools.CountFileMd5(filePath)
	fileUploadUrlInfo := Model.FileUploadUrlInfo{
		ContentMd5:   contentMd5,
		ContentType:  "application/octet-stream",
		ConvertToPDF: true,
		FileName:     "table.docx",
		FileSize:     size,
	}

	initResult := Service.GetFileUploadUrl(fileUploadUrlInfo)

	log.Println("文件id: " + initResult.Data.FileId)

	// 上传文件-start
	result := Tools.UpLoadFile(initResult.Data.FileUploadUrl, filePath, contentMd5, "application/octet-stream")
	fmt.Println(result)
	// 上传文件-end
}
