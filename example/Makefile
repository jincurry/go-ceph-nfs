# 设置 Go 编译器
GO := go

# 设置二进制文件名
BINARY_NAME := go_ceph_nfs

# 设置源文件目录
SRC_DIR := .

# 设置构建标记和编译选项
LDFLAGS := -s -w

.PHONY: all build clean install

# 默认目标
all: build

# 构建目标
build:
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) $(SRC_DIR)

# 清理目标
clean:
	rm -f $(BINARY_NAME)
	$(GO) clean

# 安装目标
install:
	$(GO) install

# 运行测试
test:
	$(GO) test ./...
