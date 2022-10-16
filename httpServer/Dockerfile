FROM ubuntu

RUN apt-get update && \
    apt-get install golang-go

COPY . .


RUN go build -o HttpServer
RUN chmod 777 HttpServer


ENTRYPOINT ["./HttpServer"]