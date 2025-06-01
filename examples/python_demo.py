#!/usr/bin/env python3

"""
xsd2code Python Generator Demo
演示 Python 代码生成功能
"""

import sys
import os

def main():
    print("=== XSD2Code Python 代码生成演示 ===")
    print()
    
    # 检查是否存在 xsd2code 可执行文件
    if not os.path.exists("./xsd2code"):
        print("错误: 未找到 xsd2code 可执行文件")
        print("请先编译项目: go build -o xsd2code ./cmd")
        return 1
    
    # 使用示例 XSD 文件
    xsd_file = "test/simple_types.xsd"
    if not os.path.exists(xsd_file):
        xsd_file = "examples/person.xsd"
    
    if not os.path.exists(xsd_file):
        print(f"错误: 未找到测试 XSD 文件")
        return 1
    
    # 生成 Python 代码
    output_file = "examples/python_demo.py"
    print(f"1. 生成 Python 代码...")
    print(f"   输入: {xsd_file}")
    print(f"   输出: {output_file}")
    
    # 执行命令
    cmd = f"./xsd2code -xsd={xsd_file} -lang=python -output={output_file} -package=models"
    print(f"   命令: {cmd}")
    
    result = os.system(cmd)
    if result == 0:
        print("✓ Python 代码生成成功")
        
        # 显示生成的代码
        if os.path.exists(output_file):
            print(f"\n生成的 Python 代码:")
            print("=" * 50)
            with open(output_file, 'r', encoding='utf-8') as f:
                print(f.read())
            print("=" * 50)
    else:
        print("✗ Python 代码生成失败")
        return 1
    
    # 显示类型映射
    print(f"\n2. Python 类型映射:")
    mapping_cmd = f"./xsd2code -xsd={xsd_file} -lang=python -show-mappings"
    print(f"   命令: {mapping_cmd}")
    os.system(mapping_cmd)
    
    print(f"\n=== 演示完成 ===")
    return 0

if __name__ == "__main__":
    sys.exit(main())
