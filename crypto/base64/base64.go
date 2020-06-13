// Package base64   提供常用base64编解码方法
package base64

import (
	"encoding/base64"
)

// Encode 编码明文字节数组为Base64字节数组
func Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// Decode 解码Base64字节数组为字节数组
func Decode(src []byte) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	_, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// EncodeString 编码明文字符串为Base64字节数组
func EncodeString(src string) []byte {
	return Encode([]byte(src))
}

// DecodeString 解码Base64字符串为字节数组
func DecodeString(src string) ([]byte, error) {
	return Decode([]byte(src))
}

// EncodeToString 编码明文字节数组为Base64字符串
func EncodeToString(src []byte) string {
	return string(Encode(src))
}

// DecodeToString 解码Base64字节数组为字符串
func DecodeToString(src []byte) string {
	dst, err := Decode(src)
	if err != nil {
		return ""
	}
	return string(dst)
}

// EncodeStringToBase64String 编码明文字符串为Base64字符串
func EncodeStringToBase64String(src string) string {
	dst := base64.StdEncoding.EncodeToString([]byte(src))
	return dst
}

// DecodeBase64StringToString 解码Base64字符串为明文字符串
func DecodeBase64StringToString(src string) string {
	dst, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(dst)
}
