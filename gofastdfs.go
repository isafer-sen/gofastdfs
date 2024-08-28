package gofastdfs

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

type FileInfo struct {
	Domain  string `json:"domain"`
	Md5     string `json:"md5"`
	Path    string `json:"path"`
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Scene   string `json:"scene"`
	Scenes  string `json:"scenes"`
	Src     string `json:"src"`
	URL     string `json:"url"`
}

type FastDFSConfig struct {
	FastDFSURL string `json:"fastdfs_url"`
	Auth       string `json:"auth"`
}

func NewFastDFSConfig(fastDFSURL, auth string) *FastDFSConfig {
	return &FastDFSConfig{
		FastDFSURL: fastDFSURL,
		Auth:       auth,
	}
}

func (c *FastDFSConfig) UploadFile(file *multipart.FileHeader) (err error, fileInfo FileInfo) {
	// 打开文件
	src, err := file.Open()
	if err != nil {
		return err, fileInfo
	}
	defer src.Close()

	// 创建一个buffer来存储文件内容
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	_ = writer.WriteField("output", "json")
	if c.Auth != "" {
		_ = writer.WriteField("auth_token", c.Auth)
	}
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return err, fileInfo
	}
	_, err = io.Copy(part, src)
	if err != nil {
		return err, fileInfo
	}
	writer.Close()

	// 发送POST请求到go-fastdfs服务器
	resp, err := http.Post(c.FastDFSURL, writer.FormDataContentType(), &buf)
	if err != nil {
		return err, fileInfo
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, fileInfo
	}
	return json.Unmarshal(body, &fileInfo), fileInfo
}
