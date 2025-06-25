FROM golang:1.23.0-alpine as build-env

RUN apk add --no-cache git ca-certificates

ARG GITHUB_TOKEN
ARG VERSION=unknown
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux TOKEN=$GITHUB_TOKEN VERSION=${VERSION}

RUN git config --global url."https://${TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /go/src/github.com/RodolfoBonis/microdetect-api/

COPY go.mod /go/src/github.com/RodolfoBonis/microdetect-api/
COPY go.sum /go/src/github.com/RodolfoBonis/microdetect-api/

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

ADD . /go/src/github.com/RodolfoBonis/microdetect-api/

COPY . ./


FROM golang:1.23.0 as builder

ARG GITHUB_TOKEN
ARG VERSION=unknown
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux TOKEN=$GITHUB_TOKEN VERSION=${VERSION}

WORKDIR /go/src/github.com/RodolfoBonis/microdetect-api/

COPY --from=build-env /go/src/github.com/RodolfoBonis/microdetect-api /go/src/github.com/RodolfoBonis/microdetect-api/

COPY --from=build-env /go/bin/swag /

RUN /swag init

RUN git config --global url."https://${TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go env -w GOPRIVATE=github.com/RodolfoBonis/go_key_guardian

RUN go get github.com/RodolfoBonis/go_key_guardian

RUN go build -o microdetect-api -v /go/src/github.com/RodolfoBonis/microdetect-api/

COPY . ./

FROM alpine:3.15 AS production

ARG GITHUB_TOKEN
ARG VERSION=unknown
ENV VERSION=${VERSION}

WORKDIR /go/src/github.com/RodolfoBonis/microdetect-api/

COPY --from=builder /go/src/github.com/RodolfoBonis/microdetect-api/microdetect-api /

CMD ["/microdetect-api"]

EXPOSE 8000