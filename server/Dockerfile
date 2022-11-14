FROM alpine:latest

RUN mkdir "/app"
RUN mkdir "/app/config"

WORKDIR "/app"
COPY main /app/
COPY config.ini /app/config/

ENTRYPOINT ["./main"]