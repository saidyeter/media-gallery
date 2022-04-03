FROM  golang:1.17-alpine
RUN apk add  --no-cache ffmpeg
RUN mkdir /app
ADD . /app
WORKDIR /app/server
RUN go mod download
RUN go build -o main ./cmd/media-gallery
CMD ["/app/server/main"]