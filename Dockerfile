FROM golang:1.18

WORKDIR /app

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["air"]

