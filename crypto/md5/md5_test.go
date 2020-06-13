package md5

import (
	"testing"

	"github.com/wenit/gotools/crypto/hex"
)

var srcData = "abc123中文ABC"
var encodeData = "89068e0e2bddb4b00d821e51bd1240e8"

func TestSum(t *testing.T) {
	src := []byte(srcData)
	resultData := Sum(src)

	t.Log("<<", src)
	t.Log("<<", srcData)
	t.Log(">>", resultData)

	if hex.EncodeToString(resultData) != encodeData {
		t.Error(false)
	}
}

func TestSumFileToString(t *testing.T) {
	resultData, err := SumFileToString("test.txt")
	if err != nil {
		t.Error(err)
	}

	t.Log("<<", srcData)
	t.Log(">>", resultData)
	//t.Log(">> 字符串：", string(resultData))

	if resultData != encodeData {
		t.Error(false)
	}
}

func TestSumStringToString(t *testing.T) {
	resultData := SumStringToString(srcData)

	t.Log("<<", srcData)
	t.Log(">>", resultData)
	//t.Log(">> 字符串：", string(resultData))

	if resultData != encodeData {
		t.Error(false)
	}
}

func TestSumToString(t *testing.T) {
	src := []byte(srcData)
	resultData := SumToString(src)

	t.Log("<<", src)
	t.Log("<<", srcData)
	t.Log(">>", resultData)
	//t.Log(">> 字符串：", string(resultData))

	if resultData != encodeData {
		t.Error(false)
	}
}
