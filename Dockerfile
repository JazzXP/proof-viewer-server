FROM golang:latest as build
WORKDIR /go/src/proofviewer
COPY . .

RUN go get -d -v ./...
RUN GOOS=linux go build -o proofviewer

FROM scratch as final

WORKDIR /proofviewer

COPY --from=build /go/src/proofviewer/proofviewer /proofviewer/proofviewer

EXPOSE 8090

CMD ["/proofviewer/proofviewer"]