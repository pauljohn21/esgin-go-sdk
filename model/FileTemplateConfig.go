package Model

// 获取文件上传地址json信息配置
type GetFileUploadUrlInfo struct {
	ContentMd5   string `json:"contentMd5,omitempty"`
	ContentType  string `json:"contentType,omitempty"`
	ConvertToPDF bool   `json:"convertToPDF,omitempty"`
	FileName     string `json:"fileName,omitempty"`
	FileSize     int64  `json:"fileSize,omitempty"`
}
type FileTemplatereq struct {
	FileId        string `json:"fileId"`
	FileUploadUrl string `json:"fileUploadUrl"`
}
