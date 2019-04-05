FROM golang:1.11 as builder

WORKDIR /box
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY domains ./domains

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/gate .
COPY conf conf

EXPOSE 80
EXPOSE 443

ENTRYPOINT [ "./gate" ]