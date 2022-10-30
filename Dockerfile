FROM alpine:latest

WORKDIR /home
RUN apk add curl
RUN curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
RUN ln -s /root/.tiup/bin/tiup /bin/tiup
RUN tiup --version

ENV TIDB_VERSION=v6.3.0
RUN tiup install tidb:${TIDB_VERSION} pd:${TIDB_VERSION} tikv:${TIDB_VERSION} playground
