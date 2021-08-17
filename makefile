.PHONY: build clean tool lint help
#声明 build / clean / tool / lint / help 为伪目标
all: build
#make 就是 make all 
#make build: 编译当前项目的包和依赖项

build:
    go build -v .
#make tool: 运行指定的 Go 工具集
tool:
    go tool vet . |& grep -v vendor; true
    gofmt -w .	
#make lint: golint 一下
lint:
    golint ./...
#make clean: 删除对象文件和缓存文件
clean:
    rm -rf FreshmanGuidanceProject
    go clean -i .

help:
    @echo "make: compile packages and dependencies"
    @echo "make tool: run specified go tool"
    @echo "make lint: golint ./..."
    @echo "make clean: remove object files and cached files"