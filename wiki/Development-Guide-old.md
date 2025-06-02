# 开发指南 (Development Guide)

本页面为希望参与XSD2Code项目开发的贡献者提供详细的开发指南，包括开发环境搭建、代码规范、贡献流程等。

## 目录
- [开发环境搭建](#开发环境搭建)
- [项目结构](#项目结构)
- [开发流程](#开发流程)
- [代码规范](#代码规范)
- [测试指南](#测试指南)
- [贡献流程](#贡献流程)
- [发布流程](#发布流程)

## 开发环境搭建

### 系统要求
- **Node.js**: >= 16.0.0
- **npm**: >= 8.0.0 或 **yarn**: >= 1.22.0
- **Git**: >= 2.0.0
- **操作系统**: Windows 10+, macOS 10.15+, Linux (Ubuntu 18.04+)

### 环境准备

#### 1. 克隆项目
```bash
# 克隆主仓库
git clone https://github.com/yourusername/xsd2code.git
cd xsd2code

# 如果是fork，添加upstream
git remote add upstream https://github.com/original/xsd2code.git
```

#### 2. 安装依赖
```bash
# 使用npm
npm install

# 或使用yarn
yarn install
```

#### 3. 环境配置
```bash
# 复制环境配置文件
cp .env.example .env

# 编辑环境变量
# NODE_ENV=development
# LOG_LEVEL=debug
# CACHE_DIR=./cache
# TEMP_DIR=./temp
```

#### 4. 验证安装
```bash
# 运行测试
npm test

# 构建项目
npm run build

# 运行开发版本
npm run dev
```

### 开发工具推荐

#### IDE/编辑器
- **VS Code** (推荐)
  - 插件：TypeScript Hero, ESLint, Prettier, Jest
- **WebStorm**
- **Vim/Neovim** (配置TypeScript支持)

#### VS Code配置
```json
// .vscode/settings.json
{
  "typescript.preferences.importModuleSpecifier": "relative",
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true,
    "source.organizeImports": true
  },
  "jest.jestCommandLine": "npm test --",
  "files.associations": {
    "*.hbs": "handlebars"
  }
}
```

#### 推荐的VS Code插件
```json
// .vscode/extensions.json
{
  "recommendations": [
    "ms-vscode.vscode-typescript-next",
    "esbenp.prettier-vscode",
    "dbaeumer.vscode-eslint",
    "orta.vscode-jest",
    "ms-vscode.vscode-json",
    "redhat.vscode-yaml",
    "ms-vscode.vscode-xml"
  ]
}
```

## 项目结构

```
xsd2code/
├── src/                          # 源代码目录
│   ├── cli/                      # 命令行接口
│   │   ├── commands/             # CLI命令
│   │   ├── options/              # 命令选项
│   │   └── index.ts              # CLI入口
│   ├── core/                     # 核心功能
│   │   ├── parser/               # XSD解析器
│   │   ├── generator/            # 代码生成器
│   │   ├── config/               # 配置管理
│   │   └── types/                # 类型定义
│   ├── generators/               # 语言生成器
│   │   ├── java/                 # Java生成器
│   │   ├── typescript/           # TypeScript生成器
│   │   ├── csharp/               # C#生成器
│   │   ├── python/               # Python生成器
│   │   ├── go/                   # Go生成器
│   │   └── base/                 # 基础生成器
│   ├── templates/                # 代码模板
│   │   ├── java/                 # Java模板
│   │   ├── typescript/           # TypeScript模板
│   │   └── shared/               # 共享模板
│   ├── utils/                    # 工具函数
│   │   ├── naming.ts             # 命名转换
│   │   ├── validation.ts         # 验证工具
│   │   └── file.ts               # 文件操作
│   ├── plugins/                  # 插件系统
│   │   ├── base/                 # 基础插件
│   │   └── examples/             # 示例插件
│   └── index.ts                  # 库入口
├── test/                         # 测试代码
│   ├── unit/                     # 单元测试
│   ├── integration/              # 集成测试
│   ├── e2e/                      # 端到端测试
│   ├── fixtures/                 # 测试数据
│   └── helpers/                  # 测试工具
├── docs/                         # 文档
│   ├── api/                      # API文档
│   ├── guides/                   # 用户指南
│   └── examples/                 # 示例
├── scripts/                      # 构建脚本
│   ├── build.js                  # 构建脚本
│   ├── test.js                   # 测试脚本
│   └── release.js                # 发布脚本
├── config/                       # 配置文件
│   ├── jest.config.js            # Jest配置
│   ├── eslint.config.js          # ESLint配置
│   └── prettier.config.js        # Prettier配置
└── package.json                  # 项目配置
```

### 核心模块说明

#### src/core/parser/
XSD解析相关代码：
```typescript
// src/core/parser/XSDParser.ts
export class XSDParser {
  parse(content: string): XSDSchema;
  parseFile(filePath: string): Promise<XSDSchema>;
  validate(schema: XSDSchema): ValidationResult;
}

// src/core/parser/types.ts
export interface XSDElement {
  name: string;
  type: string;
  minOccurs?: number;
  maxOccurs?: number | 'unbounded';
}
```

#### src/generators/
各语言代码生成器：
```typescript
// src/generators/base/BaseGenerator.ts
export abstract class BaseGenerator implements LanguageGenerator {
  abstract generateClass(complexType: XSDComplexType): string;
  abstract generateEnum(simpleType: XSDSimpleType): string;
  
  protected renderTemplate(templateName: string, context: any): string {
    // 模板渲染逻辑
  }
}
```

## 开发流程

### 1. 功能开发流程

#### 创建功能分支
```bash
# 从main分支创建功能分支
git checkout main
git pull upstream main
git checkout -b feature/new-language-support

# 或修复bug
git checkout -b fix/parser-issue-123
```

#### 开发过程
1. **编写测试** (TDD方式，推荐)
2. **实现功能**
3. **运行测试**
4. **代码审查**
5. **文档更新**

#### 示例：添加新语言支持
```typescript
// 1. 创建生成器类
// src/generators/kotlin/KotlinGenerator.ts
export class KotlinGenerator extends BaseGenerator {
  name = 'kotlin';
  version = '1.0.0';
  
  generateClass(complexType: XSDComplexType): string {
    return this.renderTemplate('kotlin/class.hbs', {
      className: this.pascalCase(complexType.name),
      properties: complexType.elements.map(el => ({
        name: this.camelCase(el.name),
        type: this.mapType(el.type),
        nullable: el.minOccurs === 0
      }))
    });
  }
  
  generateEnum(simpleType: XSDSimpleType): string {
    return this.renderTemplate('kotlin/enum.hbs', {
      enumName: this.pascalCase(simpleType.name),
      values: simpleType.enumeration
    });
  }
  
  private mapType(xsdType: string): string {
    const typeMap: Record<string, string> = {
      'xs:string': 'String',
      'xs:int': 'Int',
      'xs:boolean': 'Boolean',
      'xs:decimal': 'BigDecimal',
      'xs:dateTime': 'LocalDateTime'
    };
    
    return typeMap[xsdType] || 'Any';
  }
}
```

```handlebars
{{!-- src/templates/kotlin/class.hbs --}}
{{#if packageName}}
package {{packageName}}

{{/if}}
import kotlinx.serialization.Serializable
import kotlinx.serialization.SerialName

@Serializable
data class {{className}}(
{{#each properties}}
    @SerialName("{{@key}}")
    val {{name}}: {{type}}{{#if nullable}}?{{/if}}{{#unless @last}},{{/unless}}
{{/each}}
)
```

```typescript
// 2. 注册生成器
// src/generators/index.ts
import { KotlinGenerator } from './kotlin/KotlinGenerator';

export const BUILTIN_GENERATORS = {
  java: JavaGenerator,
  typescript: TypeScriptGenerator,
  csharp: CSharpGenerator,
  python: PythonGenerator,
  go: GoGenerator,
  kotlin: KotlinGenerator  // 新增
};
```

```typescript
// 3. 编写测试
// test/unit/generators/kotlin/KotlinGenerator.test.ts
describe('KotlinGenerator', () => {
  let generator: KotlinGenerator;
  
  beforeEach(() => {
    generator = new KotlinGenerator();
  });
  
  describe('generateClass', () => {
    it('should generate simple data class', () => {
      const complexType = {
        name: 'Person',
        elements: [
          { name: 'firstName', type: 'xs:string', minOccurs: 1 },
          { name: 'lastName', type: 'xs:string', minOccurs: 1 },
          { name: 'age', type: 'xs:int', minOccurs: 0 }
        ]
      };
      
      const result = generator.generateClass(complexType);
      
      expect(result).toContain('data class Person');
      expect(result).toContain('val firstName: String');
      expect(result).toContain('val age: Int?');
    });
  });
});
```

### 2. 调试流程

#### 本地调试
```bash
# 运行开发版本
npm run dev

# 启动调试模式
npm run debug

# 运行特定测试
npm test -- --testNamePattern="KotlinGenerator"

# 生成覆盖率报告
npm run test:coverage
```

#### VS Code调试配置
```json
// .vscode/launch.json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug CLI",
      "type": "node",
      "request": "launch",
      "program": "${workspaceFolder}/dist/cli/index.js",
      "args": ["-i", "test/fixtures/simple.xsd", "-l", "typescript"],
      "outFiles": ["${workspaceFolder}/dist/**/*.js"],
      "sourceMaps": true
    },
    {
      "name": "Debug Tests",
      "type": "node",
      "request": "launch",
      "program": "${workspaceFolder}/node_modules/.bin/jest",
      "args": ["--runInBand"],
      "console": "integratedTerminal",
      "internalConsoleOptions": "neverOpen"
    }
  ]
}
```

### 3. 性能优化

#### 代码分析
```bash
# 分析bundle大小
npm run analyze

# 性能基准测试
npm run benchmark

# 内存使用分析
npm run profile
```

#### 性能测试示例
```typescript
// test/performance/generator.performance.test.ts
describe('Generator Performance', () => {
  it('should generate 1000 classes within 5 seconds', async () => {
    const schema = createLargeSchema(1000);
    const generator = new TypeScriptGenerator();
    
    const startTime = Date.now();
    const result = await generator.generate(schema, {});
    const duration = Date.now() - startTime;
    
    expect(duration).toBeLessThan(5000);
    expect(result.files).toHaveLength(1000);
  });
});
```

## 代码规范

### 1. TypeScript规范

#### 命名约定
```typescript
// 类名：PascalCase
class XSDParser {}
class JavaGenerator {}

// 接口名：PascalCase，以I开头（可选）
interface LanguageGenerator {}
interface IConfigManager {}

// 方法名、变量名：camelCase
const configManager = new ConfigManager();
function parseXSDFile() {}

// 常量：UPPER_SNAKE_CASE
const DEFAULT_TIMEOUT = 30000;
const MAX_FILE_SIZE = 100 * 1024 * 1024;

// 枚举：PascalCase
enum ValidationLevel {
  STRICT = 'strict',
  LOOSE = 'loose'
}
```

#### 类型定义
```typescript
// 优先使用接口而不是type别名
interface GenerationConfig {
  outputDir: string;
  language: string;
  options?: GenerationOptions;
}

// 使用联合类型
type SupportedLanguage = 'java' | 'typescript' | 'csharp' | 'python' | 'go';

// 使用泛型
interface Repository<T> {
  findById(id: string): Promise<T | null>;
  save(entity: T): Promise<T>;
}

// 使用严格的null检查
function processFile(path: string): string | null {
  if (!fs.existsSync(path)) {
    return null;
  }
  return fs.readFileSync(path, 'utf8');
}
```

#### 错误处理
```typescript
// 自定义错误类
export class XSDParseError extends Error {
  constructor(
    message: string,
    public readonly line?: number,
    public readonly column?: number
  ) {
    super(message);
    this.name = 'XSDParseError';
  }
}

// 使用Result模式
type Result<T, E = Error> = 
  | { success: true; data: T }
  | { success: false; error: E };

function parseXSD(content: string): Result<XSDSchema, XSDParseError> {
  try {
    const schema = doParseXSD(content);
    return { success: true, data: schema };
  } catch (error) {
    return { 
      success: false, 
      error: new XSDParseError(error.message) 
    };
  }
}
```

### 2. ESLint配置
```javascript
// .eslintrc.js
module.exports = {
  extends: [
    '@typescript-eslint/recommended',
    'prettier',
    'prettier/@typescript-eslint'
  ],
  parser: '@typescript-eslint/parser',
  plugins: ['@typescript-eslint'],
  rules: {
    '@typescript-eslint/explicit-function-return-type': 'error',
    '@typescript-eslint/no-explicit-any': 'warn',
    '@typescript-eslint/no-unused-vars': 'error',
    '@typescript-eslint/prefer-readonly': 'error',
    'prefer-const': 'error',
    'no-var': 'error'
  }
};
```

### 3. Prettier配置
```javascript
// .prettierrc.js
module.exports = {
  semi: true,
  trailingComma: 'es5',
  singleQuote: true,
  printWidth: 80,
  tabWidth: 2,
  useTabs: false
};
```

### 4. 注释规范
```typescript
/**
 * XSD解析器，用于将XSD文档解析为内部模型
 * 
 * @example
 * ```typescript
 * const parser = new XSDParser();
 * const schema = await parser.parseFile('schema.xsd');
 * ```
 */
export class XSDParser {
  /**
   * 解析XSD文件
   * 
   * @param filePath XSD文件路径
   * @returns 解析后的Schema对象
   * @throws {XSDParseError} 当文件格式无效时
   */
  async parseFile(filePath: string): Promise<XSDSchema> {
    // 实现逻辑
  }
  
  /**
   * 验证Schema的有效性
   * 
   * @param schema 要验证的Schema
   * @returns 验证结果
   * @internal 内部方法，不对外暴露
   */
  private validateSchema(schema: XSDSchema): ValidationResult {
    // 实现逻辑
  }
}
```

## 测试指南

### 1. 测试结构
```
test/
├── unit/                    # 单元测试
│   ├── core/
│   ├── generators/
│   └── utils/
├── integration/             # 集成测试
│   ├── cli/
│   └── end-to-end/
├── fixtures/                # 测试数据
│   ├── schemas/
│   ├── configs/
│   └── expected/
└── helpers/                 # 测试工具
    ├── builders/
    └── matchers/
```

### 2. 测试框架配置
```javascript
// jest.config.js
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  roots: ['<rootDir>/src', '<rootDir>/test'],
  testMatch: ['**/*.test.ts'],
  collectCoverageFrom: [
    'src/**/*.ts',
    '!src/**/*.d.ts',
    '!src/**/index.ts'
  ],
  coverageThreshold: {
    global: {
      branches: 80,
      functions: 80,
      lines: 80,
      statements: 80
    }
  },
  setupFilesAfterEnv: ['<rootDir>/test/helpers/setup.ts']
};
```

### 3. 测试示例

#### 单元测试
```typescript
// test/unit/core/parser/XSDParser.test.ts
import { XSDParser } from '../../../../src/core/parser/XSDParser';
import { readFixture } from '../../../helpers/fixtures';

describe('XSDParser', () => {
  let parser: XSDParser;
  
  beforeEach(() => {
    parser = new XSDParser();
  });
  
  describe('parse', () => {
    it('should parse simple XSD schema', () => {
      const xsdContent = readFixture('schemas/simple.xsd');
      
      const result = parser.parse(xsdContent);
      
      expect(result.targetNamespace).toBe('http://example.com/schema');
      expect(result.complexTypes).toHaveLength(2);
      expect(result.complexTypes[0].name).toBe('Person');
    });
    
    it('should throw error for invalid XSD', () => {
      const invalidXsd = '<invalid>xml</invalid>';
      
      expect(() => parser.parse(invalidXsd)).toThrow(XSDParseError);
    });
  });
  
  describe('parseFile', () => {
    it('should parse XSD from file', async () => {
      const filePath = 'test/fixtures/schemas/person.xsd';
      
      const result = await parser.parseFile(filePath);
      
      expect(result).toBeDefined();
      expect(result.complexTypes).toHaveLength(1);
    });
    
    it('should reject for non-existent file', async () => {
      const filePath = 'non-existent.xsd';
      
      await expect(parser.parseFile(filePath)).rejects.toThrow();
    });
  });
});
```

#### 集成测试
```typescript
// test/integration/cli/generate.test.ts
import { execSync } from 'child_process';
import { tmpdir } from 'os';
import { join } from 'path';
import { readFileSync, mkdirSync } from 'fs';

describe('CLI Integration', () => {
  let tempDir: string;
  
  beforeEach(() => {
    tempDir = join(tmpdir(), 'xsd2code-test-' + Date.now());
    mkdirSync(tempDir, { recursive: true });
  });
  
  it('should generate TypeScript code from XSD', () => {
    const xsdPath = 'test/fixtures/schemas/order.xsd';
    const command = `node dist/cli/index.js -i ${xsdPath} -l typescript -o ${tempDir}`;
    
    const result = execSync(command, { encoding: 'utf8' });
    
    expect(result).toContain('Generated successfully');
    
    const generatedFile = join(tempDir, 'Order.ts');
    const content = readFileSync(generatedFile, 'utf8');
    
    expect(content).toContain('export interface Order');
    expect(content).toContain('orderId: string');
  });
});
```

#### 端到端测试
```typescript
// test/e2e/full-workflow.test.ts
describe('Full Workflow E2E', () => {
  it('should complete full generation workflow', async () => {
    // 1. 准备XSD文件
    const xsdPath = 'test/fixtures/schemas/ecommerce.xsd';
    
    // 2. 准备配置
    const config = {
      languages: [
        { language: 'typescript', package: 'models' },
        { language: 'java', package: 'com.example.models' }
      ]
    };
    
    // 3. 执行生成
    const generator = new XSD2CodeGenerator();
    const result = await generator.generateFromConfig(xsdPath, config);
    
    // 4. 验证结果
    expect(result.success).toBe(true);
    expect(result.files).toHaveLength(10); // 预期生成10个文件
    
    // 5. 验证生成的代码能够编译
    const tsFiles = result.files.filter(f => f.path.endsWith('.ts'));
    for (const file of tsFiles) {
      const compileResult = await compileTypeScript(file.content);
      expect(compileResult.success).toBe(true);
    }
    
    const javaFiles = result.files.filter(f => f.path.endsWith('.java'));
    for (const file of javaFiles) {
      const compileResult = await compileJava(file.content);
      expect(compileResult.success).toBe(true);
    }
  });
});
```

### 4. 测试工具

#### 测试数据构建器
```typescript
// test/helpers/builders/XSDSchemaBuilder.ts
export class XSDSchemaBuilder {
  private schema: Partial<XSDSchema> = {};
  
  withTargetNamespace(namespace: string): this {
    this.schema.targetNamespace = namespace;
    return this;
  }
  
  withComplexType(name: string, elements: XSDElement[]): this {
    if (!this.schema.complexTypes) {
      this.schema.complexTypes = [];
    }
    this.schema.complexTypes.push({ name, elements });
    return this;
  }
  
  build(): XSDSchema {
    return {
      targetNamespace: this.schema.targetNamespace || 'http://example.com',
      complexTypes: this.schema.complexTypes || [],
      simpleTypes: this.schema.simpleTypes || [],
      elements: this.schema.elements || []
    };
  }
}

// 使用示例
const schema = new XSDSchemaBuilder()
  .withTargetNamespace('http://test.com')
  .withComplexType('Person', [
    { name: 'firstName', type: 'xs:string' },
    { name: 'lastName', type: 'xs:string' }
  ])
  .build();
```

#### 自定义匹配器
```typescript
// test/helpers/matchers/toGenerateValidCode.ts
declare global {
  namespace jest {
    interface Matchers<R> {
      toGenerateValidCode(language: string): R;
    }
  }
}

expect.extend({
  toGenerateValidCode(received: string, language: string) {
    const isValid = validateGeneratedCode(received, language);
    
    return {
      message: () =>
        `expected generated code to be valid ${language} code`,
      pass: isValid
    };
  }
});

// 使用示例
expect(generatedCode).toGenerateValidCode('typescript');
```

## 贡献流程

### 1. 贡献类型

#### Bug修复
1. 在GitHub Issues中查找相关bug报告
2. 如果没有，创建新的issue描述bug
3. Fork项目并创建修复分支
4. 编写测试复现bug
5. 修复bug并确保测试通过
6. 提交Pull Request

#### 新功能开发
1. 在GitHub Issues中讨论功能需求
2. 等待维护者确认后开始开发
3. 遵循功能开发流程
4. 确保有充分的测试覆盖
5. 更新相关文档
6. 提交Pull Request

#### 文档改进
1. 直接提交小的文档修正
2. 大的文档改动建议先讨论
3. 确保文档格式正确
4. 检查链接和示例的有效性

### 2. Pull Request流程

#### 提交前检查
```bash
# 运行所有测试
npm test

# 检查代码格式
npm run lint

# 修复格式问题
npm run lint:fix

# 检查类型
npm run type-check

# 构建项目
npm run build
```

#### PR描述模板
```markdown
## 变更类型
- [ ] Bug修复
- [ ] 新功能
- [ ] 性能改进
- [ ] 文档更新
- [ ] 重构

## 变更描述
简要描述你的变更...

## 测试
- [ ] 添加了新的测试
- [ ] 所有现有测试通过
- [ ] 手动测试通过

## 清单
- [ ] 代码遵循项目规范
- [ ] 自测试通过
- [ ] 文档已更新
- [ ] CHANGELOG已更新（如适用）

## 相关Issue
关闭 #123
```

#### 代码审查

**审查要点：**
1. **功能正确性**：代码是否实现了预期功能
2. **代码质量**：是否遵循项目规范
3. **测试覆盖**：是否有充分的测试
4. **性能影响**：是否对性能有负面影响
5. **向后兼容**：是否破坏现有API
6. **安全性**：是否引入安全漏洞

**审查流程：**
1. 自动化检查（CI/CD）
2. 代码审查（至少一个维护者）
3. 测试验证
4. 文档检查
5. 合并到main分支

### 3. 版本发布

#### 语义化版本
- **主版本 (Major)**：不兼容的API变更
- **次版本 (Minor)**：向后兼容的功能新增
- **修订版本 (Patch)**：向后兼容的bug修复

#### 发布流程
```bash
# 1. 更新版本号
npm version patch  # 或 minor, major

# 2. 更新CHANGELOG
# 手动编辑CHANGELOG.md

# 3. 提交版本变更
git add .
git commit -m "chore: release v1.2.3"

# 4. 创建标签
git tag v1.2.3

# 5. 推送到远程
git push origin main --tags

# 6. 发布到npm
npm publish

# 7. 创建GitHub Release
# 在GitHub上创建发布说明
```

#### 自动化发布
```yaml
# .github/workflows/release.yml
name: Release
on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-node@v3
        with:
          node-version: '16'
          registry-url: 'https://registry.npmjs.org'
      
      - name: Install dependencies
        run: npm ci
      
      - name: Run tests
        run: npm test
      
      - name: Build
        run: npm run build
      
      - name: Publish to npm
        run: npm publish
        env:
          NODE_AUTH_TOKEN: ${{ secrets.NPM_TOKEN }}
      
      - name: Create GitHub Release
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          tag_name: ${{ github.ref }}
          release_name: Release ${{ github.ref }}
          draft: false
          prerelease: false
```

## 社区参与

### 1. 沟通渠道
- **GitHub Issues**: 问题报告和功能请求
- **GitHub Discussions**: 一般讨论和问答
- **Discord**: 实时聊天（如果有）
- **邮件列表**: 重要公告

### 2. 贡献者指南
- 遵循[行为准则](CODE_OF_CONDUCT.md)
- 尊重其他贡献者
- 保持友好和建设性的讨论
- 帮助新人融入社区

### 3. 认可贡献者
- Contributors页面列出所有贡献者
- 重要贡献者会被邀请成为维护者
- 定期发布贡献者感谢帖

## 总结

参与XSD2Code项目开发是一个很好的学习和贡献开源社区的机会。无论是修复bug、添加新功能，还是改进文档，每一个贡献都是有价值的。

通过遵循本指南中的流程和规范，您可以确保您的贡献能够顺利被接受并集成到项目中。我们期待您的参与！

### 快速开始贡献
1. Fork项目
2. 克隆到本地
3. 安装依赖：`npm install`
4. 运行测试：`npm test`
5. 创建功能分支：`git checkout -b feature/your-feature`
6. 进行开发
7. 提交Pull Request

如果在开发过程中遇到任何问题，请随时在GitHub Issues中提问，我们会尽快回复并提供帮助。
