FROM alpine:latest

COPY gate .
COPY conf conf

EXPOSE 80
EXPOSE 443

ENTRYPOINT [ "./gate" ]
