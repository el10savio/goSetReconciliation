FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /set

WORKDIR /set

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/set


FROM scratch

COPY --from=builder /go/bin/set /go/bin/set

ENTRYPOINT ["/go/bin/set"]

EXPOSE 8080
