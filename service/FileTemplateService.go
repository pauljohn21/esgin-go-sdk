package Service

import (
	"encoding/json"
	"log"

	Tools "github.com/pauljohn21/esgin-go-sdk/utils"

	Model "github.com/pauljohn21/esgin-go-sdk/model"
)

func GetFileUploadUrl(data Model.FileUploadUrlInfo) Tools.Res[Model.FileTemplatereq] {
	apiUrl := "/v3/files/file-upload-url"
	log.Println("获取文件上传地址：--------------")
	var dataJsonStr string
	if data, err := json.Marshal(data); err == nil {
		dataJsonStr = string(data)
	}
	initResult, err := Tools.SendCommHttp[Model.FileTemplatereq](apiUrl, dataJsonStr, "POST")
	log.Println("返回参数：------------------")
	log.Println(initResult)
	log.Println("错误信息：-----------------------")
	log.Println(err)
	return initResult
}
