FROM golang:1.17

RUN go version
ENV GOPATH=/
COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x scripts/wait-for-postgres.sh

RUN go mod download
RUN go build -o houser ./main.go

CMD ["./houser"]
