FROM golang

RUN mkdir /var/app
WORKDIR /var/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

CMD ["./zcelero"]