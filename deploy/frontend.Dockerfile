FROM golang:1.23 AS builder

WORKDIR /usr/src/gomall

# Set the GOPROXY environment variable for mainland China users
ENV GOPROXY=https://goproxy.io,direct

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen
COPY common common

RUN cd common && go mod download && go mod verify
RUN cd app/frontend/ && go mod download && go mod verify

COPY app/frontend app/frontend

RUN cd app/frontend && go get github.com/hertz-contrib/obs-opentelemetry/provider@v0.3.0 && go get go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc@v0.42.0

RUN cd app/frontend/ && go build -v -o /opt/gomall/frontend/server

# only executable and configuration files

FROM busybox

COPY --from=builder /opt/gomall/frontend/server /opt/gomall/frontend/server

COPY app/frontend/.env /opt/gomall/frontend/.env
COPY app/frontend/conf /opt/gomall/frontend/conf
COPY app/frontend/static /opt/gomall/frontend/static
COPY app/frontend/template /opt/gomall/frontend/template

WORKDIR /opt/gomall/frontend

EXPOSE 8080

CMD ["./server"]
