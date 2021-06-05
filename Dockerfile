FROM golang:1.15.13-alpine3.13
RUN apk add make
COPY . /app
RUN cd /app && make
WORKDIR /app/output
EXPOSE 8000
ENTRYPOINT ./bin/goldtalkAPI