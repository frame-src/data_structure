FROM golang:latest

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD ["bash"]
