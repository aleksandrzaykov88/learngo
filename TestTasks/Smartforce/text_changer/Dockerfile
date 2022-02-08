FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./ ./

RUN go build -o /text-changer

EXPOSE 8080

ENV ROOT=/home/openvino/workbench/datasets
ENV TOCHANGE=Мяф
ENV REPLACER=Иллидан

CMD [ "/text-changer" ]
