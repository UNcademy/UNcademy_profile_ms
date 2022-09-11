FROM golang:alpine

RUN mkdir /UNcademy_profile_ms
RUN go install github.com/beego/bee/v2@latest

WORKDIR /UNcademy_profile_ms

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

EXPOSE 8001

CMD ["bee", "run"]