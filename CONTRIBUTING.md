# 贡献指南

感谢您对 XSD2Code 项目的兴趣！我们欢迎各种形式的贡献。

## 如何贡献

### 报告问题
- 在提交问题之前，请先搜索现有的 Issues
- 使用清晰、描述性的标题
- 提供复现步骤和预期结果
- 包含相关的XSD文件和错误信息

### 提交功能请求
- 清楚地描述新功能的用途和价值
- 提供具体的使用场景
- 考虑向后兼容性

### 代码贡献

#### 开发环境设置
```bash
# 克隆项目
git clone https://github.com/suifei/xsd2code.git
cd xsd2code

# 安装依赖
go mod tidy

# 构建项目
go build -o xsd2code.exe ./cmd

# 运行测试
go test ./...
```

#### 代码规范
- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 添加适当的注释和文档
- 为新功能编写测试

#### 提交流程
1. Fork 这个项目
2. 创建新的功能分支: `git checkout -b feature/新功能名称`
3. 提交您的更改: `git commit -am '添加新功能'`
4. 推送到分支: `git push origin feature/新功能名称`
5. 创建 Pull Request

#### Pull Request 指南
- 提供清晰的PR描述
- 关联相关的Issues
- 确保所有测试通过
- 保持代码风格一致
- 更新相关文档

## 项目结构

```
xsd2code/
├── cmd/                    # 主程序入口
│   └── main.go
├── pkg/                    # 核心包
│   ├── generator/          # 代码生成器
│   ├── types/             # 类型定义
│   ├── validator/         # XML验证器
│   └── xsdparser/         # XSD解析器
├── examples/              # 示例文件
├── test/                  # 测试文件
└── docs/                  # 文档
```

## 开发指南

### 添加新的XSD特性支持
1. 在 `pkg/types/xsd_types.go` 中添加类型定义
2. 在 `pkg/xsdparser/parser.go` 中添加解析逻辑
3. 在 `pkg/generator/codegen.go` 中添加代码生成逻辑
4. 编写测试用例

### 添加新的目标语言支持
1. 在 `pkg/generator/codegen.go` 中实现 `LanguageMapper` 接口
2. 添加类型映射表
3. 实现代码生成逻辑
4. 更新命令行参数处理

### 测试
- 单元测试：`go test ./...`
- 集成测试：使用真实的XSD文件测试
- 性能测试：使用大型XSD文件进行基准测试

## 发布流程

### 版本号规则
我们使用语义化版本控制 (SemVer)：
- `MAJOR.MINOR.PATCH`
- MAJOR：不兼容的API更改
- MINOR：向后兼容的功能添加
- PATCH：向后兼容的问题修复

### 发布步骤
1. 更新 `CHANGELOG.md`
2. 更新版本号
3. 创建 git tag：`git tag v3.1.0`
4. 推送 tag：`git push origin v3.1.0`
5. GitHub Actions 自动构建和发布

## 社区准则

### 行为准则
- 保持友善和专业
- 尊重不同观点
- 建设性地处理冲突
- 帮助营造包容的环境

### 沟通渠道
- GitHub Issues：问题报告和功能请求
- GitHub Discussions：一般讨论和问答
- Pull Requests：代码审查和讨论

## 获得帮助

如果您需要帮助：
1. 查看项目文档和示例
2. 搜索现有的 Issues 和 Discussions
3. 创建新的 Issue 或 Discussion
4. 发送邮件给维护者

## 许可证

通过提交代码，您同意您的贡献将在 MIT 许可证下发布。

---

再次感谢您的贡献！🎉
