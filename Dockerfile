# ------------------------------------------------------------------------------
# build
# ------------------------------------------------------------------------------
FROM golang:1.16.3 AS builder

LABEL NAME golang

WORKDIR /app

COPY . .

RUN go mod download \
    && cd cmd \
    && go build -o /bin/go-pihole \
    && useradd go-pihole

# ------------------------------------------------------------------------------
# daemon image
# ------------------------------------------------------------------------------
FROM scratch AS runner

USER go-pihole

COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/go-pihole /bin/go-pihole

