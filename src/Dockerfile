FROM golang:1.21.4-bullseye

WORKDIR /usr/src/app

COPY go.mod go.sum ./

#If the go mod files have nothing in them yet then this returns and error. Once the project contains these files then it will work hence the force to return 0
RUN go mod download; exit 0 
RUN go mod verify; exit 0

RUN apt update && apt install snapd -y
RUN snap install core
RUN snap install solc -y

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]