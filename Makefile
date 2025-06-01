# XSD2Code Makefile

.PHONY: build test clean install help example

# 默认目标
all: build

# 构建主程序
build:
	@echo "构建 XSD2Code..."
	go build -ldflags="-s -w" -o xsd2code.exe ./cmd
	@echo "构建完成: xsd2code.exe"

# 运行测试
test:
	@echo "运行测试..."
	go test -v ./...

# 运行示例
example: build
	@echo "运行示例..."
	./xsd2code.exe -xsd=examples/simple_example.xsd -output=examples/simple_types.go -package=example -json
	@echo "示例完成，查看 examples/simple_types.go"

# 构建所有平台版本
build-all:
	@echo "构建多平台版本..."
	@mkdir -p dist
	
	# Windows
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/xsd2code-windows-amd64.exe ./cmd
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o dist/xsd2code-windows-386.exe ./cmd
	
	# Linux
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/xsd2code-linux-amd64 ./cmd
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o dist/xsd2code-linux-386 ./cmd
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o dist/xsd2code-linux-arm64 ./cmd
	
	# macOS
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o dist/xsd2code-darwin-amd64 ./cmd
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o dist/xsd2code-darwin-arm64 ./cmd
	
	@echo "多平台构建完成，文件位于 dist/ 目录"

# 安装依赖
deps:
	@echo "安装依赖..."
	go mod tidy
	go mod download

# 清理构建文件
clean:
	@echo "清理构建文件..."
	rm -f xsd2code.exe xsd2go.exe cmd.exe
	rm -f test_*.go
	rm -f examples/simple_types.go
	rm -rf dist/
	@echo "清理完成"

# 安装到系统
install: build
	@echo "安装 XSD2Code 到系统..."
	cp xsd2code.exe $(GOPATH)/bin/ || cp xsd2code.exe $(HOME)/go/bin/
	@echo "安装完成"

# 显示帮助
help:
	@echo "XSD2Code Makefile"
	@echo "=================="
	@echo ""
	@echo "可用目标:"
	@echo "  build      - 构建主程序"
	@echo "  test       - 运行测试"
	@echo "  example    - 运行示例"
	@echo "  build-all  - 构建所有平台版本"
	@echo "  deps       - 安装依赖"
	@echo "  clean      - 清理构建文件"
	@echo "  install    - 安装到系统"
	@echo "  help       - 显示此帮助信息"
	@echo ""
	@echo "示例用法:"
	@echo "  make build"
	@echo "  make test"
	@echo "  make example"
