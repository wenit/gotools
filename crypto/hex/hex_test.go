package hex

import (
	"testing"
)

var srcData = "abc123中文ABC"
var encodeData = "616263313233e4b8ade69687414243"

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

func TestEncodeString(t *testing.T) {
	resultData := EncodeString(srcData)
	t.Log("<<", srcData)
	t.Log(">>", resultData)
	t.Log(">> 字符串：", string(resultData))

	if string(resultData) != encodeData {
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

func TestEncodeStringToString(t *testing.T) {
	resultData := EncodeStringToHexString(srcData)
	t.Log("<<", srcData)
	t.Log(">>", resultData)

	if resultData != encodeData {
		t.Error(false)
	}
}

func TestDecodeHexStringToString(t *testing.T) {
	resultData := DecodeHexStringToString(encodeData)
	t.Log("<<", encodeData)
	t.Log(">>", resultData)

	if resultData != srcData {
		t.Error(false)
	}
}

func BenchmarkEncodeStringToHexString(b *testing.B) {
	EncodeStringToHexString(srcData)
}

func BenchmarkDecodeHexStringToString(b *testing.B) {
	DecodeHexStringToString(encodeData)
}
