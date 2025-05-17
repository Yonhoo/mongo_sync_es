#!/bin/bash

set -e

# 构建镜像
docker build -f Dockerfile -t monstache-release .

# 启动一个容器（后台运行，保持不退出）
docker run --rm -d --name monstache-release-container monstache-release tail -f /dev/null

# 如果 docker-build 目录存在则删除
if [ -d docker-build ]; then
  rm -rf docker-build
fi
mkdir docker-build

# 拷贝可执行文件到本地
docker cp monstache-release-container:/bin/monstache ./docker-build/monstache

# 停止并移除容器
docker stop monstache-release-container

echo "Build complete! The binary is in ./docker-build/monstache"