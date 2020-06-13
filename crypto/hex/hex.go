// Package hex 提供常用16进制编解码方法
package hex

import (
	"encoding/hex"
)

// Encode 将字节数组转换为16进制字节数组
func Encode(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

// Decode 将16进制字节数组转换为字节数组
func Decode(src []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	return dst, err
}

// EncodeString 将明文字符串编码为16进制字节数组
func EncodeString(s string) []byte {
	src := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(s)))
	hex.Encode(dst, src)
	return dst
}

// DecodeString 将16进制字符串解码为明文字节数组
func DecodeString(s string) ([]byte, error) {
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(s)))
	_, err := hex.Decode(dst, src)
	return dst, err
}

//EncodeToString 编码明文字节数组为16进制字符串
func EncodeToString(src []byte) string {
	dst := Encode(src)
	return string(dst)
}

//DecodeToString 解码16进字节数组为字符串
func DecodeToString(src []byte) string {
	dst, err := Decode(src)
	if err != nil {
		return ""
	}
	return string(dst)
}

//EncodeStringToHexString 编码明文字符串加密为16进制字符串
func EncodeStringToHexString(s string) string {
	return EncodeToString([]byte(s))
}

//DecodeHexStringToString 解码16进制字符串为明文字符串
func DecodeHexStringToString(s string) string {
	return DecodeToString([]byte(s))
}
