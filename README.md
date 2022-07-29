# line-remoteinterview-gin

```shell
go mod init github.com/panutat-p/line-remoteinterview-gin
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors
```

```shell
go build \
		-ldflags "-X main.buildCommit=`git rev-parse --short HEAD` \
		-X main.buildTime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
		-o ./bin/app
```

```shell
docker image build -t gin .

docker container run -p 8080:8080 --name gin gin
```
