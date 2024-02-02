# syntax=docker/dockerfile:experimental

FROM golang:1.21 AS build-env
ENV GO111MODULE on
WORKDIR /
RUN git clone https://github.com/coredns/coredns
COPY . /k8s_dns_chaos
# RUN ln -s /k8s_dns_chaos /coredns/plugin/k8s_dns_chaos
RUN sed -i '/kubernetes/a\k8s_dns_chaos:github.com/chaos-mesh/k8s_dns_chaos' /coredns/plugin.cfg
RUN cd coredns && \
    go mod edit -require github.com/chaos-mesh/k8s_dns_chaos@v0.2.6 && \
    go mod edit -replace github.com/chaos-mesh/k8s_dns_chaos=/k8s_dns_chaos && \ 
    go get github.com/chaos-mesh/k8s_dns_chaos && \
    go generate && \
    go mod tidy
RUN cd coredns && make 

FROM debian:stable-slim AS certs
RUN apt-get update && apt-get -uy upgrade
RUN apt-get -y install ca-certificates && update-ca-certificates

FROM scratch
LABEL org.opencontainers.image.source=https://github.com/chaos-mesh/k8s_dns_chaos
COPY --from=certs /etc/ssl/certs /etc/ssl/certs
COPY --from=build-env /coredns/coredns /coredns
EXPOSE 53 53/udp
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn
ENTRYPOINT ["/coredns"]
