FROM golang:rc-buster

COPY . /target

WORKDIR /target

EXPOSE 3000 3000

ENTRYPOINT ["go", "run", "main.go"]