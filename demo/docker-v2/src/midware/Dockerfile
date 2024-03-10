FROM golang:latest as build
RUN mkdir -p /workspace
ADD main.go typed.go utils.go /workspace/
WORKDIR /workspace
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./main *.go

FROM alpine:latest
RUN mkdir -p /workspace
COPY --from=build /workspace/main /workspace
ENTRYPOINT ["/workspace/main"]
