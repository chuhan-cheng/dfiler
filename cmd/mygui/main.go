package main

import (
	"fyne.io/fyne/v2"           // 導入 Fyne 框架的主包
	"fyne.io/fyne/v2/app"       // 導入 Fyne 框架的應用程序包
	"fyne.io/fyne/v2/container" // 導入 Fyne 的容器包

	// 導入 Fyne 的存儲包
	"fyne.io/fyne/v2/storage" // 導入 Fyne 的存儲包
	"fyne.io/fyne/v2/widget"  // 導入 Fyne 的小部件包
)

// config 結構體用於存儲應用程序的配置信息
type config struct {
	EditWidget    *widget.Entry    // 編輯區域，使用 Entry 小部件
	PreviewWidget *widget.RichText // 預覽區域，使用 RichText 小部件
	CurrentFile   fyne.URI         // 當前打開的文件 URI
	MenuItem      *fyne.MenuItem   // 菜單項
	BaseTitle     string           // 窗口標題的基本字符串
}

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

var cfg config // 聲明一個全局變量 cfg，類型為 config

func main() {
	a := app.New() // 創建一個新的 Fyne 應用程序實例
	// 創建新的窗口，標題為 "Markdown編輯器"
	w := a.NewWindow("Markdown編輯器")
	cfg.BaseTitle = "Markdown編輯器" // 設置基本標題
	// 創建編輯區域和預覽區域的 UI
	edit, preview := cfg.makeUI()
	// 創建菜單
	cfg.createMenu(w)
	// 設置窗口內容為水平分割的編輯區域和預覽區域
	w.SetContent(container.NewHSplit(edit, preview))
	// 設置窗口的初始大小
	w.Resize(fyne.Size{Width: 800, Height: 600})
	// 窗口居中顯示在屏幕上
	w.CenterOnScreen()
	// 顯示窗口並運行應用程序
	w.ShowAndRun()
}
