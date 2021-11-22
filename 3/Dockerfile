
FROM golang:1.13 as base

WORKDIR  /var/opt
COPY src/main.go .
RUN go build main.go

FROM ubuntu
WORKDIR /var/opt
COPY --from=base /var/opt/main .
RUN  chmod +x main
EXPOSE 8000


ENTRYPOINT [ "./main" ]

