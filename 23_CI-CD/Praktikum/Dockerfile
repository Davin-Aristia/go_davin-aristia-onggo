FROM golang:1.18-alpine

WORKDIR /app

COPY ./23_CI-CD/Praktikum/go.mod ./
COPY ./23_CI-CD/Praktikum/go.sum ./
RUN go mod download

COPY ./23_CI-CD/Praktikum/ ./

RUN go build -o /cicd

EXPOSE 8080

CMD ["/cicd"]