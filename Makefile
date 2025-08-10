# 變數：可執行檔名稱，用空格分隔多個 target
BINARY_NAMES = server mygui

BUILD_DIR = bin
CONFIG_FILE = etc/config.json

# 來源碼路徑對應
SERVER_MAIN = cmd/server/main.go
MYGUI_MAIN = cmd/mygui/main.go
MYGUI_SRCS := $(wildcard cmd/mygui/*.go)


# 預設目標
.PHONY: all
all: build

# 編譯專案，分別建置每個二進位檔
.PHONY: build
build: $(BINARY_NAMES)

server:
	go build -o $(BUILD_DIR)/server $(SERVER_MAIN)

mygui:
	go build -o $(BUILD_DIR)/mygui $(MYGUI_SRCS)

# 執行專案（需要指定執行哪個）
.PHONY: run-server run-mygui
run-server:
	go run $(SERVER_MAIN) -config $(CONFIG_FILE)

run-mygui:
	go run $(MYGUI_MAIN) -config $(CONFIG_FILE)

# 清除編譯產物
.PHONY: clean
clean:
	rm -f $(BUILD_DIR)/server $(BUILD_DIR)/mygui

# 自動格式化程式碼
.PHONY: fmt
fmt:
	go fmt ./...

# 靜態檢查（需安裝 golangci-lint）
.PHONY: lint
lint:
	golangci-lint run

# 安裝依賴工具（選用）
.PHONY: install-tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
