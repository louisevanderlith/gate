FROM golang:1.12 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY domains ./domains
COPY routers ./routers 

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/gate .
COPY conf conf

EXPOSE 80
EXPOSE 443

ENTRYPOINT [ "./gate" ]