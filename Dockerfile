FROM alpine:latest

COPY gate .
COPY conf conf
VOLUME [ "/certs" ]
CMD [ "./gate" ]
