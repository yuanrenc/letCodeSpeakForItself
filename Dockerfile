ARG arch 
FROM --platform=linux/amd64 golang:alpine AS build

WORKDIR /letcodespeakforitself

COPY . .

RUN go mod tidy

RUN CGO_ENABLED="0" GOARCH=${arch} go build -ldflags="-s -w" -a -o App


FROM --platform=linux/amd64 alpine:latest

WORKDIR /App

COPY config.yaml ./config.yaml
COPY --from=build /letcodespeakforitself/App App

ENTRYPOINT [ "./App" ]
