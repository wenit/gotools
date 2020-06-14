package watcher

import (
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func TestFileWatcher(t *testing.T) {
	filePath := "test.txt"
	// 1、创建监听器
	watcher := NewFileWatcher(filePath)

	// 2、设置文件修改事件
	watcher.OnWriteEvent(func(e fsnotify.Event) {
		t.Log("Config file changed:", e.Name)
	})
	// 3、开始监听
	watcher.Start()

	time.Sleep(time.Second * 10)

	// 4、停止监听
	watcher.Stop()
}
