
FROM golang:1.17-alpine

COPY . /app
ENV HOME /app
WORKDIR /app
RUN go build 
# RUN useradd -m heroku
# USER heroku
RUN go version
CMD /app/goplayground
