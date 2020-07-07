# build stage
FROM golang:alpine AS build-env
ENV GOPROXY=direct
RUN apk --no-cache add build-base git bzr mercurial gcc protoc protobuf protobuf-dev
RUN go get -u github.com/golang/protobuf/protoc-gen-go && \
  go get -u google.golang.org/grpc
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o server example/cmd/main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/app/server /app/
COPY scripts/entrypoint.sh /app/
RUN chmod +x ./entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
CMD []
