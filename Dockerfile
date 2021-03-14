FROM golang
WORKDIR /go/src/github.com/shambanna-u/prome-go-lib
RUN go get -d -v github.com/shambanna-u/prome-go-lib
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/github.com/shambanna-u/prome-go-lib/prome-go-lib .
EXPOSE 2222
CMD ["./app"]  

# FROM golang
#  RUN go get -u github.com/shambanna-u/prome-go-lib
#  ENV GO111MODULE=on
#  ENV GOFLAGS=-mod=vendor
#  ENV APP_HOME /go/src/mathapp
#  RUN mkdir -p $APP_HOME
#  WORKDIR $APP_HOME
#  EXPOSE 2222
# #  CMD ["/prome-go-lib"]
