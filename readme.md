# Go语言上传文件到Go-FastDFS
## 安装
```shell
go get -u github.com/isafer-sen/gofastdfs
```
## 使用
```go
func Upload(file *multipart.FileHeader) (err error, fileInfo gofastdfs.FileInfo) {
	dfs := gofastdfs.NewFastDFSConfig(config.FastDFSURL, "")
	return dfs.UploadFile(file)
}
```