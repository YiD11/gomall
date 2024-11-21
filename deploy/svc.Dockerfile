FROM golang:1.23 AS builder

ARG SVC

WORKDIR /usr/src/gomall

# Set the GOPROXY environment variable for mainland China users
ENV GOPROXY=https://goproxy.io,direct

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY app/${SVC}/go.mod app/${SVC}/go.sum ./app/${SVC}/
COPY rpc_gen rpc_gen
COPY common common

RUN cd common && go mod download && go mod verify
RUN cd app/${SVC}/ && go mod download && go mod verify

COPY app/${SVC} app/${SVC}

RUN cd app/${SVC} && go get github.com/hertz-contrib/obs-opentelemetry/provider@v0.3.0 && go get go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc@v0.42.0

RUN cd app/${SVC}/ && go build -v -o /opt/gomall/${SVC}/server

# only executable and configuration files

FROM busybox

ARG SVC

COPY --from=builder /opt/gomall/${SVC}/server /opt/gomall/${SVC}/server

COPY app/${SVC}/conf /opt/gomall/${SVC}/conf

WORKDIR /opt/gomall/${SVC}

CMD ["./server"]
