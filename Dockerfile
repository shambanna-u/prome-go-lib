FROM golang
WORKDIR /go/src/github.com/shambanna-u/prome-go-lib
RUN go get -d -v github.com/shambanna-u/prome-go-lib
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/shambanna-u/prome-go-lib/app .
EXPOSE 2222
CMD ["./app"]
