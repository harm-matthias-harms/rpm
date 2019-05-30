# RPM - Backend
The backend of the RPM Application is placed here.

## Install dependencies
Simple run:
```
go mod download
```

## Test
```
go test ./...
```
With coverage
```
go test ./... --cover
```

## Build
Generate a binary:
```
go build
```
Build the docker container:
```
docker build -t rpm-backend:latest .
```

## Run
Don't run it, use tests. If you want to run it local just build a docker container or use `go build`.

Run the container:
```
docker run --rm -p 3000:3000 rpm-backend:latest
```