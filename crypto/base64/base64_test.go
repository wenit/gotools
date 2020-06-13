package base64

import (
	"testing"
)

var srcData = "abc123中文ABC"
var encodeData = "YWJjMTIz5Lit5paHQUJD"

func TestInitData(t *testing.T) {
	t.Log("加密明文：", srcData)
	t.Log("加密密文：", encodeData)
}


func TestEncode(t *testing.T) {
	src := []byte(srcData)
	resultData := Encode(src)

	t.Log("<<", src)
	t.Log(">>", resultData)
	t.Log(">> 字符串：", string(resultData))

	if string(resultData) != encodeData {
		t.Error(false)
	}
}

func TestEncodeString(t *testing.T) {
	resultData := EncodeString(srcData)
	t.Log("<<", srcData)
	t.Log(">>", resultData)
	t.Log(">> 字符串：", string(resultData))

	if string(resultData) != encodeData {
		t.Error(false)
	}
}

func TestEncodeToString(t *testing.T) {
	src := []byte(srcData)
	resultData := EncodeToString(src)
	t.Log("字符串 <<", srcData)
	t.Log("<<", src)
	t.Log(">>", resultData)

	if resultData != encodeData {
		t.Error(false)
	}
}

func TestEncodeStringToBase64String(t *testing.T) {
	resultData := EncodeStringToBase64String(srcData)
	t.Log("<<", srcData)
	t.Log(">>", resultData)
	t.Log(resultData == encodeData)

	if resultData != encodeData {
		t.Error(false)
	}
}

func TestDecode(t *testing.T) {
	src := []byte(encodeData)
	resultData, _ := Decode(src)

	t.Log("<<", src)
	t.Log(">>", resultData)
	t.Log(">> 字符串：", string(resultData))

	if string(resultData) != srcData {
		t.Error(false)
	}
}

func TestDecodeString(t *testing.T) {
	resultData, _ := DecodeString(encodeData)
	t.Log("<<", encodeData)
	t.Log(">>", resultData)
	t.Log(">> 字符串：", string(resultData))

	if string(resultData) != srcData {
		t.Error(false)
	}
}

func TestDecodeToString(t *testing.T) {
	src := []byte(encodeData)
	resultData := DecodeToString(src)
	t.Log("字符串 <<", encodeData)
	t.Log("<<", src)
	t.Log(">>", resultData)
	if resultData != srcData {
		t.Error(false)
	}
}

func TestDecodeBase64StringToString(t *testing.T) {
	resultData := DecodeBase64StringToString(encodeData)
	t.Log("<<", encodeData)
	t.Log(">>", resultData)
	t.Log(resultData == srcData)

	if resultData != srcData {
		t.Error(false)
	}
}


func BenchmarkEncodeStringToBase64String(b *testing.B) {
	EncodeStringToBase64String(srcData)
}

func BenchmarkDecodeBase64StringToString(b *testing.B) {
	DecodeBase64StringToString(encodeData)
}



