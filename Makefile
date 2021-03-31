.PHONY: build docker-build

all: build

build:
	go build main.go

docker-build:
	docker build -t dpakach/kube-go .

docker-push:
	docker push dpakach/kube-go

docker-run:
	docker run -p 8080:8080 -e SIMPLE_ENV='simple_value' -e CONFIG_VALUE='config' -e SECRET_PASSWORD='Secret' dpakach/kube-go
