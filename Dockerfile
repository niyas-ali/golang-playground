
FROM golang:1.17-alpine

COPY . /app
ENV HOME /app
WORKDIR /app
RUN go build 
RUN go version
CMD /app/goplayground
