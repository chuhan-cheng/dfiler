package main

import (
	"io" // 導入 io，用於讀取文件內容
	"strings"
	// 導入 strings，用於字符串操作

	"fyne.io/fyne/v2"         // 導入 Fyne 框架的核心包
	"fyne.io/fyne/v2/dialog"  // 導入 Fyne 的對話框包
	"fyne.io/fyne/v2/storage" // 導入 Fyne 的存儲包
)

// saveFunc 方法返回一個保存當前文件的函數
func (cfg *config) saveFunc(win fyne.Window) func() {
	return func() {
		// 檢查當前是否有打開的文件
		if cfg.CurrentFile != nil {
			// 獲取當前文件的寫入器
			write, err := storage.Writer(cfg.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			// 將編輯器中的文本寫入文件
			write.Write([]byte(cfg.EditWidget.Text))
			// 確保文件在寫入後關閉
			defer write.Close()
		}
	}
}

// openFunc 方法返回一個打開文件的函數
func (cfg *config) openFunc(win fyne.Window) func() {
	return func() {
		// 創建文件打開對話框
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			// 錯誤處理：如果發生錯誤，則顯示錯誤對話框
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			// 如果沒有選擇文件，直接返回
			if read == nil {
				return
			}
			// 讀取文件內容
			data, err := io.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			// 確保文件在讀取後關閉
			defer read.Close()
			// 將讀取的內容設置到編輯器中
			cfg.EditWidget.SetText(string(data))
			// 更新當前文件的 URI
			cfg.CurrentFile = read.URI()
			// 更新窗口標題，包含當前文件名
			win.SetTitle(cfg.BaseTitle + "-" + read.URI().Name())
			cfg.MenuItem.Disabled = false // 啟用保存菜單項
		}, win)
		// 設置文件過濾器
		openDialog.SetFilter(filter)
		openDialog.Show() // 顯示打開對話框
	}
}

// saveAsFunc 方法返回一個保存文件的函數
func (cfg *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		// 創建文件保存對話框
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			// 錯誤處理：如果發生錯誤，則顯示錯誤對話框
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			// 如果沒有選擇文件，直接返回
			if write == nil {
				return
			}
			// 檢查文件擴展名是否為 .md
			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("錯誤", "必須是.md擴展名", win)
				return
			}
			// 將編輯器中的文本寫入文件
			write.Write([]uint8(cfg.EditWidget.Text))
			// 更新當前文件的 URI
			cfg.CurrentFile = write.URI()
			// 確保文件在寫入後關閉
			defer write.Close()
			// 更新窗口標題，包含當前文件名
			win.SetTitle(cfg.BaseTitle + "-" + write.URI().Name())
			cfg.MenuItem.Disabled = false // 啟用保存菜單項
		}, win)
		// 設置默認文件名和過濾器
		saveDialog.SetFileName("未命名.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show() // 顯示保存對話框
	}
}
