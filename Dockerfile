# Building
FROM golang:1.24-alpine3.23 AS builder

# define working directory
WORKDIR /opt/app

# copy go.mod first to enable caching
COPY ./go.mod ./go.sum ./
RUN go mod download

# copy source from current dir to working dir
COPY . .

# update metadata and build
RUN chmod +x ./metadata/set_metadata.sh \
  && sh ./metadata/set_metadata.sh $(pwd) \
  && go build -v -o main .

# Running
FROM alpine:3.23.2 AS runner

WORKDIR /opt/

COPY --from=builder /opt/app/main .

# inform exposed ports
EXPOSE 8000

CMD ["./main"]
