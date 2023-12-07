FROM golang:1.20
ENV GO111MODULE=on
RUN apt-get update && apt-get install -y build-essential
WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN go build -o main .
CMD ["./main", "-mode", "graceful"]
EXPOSE 8081