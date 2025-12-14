# 设置 GOBIN 环境变量，默认为 $(go env GOPATH)/bin
# ?= 表示如果变量未定义才赋值，$$ 是为了在 Makefile 中转义 $
GOBIN ?= $$(go env GOPATH)/bin

# .PHONY 声明伪目标，表示这些不是真实的文件名，每次都需要执行
# 安装测试覆盖率检查工具
.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

# 检查测试覆盖率是否达标
# 依赖 install-go-test-coverage，会先安装工具
.PHONY: check-coverage
check-coverage: install-go-test-coverage
	# 运行测试并生成覆盖率报告到 cover.out
	# -covermode=atomic 支持并发安全的覆盖率统计
	# -coverpkg=./... 统计所有包的覆盖率
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	# 使用配置文件检查覆盖率是否达到要求
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml

# 运行测试
.PHONY: test
test:
	# -v 显示详细输出，-race 开启竞态检测
	go test -v -race ./...

# 生成覆盖率报告并在终端显示
.PHONY: cover
cover:
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	# 在终端按函数显示覆盖率
	go tool cover -func=./cover.out

# 生成覆盖率报告并在浏览器中打开 HTML 可视化页面
.PHONY: cover-html
cover-html:
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	# 生成 HTML 报告并自动在浏览器中打开
	go tool cover -html=./cover.out

# 检查每个目录下非测试 Go 文件数量不超过 20 个
.PHONY: check-file-count
check-file-count:
	@bash scripts/check-file-count.sh 20

# 运行代码静态检查（需要先安装 golangci-lint）
.PHONY: lint
lint:
	golangci-lint run

# 清理生成的覆盖率文件
.PHONY: clean
clean:
	rm -f cover.out coverage.out

