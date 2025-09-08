// package model
//
// //这个模块应写在服务层
// import (
//
//	"My-Blog/utils"
//	"My-Blog/utils/errmsg"
//	"context"
//
//	//"github.com/qiniu/api.v7/v7/auth/qbox"
//	//"github.com/qiniu/api.v7/v7/storage"
//	"github.com/qiniu/go-sdk/v7/auth/qbox"
//	"github.com/qiniu/go-sdk/v7/storage"
//	"mime/multipart"
//
// )
//
// var AccessKey = utils.AccessKey
// var SecretKey = utils.SecretKey
// var Bucket = utils.Bucket
// var ImgUrl = utils.QiniuSever
//
//	func UploadFile(file multipart.File, fileSize int64) (string, int) {
//		putPolicy := storage.PutPolicy{
//			Scope: Bucket,
//		}
//		mac := qbox.NewMac(AccessKey, SecretKey)
//		upToken := putPolicy.UploadToken(mac)
//		cfg := storage.Config{
//			Zone:          &storage.ZoneHuabei,
//			UseCdnDomains: false,
//			UseHTTPS:      false,
//		}
//		putExtra := storage.PutExtra{}
//		formUploader := storage.NewFormUploader(&cfg)
//		ret := storage.PutRet{}
//
//		err := formUploader.PutFileWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
//		if err != nil {
//			return "", errmsg.ERROR
//		}
//		url := ImgUrl + "/" + ret.Key
//		return url, errmsg.SUCCSE
//	}
package model

import (
	"My-Blog/utils"
	"My-Blog/utils/errmsg"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuSever

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	// 重置文件指针到开头
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", errmsg.ERROR
	}

	// 创建临时文件保存multipart.File内容
	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		fmt.Printf("创建临时文件失败: %v\n", err)
		return "", errmsg.ERROR
	}
	defer os.Remove(tempFile.Name()) // 确保临时文件会被删除

	// 将multipart.File内容复制到临时文件
	if _, err := io.Copy(tempFile, file); err != nil {
		fmt.Printf("复制文件内容失败: %v\n", err)
		return "", errmsg.ERROR
	}
	tempFile.Close()

	putPolicy := storage.PutPolicy{
		Scope:   Bucket,
		Expires: 3600,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	region, err := storage.GetRegion(AccessKey, Bucket)
	if err != nil {
		fmt.Printf("自动获取区域失败，使用默认华北区域: %v\n", err)
		cfg.Zone = &storage.ZoneHuabei
	} else {
		cfg.Zone = region
	}

	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 使用临时文件路径作为参数，匹配SDK要求
	err = formUploader.PutFileWithoutKey(context.Background(), &ret, upToken, tempFile.Name(), &putExtra)
	if err != nil {
		fmt.Printf("上传错误详情: %v\n", err)
		return "", errmsg.ERROR
	}

	url := strings.TrimSuffix(ImgUrl, "/") + "/" + ret.Key
	return url, errmsg.SUCCSE
}
