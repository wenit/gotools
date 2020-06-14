// Package system 系统工具类
package system

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// substr 字符串截取
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// ParentDirectory 获取父级目录
func ParentDirectory(directory string) string {
	return substr(directory, 0, strings.LastIndex(directory, "/"))
}

// CurrentDirectory 获取当前执行程序目录
func CurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// UIDAndGID 获取当前进程用户ID和组ID
func UIDAndGID() (int, int) {
	uid := os.Getuid()
	gid := os.Getgid()
	return uid, gid
}

// LookupUIDAndGID 根据用户名获取指定用户的uid和gid
// 获取失败返回-1
func LookupUIDAndGID(name string) (int, int) {
	u, err := user.Lookup(name)
	if err != nil {
		return -1, -1
	}
	uid, _ := strconv.Atoi(u.Uid)
	gid, _ := strconv.Atoi(u.Gid)
	return uid, gid
}

// LookupGid 根据组名称获取组ID
// 获取失败返回-1
func LookupGid(name string) int {
	u, err := user.LookupGroup(name)
	if err != nil {
		return -1
	}
	gid, _ := strconv.Atoi(u.Gid)
	return gid
}

// GetPid 获取PID
func GetPid() int {
	return os.Getpid()
}

// GetWd 获取工作目录
func GetWd() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

// GetTime 获取当前时间，格式：2006-01-02 15:04:05
func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetTimestamp 获取当前时间戳(秒)
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// GetTimestampMilli 获取当前时间戳(毫秒)
func GetTimestampMilli() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

// NumGoroutine 获取协程数量
func NumGoroutine() int {
	return runtime.NumGoroutine()
}
