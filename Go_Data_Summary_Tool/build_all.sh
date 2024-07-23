#!/bin/bash

# 设置项目根目录（相对于脚本的位置）
PROJECT_ROOT="$(dirname "$0")"
cd "$PROJECT_ROOT"

# 创建或清空releases目录
mkdir -p releases
rm -f releases/*

# 定义函数来构建并移动编译后的二进制文件
function build_and_move() {
    local BINARY_NAME="GDST"  # 替换为你的程序名
    local OS=$1
    local ARCH=$2
    local OUTPUT_NAME

    # 根据操作系统设置输出文件名
    if [ "$OS" = "windows" ]; then
        OUTPUT_NAME="${BINARY_NAME}_${OS}_${ARCH}.exe"
    else
        OUTPUT_NAME="${BINARY_NAME}_${OS}_${ARCH}"
    fi

    # 设置环境变量
    export GOOS=$OS
    export GOARCH=$ARCH

    # 编译
    echo "Building $OUTPUT_NAME..."
    go build -o "./releases/$OUTPUT_NAME"

    # 检查是否编译成功
    if [ $? -ne 0 ]; then
        echo "Error building $OUTPUT_NAME"
        exit 1
    fi

    echo "$OUTPUT_NAME built successfully."
}

# MacOS 架构
build_and_move darwin amd64
build_and_move darwin arm64

# Linux 架构
build_and_move linux amd64
build_and_move linux arm64
build_and_move linux arm
build_and_move linux 386  # 32位x86

# Windows 架构
build_and_move windows amd64
build_and_move windows 386  # 32位x86

echo "All binaries have been built and moved to the releases directory."