FROM wynn/monstache-builder-cache-rel6:latest AS build-app

RUN mkdir /app

WORKDIR /app

COPY . .

COPY /docker/plugin/${PLUGIN}.go /app/docker/plugin/

RUN go mod download

ARG PLUGIN

WORKDIR /app/docker/plugin

RUN go mod download

# Compile the specific plugin file
RUN go build -buildmode=plugin -o ${PLUGIN}.so ./${PLUGIN}.go
