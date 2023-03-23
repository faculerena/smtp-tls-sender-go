FROM golang:1.20

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o smtp-tls-api .

EXPOSE 8080

#ENV SERVER=smtp.example.com \
#    USER=your-email@example.com \
#    PASSWORD=your-email-password

CMD ["./smtp-tls-api"]