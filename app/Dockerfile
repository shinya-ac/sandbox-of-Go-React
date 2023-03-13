FROM golang:1.19
ENV TZ=Asia/Tokyo

WORKDIR /app/src

COPY ./src /app/src

RUN go get cloud.google.com/go/vision/apiv1
RUN go get gopkg.in/ini.v1

RUN go mod tidy

CMD ["go","run","main.go"]