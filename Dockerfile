FROM golang:latest as build
RUN mkdir -p /go/src/proofviewer
WORKDIR /go/src/proofviewer
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go get -d -v ./...
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o proofviewer

FROM ubuntu:latest as hc
ADD https://github.com/chrisaxiom/docker-health-check/archive/v0.3.tar.gz /
RUN tar -xvzf /v0.3.tar.gz

FROM scratch as final

WORKDIR /proofviewer

COPY --from=build /go/src/proofviewer/proofviewer .
COPY --from=hc /docker-health-check-0.3/docker-health-check /docker-health-check
HEALTHCHECK --interval=8s --timeout=120s --retries=8 CMD ["/docker-health-check", "-url=http://127.0.0.1:8090/health"]

EXPOSE 8090

CMD ["./proofviewer"]