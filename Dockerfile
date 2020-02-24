# build stage
FROM golang:1 AS builder
#ADD . /src
COPY / /src/
WORKDIR /src/
RUN cd /src/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /newsapi

# final stage
FROM alpine:latest
#RUN apk --no-cache add ca-certificates
COPY --from=builder /newsapi ./newsapi
RUN chmod +x ./newsapi
ENTRYPOINT ["./newsapi"]
EXPOSE 80