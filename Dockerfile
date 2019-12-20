# Building
FROM golang:1.13-alpine AS build

# define working directory
WORKDIR /opt/app

# copy source from current dir to working dir
COPY . .

# build
# NOTE: usage of CGO_ENABLED and installSuffix is no longer needed since go 1.10
#       like in the following command
#       https://github.com/golang/go/issues/9344#issuecomment-69944514
# RUN GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o app .
RUN go build -o main .

# Running
FROM alpine:latest AS runner 

WORKDIR /opt/

COPY --from=build /opt/app/main .
COPY ./config/ ./config/

# inform exposed ports 
EXPOSE 3000 3001

CMD ["./main"]
