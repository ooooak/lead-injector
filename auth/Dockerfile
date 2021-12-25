FROM golang:alpine


RUN mkdir /app
WORKDIR /app
COPY ./ /app
RUN go mod download
RUN go build -o ./auth .

EXPOSE 3000

ENTRYPOINT ./auth