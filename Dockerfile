# Building
FROM golang:1.15-alpine AS builder

# define working directory
WORKDIR /opt/app

# copy source from current dir to working dir
COPY . .

# build
RUN go build -o main .

# Running
FROM alpine:latest AS runner

WORKDIR /opt/

COPY --from=builder /opt/app/main .
COPY ./configs/ ./configs/

# inform exposed ports
EXPOSE 3000 3001

CMD ["./main"]
