# Building
FROM golang:1.24-alpine AS builder

# define working directory
WORKDIR /opt/app

# copy source from current dir to working dir
COPY . .

# update metadata
RUN chmod +x ./metadata/set_metadata.sh \
  && sh ./metadata/set_metadata.sh $(pwd)

# build
RUN go build -o main .

# Running
FROM alpine:3.22.2 AS runner

WORKDIR /opt/

COPY --from=builder /opt/app/main .

# inform exposed ports
EXPOSE 8000

CMD ["./main"]
