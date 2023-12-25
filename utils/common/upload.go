package common

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/xuri/excelize/v2"

	"time"
)

// 上传图片到七牛云，然后返回状态和图片的url
func UploadToQiNiu(src *excelize.File, AccessKey, SerectKey, Bucket, QnUrl string) (int, string) {

	upToken := GetQNUploadToken(AccessKey, SerectKey, Bucket)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南区
		UseCdnDomains: true,
		UseHTTPS:      true, // 非https
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{} // 上传后返回的结果
	putExtra := storage.PutExtra{
		//MimeType: "",
	} // 额外参数

	// 上传 自定义key，可以指定上传目录及文件名和后缀，
	key := fmt.Sprintf("用户反馈_%s%d%s", "", time.Now().Unix(), ".xlsx") // 上传路径，如果当前目录中已存在相同文件，则返回上传失败错误

	buf, err := src.WriteToBuffer()
	if err != nil {
		return 0, ""
	}
	rd := bytes.NewReader(buf.Bytes())
	//err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, rd, int64(len(buf.Bytes())), &putExtra)
	err = formUploader.Put(context.Background(), &ret, upToken, key, rd, int64(len(buf.Bytes())), &putExtra)

	// 以默认key方式上传
	// err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, src, fileSize, &putExtra)

	// 自定义key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	// 默认key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	if err != nil {
		code := 501
		return code, err.Error()
	}

	url := QnUrl + "/" + ret.Key // 返回上传后的文件访问路径
	return 0, url
}

func GetQNUploadToken(qnAccessKey, qnSerectKey, qnBucket string) string {
	putPolicy := storage.PutPolicy{
		Scope: qnBucket,
	}
	mac := qbox.NewMac(qnAccessKey, qnSerectKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
