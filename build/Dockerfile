# Derived from https://www.cloudreach.com/en/resources/blog/containerize-this-how-to-build-golang-dockerfiles/
# Build with:
#   cd <project root>
#   docker build -f build/Dockerfile -t foo:1.0 .
# Debug build failures with
#   docker run -it <id_last_working_layer> sh

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