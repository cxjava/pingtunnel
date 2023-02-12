FROM alpine:latest
LABEL maintainer="cxjava <cxjava@github.com>"

RUN set -ex \
    && apk add --no-cache tzdata openssl ca-certificates \
    && wget -O - https://raw.githubusercontent.com/cxjava/pingtunnel/main/install.sh | sh -s -- -b /usr/local/bin

ENTRYPOINT ["/usr/local/bin/pingtunnel"]
