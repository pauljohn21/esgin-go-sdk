package service

import (
	"log"

	Tools "esgin-go-sdk/tools"
)

func DocTemplatesList() *Tools.Res[DocTemplatesListRes] {
	apiurl := "/v3/doc-templates?pageNum=1&pageSize=20"
	result, err := Tools.SendCommHttp[DocTemplatesListRes](apiurl, "", "GET")
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

type DocTemplatesListRes struct {
	Total        int32                             `json:"total"`
	DocTemplates []DocTemplatesListResdocTemplates `json:"docTemplates"`
}
type DocTemplatesListResdocTemplates struct {
	DocTemplateId   string `json:"docTemplateId"`
	DocTemplateName string `json:"docTemplateName"`
	CreateTime      int64  `json:"createTime"`
	UpdateTime      int64  `json:"updateTime"`
}
