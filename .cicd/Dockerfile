FROM golang:1.14-alpine AS builder
WORKDIR /build
ADD . /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o gopherlol


FROM scratch
COPY --from=builder /build/gopherlol /gopherlol
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/gopherlol"]
#CMD ["-port", "4000"]
#HEALTHCHECK --interval=5m --timeout=3s CMD /dl -buildinfo || exit 1
