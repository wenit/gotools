package watcher

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"time"
)

type FileWatcher struct {
	watcher       *fsnotify.Watcher    //监听器
	LastModify    int64                //上次变更时间戳（秒）
	FilePath      string               //文件路径
	onCreateEvent func(fsnotify.Event) //创建事件
	onWriteEvent  func(fsnotify.Event) //写事件
	onRemoveEvent func(fsnotify.Event) //删除事件
	onRenameEvent func(fsnotify.Event) //重命名事件
	onChmodEvent  func(fsnotify.Event) //权限变更事件
}

// 创建监听器
func NewFileWatcher(filePath string) *FileWatcher {
	watcher, _ := fsnotify.NewWatcher()
	w := FileWatcher{
		watcher:    watcher,
		LastModify: 0,
		FilePath:   filePath,
	}
	w.watcher.Add(w.FilePath)
	return &w
}

// 设置创建事件
func (w *FileWatcher) OnCreateEvent(f func(fsnotify.Event)) {
	w.onCreateEvent = f
}

// 设置写事件
func (w *FileWatcher) OnWriteEvent(f func(fsnotify.Event)) {
	w.onWriteEvent = f
}

// 设置移除事件
func (w *FileWatcher) OnRemoveEvent(f func(fsnotify.Event)) {
	w.onRemoveEvent = f
}

// 设置重命名事件
func (w *FileWatcher) OnRenameEvent(f func(fsnotify.Event)) {
	w.onRenameEvent = f
}

// 设置权限修改事件
func (w *FileWatcher) OnChmodEvent(f func(fsnotify.Event)) {
	w.onChmodEvent = f
}

// 开始监听
func (w *FileWatcher) Start() {
	go func() {
		defer w.Stop()
		for {
			select {
			case ev := <-w.watcher.Events:
				{
					current := time.Now().Unix()
					c := current - w.LastModify
					w.LastModify = current
					if c > 0 {
						//判断事件发生的类型，如下5种
						// Create 创建
						// Write 写入
						// Remove 删除
						// Rename 重命名
						// Chmod 修改权限
						if ev.Op&fsnotify.Create == fsnotify.Create {
							// log.Println("创建文件 : ", ev.Name);
							if w.onCreateEvent != nil {
								w.onCreateEvent(ev)
							}
						}
						if ev.Op&fsnotify.Write == fsnotify.Write {
							// log.Println("写入文件 : ", ev.Name)
							if w.onWriteEvent != nil {
								w.onWriteEvent(ev)
							}
						}
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						if w.onRemoveEvent != nil {
							w.onRemoveEvent(ev)
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						if w.onRenameEvent != nil {
							w.onRenameEvent(ev)
						}
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						if w.onChmodEvent != nil {
							w.onChmodEvent(ev)
						}
					}
				}
			case err := <-w.watcher.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()
}

// 停止监听
func (w *FileWatcher) Stop() {
	w.watcher.Close()
}
