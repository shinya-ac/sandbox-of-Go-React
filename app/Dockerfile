FROM golang:1.19
ENV TZ=Asia/Tokyo

WORKDIR /app/src

COPY ./src /app/src

RUN go mod tidy

CMD ["go","run","main.go"]