# Stage 1: BUILD
FROM golang:1.19.1-bullseye as BUILDER

RUN apt-get update && apt-get -y install --no-install-recommends \
    ca-certificates \
    bash \
    tzdata

COPY / /build
WORKDIR /build
ENV GO111MODULE=on 
RUN go get || true

RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN GODEBUG="madvdontneed=1" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app


# Stage 2: DEPLOY
FROM ubuntu:22.04
ARG TZ
ENV TZ=${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && apt-get -y install --no-install-recommends \
    ca-certificates \
    tzdata \
    bash \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean

COPY --from=BUILDER /app /app
COPY /lib/server/ca/binaries/mkcert-linux /bin/mkcert
RUN chmod a+x /bin/mkcert

VOLUME ["/tmp"]
EXPOSE 9000

ENV GODEBUG="madvdontneed=1"
ENTRYPOINT ["/app"]
