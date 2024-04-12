FROM golang:1.21 AS build
ENV GO111MODULE=on
RUN apt-get update && apt-get install -y build-essential
WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN make build

FROM alpine:3.17.3
RUN apk update && apk add --no-cache libc6-compat
RUN mkdir -p /home/jake/templates
RUN mkdir -p /home/jake/result
WORKDIR /home/jake
COPY --from=build /usr/src/app/templates/ /home/jake/templates/
COPY --from=build /usr/src/app/bin/upload-app .
CMD ["./upload-app"]
EXPOSE 8082