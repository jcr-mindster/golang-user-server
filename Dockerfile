FROM golang:rc-buster

LABEL version="0.1.0"

COPY . /target

WORKDIR /target

EXPOSE 3000 3000

ENTRYPOINT ["go", "run", "main.go"]