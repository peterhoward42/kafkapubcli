# Build:
#   cd <project root>
#   docker build -f build/Dockerfile -t kafkapubcli:1.0 .  

FROM golang:alpine as builder
RUN mkdir /project
ADD . /project/
WORKDIR /project/cmd
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /project/cmd/main /app/
WORKDIR /app
CMD ["./main"]