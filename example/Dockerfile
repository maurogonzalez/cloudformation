FROM alpine:3.11

COPY s3example .

RUN apk add --purge libc6-compat && \
    rm -rf /tmp/* && \
    rm -rf /var/cache/apk/*

ENTRYPOINT ["./s3example"]
