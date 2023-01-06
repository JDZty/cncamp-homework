export tag=myserver-1.0
export docker_repo=jdzzty/cncamp
export os=linux
export arch=amd64
root:
        export ROOT=github.com/cncamp-homework/httpserver
build:
        echo "======== build httpserver binary ========"
        mkdir -p bin/amd64
        CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -o bin/amd64
release: build
        echo "========= build httserver image ========="
        docker build -t ${docker_repo}:${tag} .
push: release
        echo "======== push httpserver image to docker repo ========"
        docker push ${docker_repo}:${tag}