FROM golang:1.19-alpine3.18 as buildbase

WORKDIR /go/src/github.com/apodeixis/mailjet-svc

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/mailjet-svc main.go

FROM alpine:3.18

COPY --from=buildbase /usr/local/bin/mailjet-svc /usr/local/bin/mailjet-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["mailjet-svc"]
