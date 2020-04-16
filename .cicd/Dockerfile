FROM golang:1.14-alpine AS builder
WORKDIR /build
ADD . /build

# "github.com/mattn/go-sqlite3" needs gcc (libc-dev for headers)
#RUN apk add git protobuf gcc libc-dev

#RUN for p in github.com/golang/protobuf/proto \
#    "github.com/golang/protobuf/protoc-gen-go" \
#    "github.com/jinzhu/gorm" \
#    "github.com/mattn/go-sqlite3" \
#    "google.golang.org/grpc" \
#    "google.golang.org/grpc/credentials" \
#    "google.golang.org/grpc/testdata" \
#    ; do GO111MODULES=off go get ${p}; done

#RUN protoc -Iproto --go_out=plugins=grpc:grpc proto/dailylog.proto

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o gopherlol

FROM scratch
COPY --from=builder /build/gopherlol /gopherlol
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/gopherlol"]
#CMD ["-port", "4000"]
#HEALTHCHECK --interval=5m --timeout=3s CMD /dl -buildinfo || exit 1
