#!/bin/bash
set -e

# 环境变量，darwin/arm64 架构，CGO 启用
export GOOS=darwin
export GOARCH=arm64
export CGO_ENABLED=1

# 你的主程序源码路径（根据实际改）
MAIN_PKG="./monstache.go"
# 编译输出主程序文件名
MAIN_BIN="monstache"

# 你的插件源码路径（根据实际改）
#PLUGIN_SRC="./collection_process.go"
# 编译输出插件文件名
#PLUGIN_SO="collection_process.so"

echo "开始编译主程序..."
go build -o $MAIN_BIN $MAIN_PKG
echo "主程序编译完成: $MAIN_BIN"

echo "开始编译插件..."
#go build -buildmode=plugin -o $PLUGIN_SO $PLUGIN_SRC
#echo "插件编译完成: $PLUGIN_SO"

echo "编译全部完成！"

