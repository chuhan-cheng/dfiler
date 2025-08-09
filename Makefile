# 專案參數
BINARY_NAME=server
BUILD_DIR=bin
CONFIG_FILE=etc/config.json
MAIN_FILE=cmd/server/main.go

# 預設目標
.PHONY: all
all: build

# 編譯專案
.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# 執行專案（記得使用絕對或相對路徑）
.PHONY: run
run:
	go run $(MAIN_FILE) -config $(CONFIG_FILE)

# 清除編譯產物
.PHONY: clean
clean:
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

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
