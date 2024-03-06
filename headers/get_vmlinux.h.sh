#!/bin/bash

# 定义一个函数来下载和安装 bpftool
install_bpftool() {
    # 获取最新的 bpftool 版本
    local BPFTOOL_VERSION=$(curl -s https://api.github.com/repos/libbpf/bpftool/releases/latest | jq -r '.tag_name')

    if [ -z "$BPFTOOL_VERSION" ]; then
        echo "Failed to fetch the latest bpftool version"
        exit 1
    fi

    echo "Latest bpftool version: $BPFTOOL_VERSION"

    # 检测系统架构
    local ARCH=$(uname -m)
    local BPFTOOL_ARCH=""
    case $ARCH in
        x86_64)
            BPFTOOL_ARCH="amd64"
            ;;
        aarch64)
            BPFTOOL_ARCH="arm64"
            ;;
        *)
            echo "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac

    # 构建下载 URL
    local BPFTOOL_URL="https://github.com/libbpf/bpftool/releases/download/${BPFTOOL_VERSION}/bpftool-${BPFTOOL_VERSION}-${BPFTOOL_ARCH}.tar.gz"

    # 下载 bpftool
    echo "Downloading bpftool for ${ARCH}..."
    wget -O bpftool-${BPFTOOL_VERSION}-${BPFTOOL_ARCH}.tar.gz $BPFTOOL_URL

    if [ $? -ne 0 ]; then
        echo "Download failed!"
        exit 1
    else
        echo "Download successful!"
    fi

    # 解压下载的文件
    echo "Extracting bpftool..."
    tar -xzf bpftool-${BPFTOOL_VERSION}-${BPFTOOL_ARCH}.tar.gz

    if [ $? -ne 0 ]; then
        echo "Extraction failed!"
        exit 1
    else
        echo "Extraction successful!"
    fi

    # 可以选择移动 bpftool 到系统路径中
    sudo mv bpftool /usr/local/bin/

    echo "bpftool is ready to use."

    rm bpftool-${BPFTOOL_VERSION}-${BPFTOOL_ARCH}.tar.gz
}

# 主流程
# 检查 bpftool 是否已安装
if ! command -v bpftool &> /dev/null; then
    echo "bpftool could not be found, installing..."
    install_bpftool
else
    echo "bpftool is already installed."
    sudo bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
    echo "vmlinux dumped successfully"
fi


