PROJECT:=herman

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix -o herman .

# make build-linux
build-linux:
	@docker build -t herman:latest .
	@echo "build successful"

# make run
run:
    # delete herman-api container
	@if [ $(shell docker ps -aq --filter name=herman --filter publish=8000) ]; then docker rm -f herman; fi

    # 启动方法一 run herman-api container  docker-compose 启动方式
    # 进入到项目根目录 执行 make run 命令
	@docker-compose up -d

	# 启动方式二 docker run  这里注意-v挂载的宿主机的地址改为部署时的实际决对路径
    #@docker run --name=herman -p 8000:8000 -v C:\Users\86150\Desktop\Golang\herman:/code/herman/herman-api -d --restart=always herman:latest

	@echo "herman service is running···"

	# delete Tag=<none> 的镜像
	@docker image prune -f
	@docker ps -a | grep "herman"

# make stop
stop:
    # delete herman-api container
	@if [ $(shell docker ps -aq --filter name=herman --filter publish=8000) ]; then docker-compose down; fi

# make deploy
deploy:
	make build-linux
	make run