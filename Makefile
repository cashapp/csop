APP_NAME:=csop
GIT_COMMIT?=$(shell git rev-parse --short HEAD)
DOCKER_TAG:=${APP_NAME}:${GIT_COMMIT}

lint:
	golangci-lint run -c ./.golangci.yml

test:
	go test -mod=vendor ./...

integration-test:
	operator-sdk test local

image:
	operator-sdk build ${DOCKER_TAG} --verbose --image-build-args="--build-arg SERVICE=\"${APP_NAME}\" --build-arg GIT_COMMIT=\"${GIT_COMMIT}\""

clean:
	go clean
