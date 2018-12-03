build:
	# 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统
	GOOS=darwin GOARCH=amd64 go build
	docker build -t sensitiveword .
run:
	# 可添加 -d 参数将微服务放到后台运行
	docker run -p 8080:8080 sensitiveword

