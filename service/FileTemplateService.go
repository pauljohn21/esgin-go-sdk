package Service

import (
	"log"

	Tools "XuanYuanAPI-Golang/utils"
)

func GetFileUploadUrl(dataJsonStr string) Tools.Res[FileTemplatereq] {
	apiUrl := "/v3/files/file-upload-url"
	log.Println("获取文件上传地址：--------------")
	initResult, err := Tools.SendCommHttp[FileTemplatereq](apiUrl, dataJsonStr, "POST")
	log.Println("返回参数：------------------")
	log.Println(initResult)
	log.Println("错误信息：-----------------------")
	log.Println(err)
	return initResult
}

type FileTemplatereq struct {
	FileId        string `json:"fileId"`
	FileUploadUrl string `json:"fileUploadUrl"`
}
