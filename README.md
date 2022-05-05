# caribou-next

Caribou Project is used to search and display Pixiv illustration.

This repository is the backend server, which provides illustration information through APIs for caribou. The frontend part is stored in a separate repository [caribou-web](https://github.com/YKMeIz/caribou-web). It is required to bundle [caribou-web](https://github.com/YKMeIz/caribou-web) and [caribou-next](https://github.com/YKMeIz/caribou-next) together.

> Please be aware that you may want to modify source if you plan to put this project into your production. For example, changing the listened port.

## Project setup
```
go mod tidy
```

## Compiles for Production
```
go build -o caribou-next main.go
```
> This repository comes with a pre-build front-end website for your convenience.
> In deployment, please copy `dist` folder accompany with executable golang binary file.

### Production Scenario

The website may run as a virtual host served by docker container instance.

Here is a sample of Dockerfile:
```dockerfile
FROM golang:alpine

RUN apk add --no-cache tzdata

WORKDIR /
COPY ./caribou-next .
ADD ./dist/ ./dist/

EXPOSE 9090/tcp

ENTRYPOINT ["/caribou-next"]
```
