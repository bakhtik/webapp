FROM golang:latest as builder
RUN go get -u -v github.com/go-redis/redis
RUN go get -u -v github.com/denisenkom/go-mssqldb
ADD . /go/src/github.com/bakhtik/webapp
WORKDIR /go/src/github.com/bakhtik/webapp
RUN go build -o main cmd/webapp/main.go

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# FROM scratch
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY public/ /app/public
# COPY templates/ /app/templates
EXPOSE 8080
CMD ["./main"]