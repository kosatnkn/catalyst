# Building
FROM golang:1.13-alpine AS build

# define working directory
WORKDIR /opt/app

# copy source from current dir to working dir
COPY . .
RUN go build -o main .

# NOTE: usage of CGO_ENABLED and installSuffix is no longer needed since go 1.10
#       like in the following command
#       https://github.com/golang/go/issues/9344#issuecomment-69944514
# RUN GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

# Running
FROM alpine:latest AS runner 
# RUN apk --no-cache add ca-certificates
WORKDIR /opt/
COPY --from=build /opt/app/main .
COPY ./config/ ./config/

# inform exposed ports 
EXPOSE 3000 3001

CMD ["./main"]
