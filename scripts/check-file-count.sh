#!/bin/bash
# 检查每个目录下非测试 Go 文件数量不超过指定上限

MAX_FILES=${1:-20}  # 默认上限 20，可通过参数传入

has_error=0

# 遍历所有包含 .go 文件的目录
for dir in $(find . -name '*.go' -not -path './vendor/*' -not -path './.git/*' -exec dirname {} \; | sort -u); do
    # 统计非测试文件数量
    count=$(find "$dir" -maxdepth 1 -name '*.go' ! -name '*_test.go' -type f | wc -l)
    
    if [ "$count" -gt "$MAX_FILES" ]; then
        echo "Error: $dir has $count non-test .go files (max $MAX_FILES)"
        has_error=1
    fi
done

if [ "$has_error" -eq 1 ]; then
    exit 1
fi

echo "✓ File count check passed (max $MAX_FILES per directory)"

