ARG arch
FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED="0" GOARCH=${arch} go build -ldflags="-s -w" -a -o App


ENTRYPOINT [ "./App" ]
