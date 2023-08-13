FROM cgr.dev/chainguard/go:1.20 AS builder
COPY . /app
RUN cd /app && go build -o server .

FROM cgr.dev/chainguard/glibc-dynamic:latest
COPY --from=builder /app/server /usr/bin/
CMD ["/usr/bin/server"]
