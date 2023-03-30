FROM golang:alpine as base
COPY . /app
WORKDIR /app

FROM base as test
RUN ["go","test","./..."]

FROM base as run
CMD go run main.go
