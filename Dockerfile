FROM golang:1.11.11

COPY main.go /app/

CMD ["go", "run", "/app/main.go"]