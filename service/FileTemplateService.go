package Service

import (
	"log"

	Model "XuanYuanAPI-Golang/model"
	Tools "XuanYuanAPI-Golang/utils"
)

func GetFileUploadUrl(dataJsonStr string) Tools.Res[Model.FileTemplatereq] {
	apiUrl := "/v3/files/file-upload-url"
	log.Println("获取文件上传地址：--------------")
	initResult, err := Tools.SendCommHttp[Model.FileTemplatereq](apiUrl, dataJsonStr, "POST")
	log.Println("返回参数：------------------")
	log.Println(initResult)
	log.Println("错误信息：-----------------------")
	log.Println(err)
	return initResult
}
