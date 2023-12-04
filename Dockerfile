#First Stage
#Finds image from AWS public repo.
FROM public.ecr.aws/docker/library/golang:latest AS build
WORKDIR /build
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .


#Second stage
#From AWS public repo.
FROM public.ecr.aws/docker/library/alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add ca-certificates && \
    apk add tzdata

WORKDIR /app

COPY --from=build /build/main ./

RUN pwd && find .


#Runs binary with 
CMD ["./main",  "-i", "120"]
