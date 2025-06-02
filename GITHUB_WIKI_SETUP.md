# GitHub Wiki 设置指南

本文档说明如何为XSD2Code项目设置GitHub Wiki。

## 📋 前提条件

- 拥有GitHub仓库的管理员权限
- 已创建wiki内容文件（位于`wiki/`文件夹）

## 🚀 设置步骤

### 第一步：启用Wiki功能

1. 访问GitHub仓库主页：`https://github.com/suifei/xsd2code`
2. 点击 **Settings** 选项卡
3. 在左侧菜单中找到 **Features** 部分
4. 勾选 **Wikis** 复选框启用Wiki功能

### 第二步：创建Wiki

1. 在仓库主页点击 **Wiki** 选项卡
2. 点击 **Create the first page** 按钮
3. 在页面标题输入：`Home`
4. 将 `wiki/Home.md` 文件的内容复制到编辑器中
5. 点击 **Save Page** 保存

### 第三步：添加其他Wiki页面

按照以下顺序创建其他Wiki页面：

1. **Quick-Start** - 复制 `wiki/Quick-Start.md` 内容
2. **Installation** - 复制 `wiki/Installation.md` 内容
3. **Basic-Usage** - 复制 `wiki/Basic-Usage.md` 内容
4. **XSD-Features** - 复制 `wiki/XSD-Features.md` 内容
5. **Command-Line-Reference** - 复制 `wiki/Command-Line-Reference.md` 内容
6. **Advanced-Examples** - 复制 `wiki/Advanced-Examples.md` 内容
7. **FAQ** - 复制 `wiki/FAQ.md` 内容

### 第四步：设置页面链接

确保每个页面的内部链接正确：

- 使用格式：`[[页面标题]]` 或 `[[显示文本|页面标题]]`
- 例如：`[[快速开始|Quick-Start]]`

## 📝 Wiki页面结构

我们的Wiki包含以下页面：

```
Home.md                    - Wiki主页和导航
├── Quick-Start.md         - 5分钟快速开始指南
├── Installation.md        - 详细安装指南
├── Basic-Usage.md         - 基本用法说明
├── XSD-Features.md        - XSD特性支持详解
├── Command-Line-Reference.md - 完整命令行参考
├── Advanced-Examples.md   - 高级使用示例
└── FAQ.md                 - 常见问题解答
```

## 🎯 Wiki内容特色

- **中文本地化**: 所有内容使用中文编写，适合中文用户
- **丰富示例**: 包含大量代码示例和实际案例
- **完整教程**: 从入门到高级的完整学习路径
- **实际场景**: 企业级应用、工业自动化、数据分析等真实案例
- **最佳实践**: CI/CD集成、性能优化、故障排除等最佳实践

## 🔧 维护建议

### 定期更新

1. 当项目有新功能时，及时更新相关Wiki页面
2. 收集用户反馈，改进文档内容
3. 添加新的使用案例和示例

### 版本同步

1. 确保Wiki内容与代码版本保持同步
2. 在发布新版本时更新版本号和功能说明
3. 维护CHANGELOG与Wiki的一致性

### 用户反馈

1. 监控GitHub Issues中关于文档的反馈
2. 定期检查Wiki页面的访问统计
3. 根据用户需求调整文档结构和内容

## 📊 访问统计

启用Wiki后，可以通过以下方式查看访问统计：

1. GitHub仓库的Insights标签
2. Wiki页面的编辑历史
3. 第三方分析工具（如Google Analytics）

## 🔗 相关链接

- **主仓库**: https://github.com/suifei/xsd2code
- **Wiki地址**: https://github.com/suifei/xsd2code/wiki
- **Issues**: https://github.com/suifei/xsd2code/issues
- **Releases**: https://github.com/suifei/xsd2code/releases

---

📝 **注意**: 本文档仅供项目维护者参考，用户可以直接访问Wiki获取使用说明。
