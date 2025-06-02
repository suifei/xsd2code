# 安装指南

本页面详细介绍了在不同操作系统上安装XSD2Code的各种方法。

## 📋 系统要求

### 最低要求
- **操作系统**: Windows 10+, macOS 10.15+, Linux (Ubuntu 18.04+, CentOS 7+)
- **内存**: 512MB RAM
- **磁盘空间**: 50MB 可用空间

### 开发环境要求（从源码构建）
- **Go版本**: 1.21 或更高版本
- **Git**: 用于克隆源码

## 🚀 安装方法

### 方法1：预构建二进制文件（推荐）

这是最简单快捷的安装方式：

#### Windows

```powershell
# 下载Windows版本
curl -L -o xsd2code.exe https://github.com/suifei/xsd2code/releases/latest/download/xsd2code-windows-amd64.exe

# 或者使用PowerShell
Invoke-WebRequest -Uri "https://github.com/suifei/xsd2code/releases/latest/download/xsd2code-windows-amd64.exe" -OutFile "xsd2code.exe"

# 添加到PATH（可选）
# 将xsd2code.exe移动到PATH中的目录，如 C:\Windows\System32\
```

#### macOS

```bash
# 下载macOS版本
curl -L -o xsd2code https://github.com/suifei/xsd2code/releases/latest/download/xsd2code-darwin-amd64

# 设置执行权限
chmod +x xsd2code

# 移动到PATH目录
sudo mv xsd2code /usr/local/bin/

# Apple Silicon Mac (M1/M2)
curl -L -o xsd2code https://github.com/suifei/xsd2code/releases/latest/download/xsd2code-darwin-arm64
chmod +x xsd2code
sudo mv xsd2code /usr/local/bin/
```

#### Linux

```bash
# 下载Linux版本
wget https://github.com/suifei/xsd2code/releases/latest/download/xsd2code-linux-amd64 -O xsd2code

# 或使用curl
curl -L -o xsd2code https://github.com/suifei/xsd2code/releases/latest/download/xsd2code-linux-amd64

# 设置执行权限
chmod +x xsd2code

# 移动到PATH目录
sudo mv xsd2code /usr/local/bin/

# 或者创建符号链接
sudo ln -s $(pwd)/xsd2code /usr/local/bin/xsd2code
```

### 方法2：从源码构建

如果您想要最新的开发版本或需要自定义构建：

#### 步骤1：安装Go

首先确保已安装Go 1.21或更高版本：

```bash
# 检查Go版本
go version

# 如果没有安装Go，请访问 https://golang.org/dl/ 下载安装
```

#### 步骤2：克隆源码

```bash
# 克隆仓库
git clone https://github.com/suifei/xsd2code.git
cd xsd2code
```

#### 步骤3：构建

```bash
# 基本构建
go build -o xsd2code cmd/main.go

# 优化构建（推荐生产环境）
go build -ldflags="-s -w" -o xsd2code cmd/main.go

# 交叉编译（为其他平台构建）
# Windows
GOOS=windows GOARCH=amd64 go build -o xsd2code.exe cmd/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o xsd2code-macos cmd/main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o xsd2code-linux cmd/main.go
```

#### 步骤4：安装到系统

```bash
# Linux/macOS
sudo cp xsd2code /usr/local/bin/

# 或者添加到个人bin目录
mkdir -p ~/bin
cp xsd2code ~/bin/
echo 'export PATH=$PATH:~/bin' >> ~/.bashrc
source ~/.bashrc
```

### 方法3：使用Go Install

如果您已经安装了Go，可以直接使用go install：

```bash
go install github.com/suifei/xsd2code/cmd@latest
```

这会自动下载、编译并安装xsd2code到您的GOPATH/bin目录。

### 方法4：使用包管理器

#### macOS - Homebrew（计划中）

```bash
# 计划加入Homebrew，敬请期待
# brew install xsd2code
```

#### Linux - Snap（计划中）

```bash
# 计划制作Snap包，敬请期待
# sudo snap install xsd2code
```

## ✅ 验证安装

安装完成后，验证是否正确安装：

```bash
# 检查版本
xsd2code -version

# 显示帮助
xsd2code -help

# 测试基本功能
xsd2code -xsd=examples/simple_example.xsd
```

预期输出：

```
XSD2Code v3.1.0
Build: 2025-06-02
Go version: go1.21.0
```

## 🔧 配置环境

### 环境变量

您可以设置以下环境变量来自定义行为：

```bash
# 设置默认输出目录
export XSD2CODE_OUTPUT_DIR="/path/to/output"

# 设置默认包名
export XSD2CODE_DEFAULT_PACKAGE="models"

# 启用调试模式
export XSD2CODE_DEBUG=true
```

### 配置文件（计划中）

未来版本将支持配置文件：

```yaml
# ~/.xsd2code/config.yaml
default:
  language: go
  package: models
  output_dir: ./generated
  json_tags: true
  validation: true
```

## 🚀 IDE集成

### VS Code

创建VS Code任务来快速生成代码：

```json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "XSD2Code Generate",
            "type": "shell",
            "command": "xsd2code",
            "args": [
                "-xsd=${input:xsdFile}",
                "-output=${input:outputFile}",
                "-package=${input:packageName}"
            ],
            "group": "build",
            "presentation": {
                "echo": true,
                "reveal": "always",
                "focus": false,
                "panel": "shared"
            }
        }
    ],
    "inputs": [
        {
            "id": "xsdFile",
            "description": "XSD file path",
            "default": "schema.xsd",
            "type": "promptString"
        },
        {
            "id": "outputFile", 
            "description": "Output file path",
            "default": "types.go",
            "type": "promptString"
        },
        {
            "id": "packageName",
            "description": "Package name",
            "default": "models",
            "type": "promptString"
        }
    ]
}
```

### JetBrains IDEs

创建External Tool：

1. 打开 Settings → Tools → External Tools
2. 点击 + 添加新工具
3. 配置：
   - Name: XSD2Code
   - Program: xsd2code
   - Arguments: -xsd=$FilePath$ -output=$FileDir$/types.go
   - Working directory: $ProjectFileDir$

## 🐳 Docker支持

您也可以使用Docker运行XSD2Code：

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o xsd2code cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/xsd2code .
ENTRYPOINT ["./xsd2code"]
```

构建和使用：

```bash
# 构建镜像
docker build -t xsd2code .

# 使用Docker运行
docker run --rm -v $(pwd):/workspace xsd2code -xsd=/workspace/schema.xsd -output=/workspace/types.go
```

## 🔍 故障排除

### 常见安装问题

#### "command not found"错误

确保二进制文件在PATH中：

```bash
# 检查PATH
echo $PATH

# 查找xsd2code
which xsd2code

# 手动添加到PATH
export PATH=$PATH:/path/to/xsd2code/directory
```

#### 权限错误

```bash
# 添加执行权限
chmod +x xsd2code

# 如果是权限被拒绝
sudo chown $USER:$USER xsd2code
```

#### Go版本不兼容

```bash
# 检查Go版本
go version

# 需要Go 1.21+，升级Go：
# 访问 https://golang.org/dl/ 下载最新版本
```

### 网络问题

如果下载失败，可以：

1. 使用代理或VPN
2. 从备用镜像下载
3. 手动下载release文件

## 📞 获取帮助

如果安装过程中遇到问题：

1. 查看 [[常见问题|FAQ]]
2. 查看 [[故障排除|Troubleshooting]]
3. 在 [GitHub Issues](https://github.com/suifei/xsd2code/issues) 报告问题

## 🔄 更新

### 更新预构建版本

重新下载最新版本的二进制文件，替换旧版本即可。

### 更新源码构建版本

```bash
cd xsd2code
git pull origin main
go build -o xsd2code cmd/main.go
```

### 检查更新

```bash
# 检查当前版本
xsd2code -version

# 检查最新版本
curl -s https://api.github.com/repos/suifei/xsd2code/releases/latest | grep "tag_name"
```

---

🎉 **安装完成！** 现在您可以查看 [[快速开始|Quick-Start]] 指南开始使用XSD2Code。
