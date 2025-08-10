package main

import (
	"fyne.io/fyne/v2"        // 導入 Fyne 框架的核心包
	"fyne.io/fyne/v2/widget" // 導入 Fyne 的小部件包
)

// makeUI 方法用於創建編輯和預覽區域的用戶界面
func (cfg *config) makeUI() (*widget.Entry, *widget.RichText) {
	// 創建一個多行文本輸入框用於編輯
	edit := widget.NewMultiLineEntry()
	// 創建一個富文本小部件用於Markdown預覽
	preview := widget.NewRichTextFromMarkdown("")
	// 將創建的編輯和預覽小部件保存到 config 結構體中
	cfg.EditWidget = edit
	cfg.PreviewWidget = preview
	// 當編輯內容改變時，解析Markdown並更新預覽
	edit.OnChanged = preview.ParseMarkdown
	// 返回編輯和預覽小部件
	return edit, preview
}

// createMenu 方法用於創建窗口的菜單
func (cfg *config) createMenu(win fyne.Window) {
	// 創建 "打開..." 菜單項，並指定點擊時調用的函數
	open := fyne.NewMenuItem("打開文件", cfg.openFunc(win))
	// 創建 "保存" 菜單項，並指定點擊時調用的函數
	save := fyne.NewMenuItem("保存文件", cfg.saveFunc(win))
	// 將保存菜單項存儲到 config 結構體中，並禁用它
	cfg.MenuItem = save
	cfg.MenuItem.Disabled = true // 空文件不能保存
	// 創建 "另存為..." 菜單項，並指定點擊時調用的函數
	saveAs := fyne.NewMenuItem("另存為", cfg.saveAsFunc(win))
	// 創建文件菜單，將上述菜單項添加到文件菜單中
	fileMenu := fyne.NewMenu("文件", open, save, saveAs)
	// 創建主菜單，並將文件菜單作為其內容
	menu := fyne.NewMainMenu(fileMenu)
	// 將創建的主菜單設置為窗口的菜單
	win.SetMainMenu(menu)

}
