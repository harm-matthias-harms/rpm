# frontend


## Build Setup

``` bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn run dev

# build for production and launch server
$ yarn run build
$ yarn start

# generate static project
$ yarn run generate

# test with jest
$ yarn run test

# test with cypress
$ yarn run test:e2e

# test with cypress and  open it
$ yarn run test:e2e:open

# Build with docker
$ docker build -t rpm-frontend .

# Run docker container
$ docker run --rm -p 3000:80 rpm-frontend:latest
```

## Options

| ENV_VAR 	| description        	|
|---------	|--------------------	|
| API_URL 	| url of the backend 	|

For detailed explanation on how things work, checkout [Nuxt.js docs](https://nuxtjs.org).
