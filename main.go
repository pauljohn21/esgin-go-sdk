package main

import (
	"fmt"

	"esgin-go-sdk/service"
	Tools "esgin-go-sdk/tools"
)

func main() {
	service.EsignInitService()
	fileid, md5Str := service.FileUpload()
	fmt.Println(md5Str)
	uploadurl := fileid.Data.FileUploadUrl
	fmt.Println(uploadurl)
	res := Tools.UpLoadFile(uploadurl, "table.docx", md5Str, "application/octet-stream")
	fmt.Println(res)

	// res := service.DocTemplatesList()
	// fmt.Println(res.Data)
}
