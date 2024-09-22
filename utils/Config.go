package Tools

type EsignInitConfig struct {
	host         string
	projectId    string
	projectScert string
}

var esignInitConfig *EsignInitConfig = new(EsignInitConfig)

func InstaneEsignInitConfig() *EsignInitConfig {
	return esignInitConfig
}

func (e *EsignInitConfig) Host() string {
	return e.host
}

func (e *EsignInitConfig) SetHost(host string) {
	e.host = host
}

func (e *EsignInitConfig) ProjectId() string {
	return e.projectId
}

func (e *EsignInitConfig) SetProjectId(projectId string) {
	e.projectId = projectId
}

func (e *EsignInitConfig) ProjectScert() string {
	return e.projectScert
}

func (e *EsignInitConfig) SetProjectScert(projectScert string) {
	e.projectScert = projectScert
}
