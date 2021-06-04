FROM golang:1.16.3
COPY . /app
RUN cd /app && make
WORKDIR /app/output
EXPOSE 8000
ENTRYPOINT ./bin/goldtalkAPI