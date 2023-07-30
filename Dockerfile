FROM alpine:latest

COPY bin/ports ports
COPY json_files json_files

EXPOSE 8080
CMD ["./ports", "run"]