package service

import (
	"encoding/json"
	"log"

	Tools "esgin-go-sdk/tools"
)

func FileUpload() (*Tools.Res[FileUploadUrlResponse], string) {
	apiUrl := "/v3/files/file-upload-url"
	filePath := "table.docx"
	md5str, size := Tools.CountFileMd5(filePath)
	// var uploadUrlRequest FileUploadRequest
	data := FileUploadRequest{
		ContentMd5:    md5str,
		ContentType:   "application/octet-stream",
		FileName:      "测试.pdf",
		FileSize:      size,
		ConvertToPDF:  true,
		ConvertToHTML: false,
	}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON:", err)
		return nil, ""
	}
	// err := json.Unmarshal([]byte(data), &uploadUrlRequest)
	// if err != nil {
	// 	log.Fatalf("json.Unmarshal failed: %", err)
	// }

	res, err := Tools.SendCommHttp[FileUploadUrlResponse](apiUrl, string(jsonStr), "POST")
	if err != nil {
		log.Fatalf("SendHttp failed: %s", err)
	}

	return res, md5str
}

type FileUploadUrlResponse struct {
	FileId        string `json:"fileId"`
	FileUploadUrl string `json:"fileUploadUrl"`
}
type FileUploadRequest struct {
	ContentMd5    string `json:"contentMd5,omitempty"`
	ContentType   string `json:"contentType,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	FileSize      int64  `json:"fileSize,omitempty"`
	ConvertToPDF  bool   `json:"convertToPDF,omitempty"`
	ConvertToHTML bool   `json:"convertToHTML,omitempty"`
}
