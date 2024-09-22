package service

import Tools "esgin-go-sdk/tools"

const (
	host         = "https://smlopenapi.esign.cn"
	projectId    = "7439012567"
	projectScert = "c0d5ea7e936a53d515938cd5dabcba37"
)

func EsignInitService() {
	config := Tools.InstaneEsignInitConfig()
	config.SetHost(host)
	config.SetProjectId(projectId)
	config.SetProjectScert(projectScert)
}
