FROM alpine:3.12

RUN mkdir "/app"
WORKDIR "/app"

COPY chat /app/app
ENTRYPOINT ["./app"]