// Package md5 摘要算法基础库
package md5

import (
	smd5 "crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

//Sum 将字节数组计算为MD5字节数组
func Sum(b []byte) []byte {
	rst := smd5.Sum(b)
	return rst[:]
}

//SumToString 将字节数组计算为MD5字符串，字符串为16进制字符串
func SumToString(b []byte) string {
	return hex.EncodeToString(Sum(b))
}

//SumStringToString 将字符串计算为MD5字符串，字符串为16进制字符串
func SumStringToString(s string) string {
	d := []byte(s)
	return SumToString(d)
}

//SumFileToString 将文件计算为MD5字符串，字符串为16进制字符串
func SumFileToString(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := smd5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
