FROM golang:1.22.5

COPY ./ /app/
WORKDIR /app/

ENV PRODUCTION="1"

RUN go get .
RUN go build -o main .

CMD [ "/app/main" ]

EXPOSE 8000