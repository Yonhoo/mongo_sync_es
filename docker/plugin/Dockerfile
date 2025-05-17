FROM rwynn/monstache-builder-cache-rel6:1.0.8 AS build-app

ARG PLUGIN


RUN mkdir /app

WORKDIR /app

COPY . .



RUN go mod download



WORKDIR /app/docker/plugin

RUN go mod download

# Compile the specific plugin file
RUN GOOS=darwin CGO_ENABLED=1 GOARCH=arm64 go build -buildmode=plugin -o ${PLUGIN}.so ./${PLUGIN}.go
