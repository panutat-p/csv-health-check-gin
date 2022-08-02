# csv-health-check-gin

## Requirements
* ✅ Must be written in Golang 
* ✅ Check website available with net/http package; Up: any http status code that is returned from the website; Down: cannot reach the website (request timeout and not found)
* ✅ Handle errors without stopping the entire process 
* ✅ Write readable code 
* ✅ Write automated unit tests 
* ✅ Run program as fast as possible on multi-core CPU
* 🟨 Bonus if you have a good package structure 
* Bonus if you can handle a vast file

## Note
* You can run binary file without building the project
* Build & run docker container on port 8080
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

GOOS=windows GOARCH=amd64 go build \
    -ldflags "-X main.buildTime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
    -o ./bin/app.exe
```

```
.
├── health
│   └── handler.go
├── logging
│   └── color.go
├── main.go
└── upload
    ├── handler.go
    ├── models.go
    └── service
        ├── csv.go
        ├── csv_test.go
        ├── http.go
        └── http_test.go

```
