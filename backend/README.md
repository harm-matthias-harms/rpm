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

## Options

| ENV_VAR 	    | description        	            |
|---------	    |--------------------	            |
| DOMAIN 	    | domain to allow origins from 	    |
| JWT_SECRET    | JWT secret to use (keep it secret)|
| MONGO_URL     | url to reach mongo db 	        |
| MONGO_DATABASE| name of the database 	    |