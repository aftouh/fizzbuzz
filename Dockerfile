FROM golang:1.15 as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN make build

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/src/app/out/bin /
CMD ["/fizzbuzz"]
