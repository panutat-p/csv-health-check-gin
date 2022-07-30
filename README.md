# line-remoteinterview-gin

## Requirements
* ✅ Must be written in Golang 
* ✅ Check website available with net/http package; Up: any http status code that is returned from the website; Down: cannot reach the website (request timeout and not found)
* ✅ Handle errors without stopping the entire process 
* ✅ Write readable code 
* 🟧 Write automated unit tests 
* ✅ Run program as fast as possible on multi-core CPU
* ❓ Bonus if you have a good package structure 
* Bonus if you can handle a vast file

## Note
* timeout for HTTP client is 3s
* slow upload does not affect this timeout

## Run

```shell
docker image build -t gin .
docker container run -p 8080:8080 --name gin gin
```

```shell
go build \
    -ldflags "-X main.buildTime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
    -o ./bin/app
```
