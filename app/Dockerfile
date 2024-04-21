FROM golang:1.21 as golang
WORKDIR /src
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app .
FROM gcr.io/distroless/static-debian11
COPY --from=golang /app .
COPY --from=golang /bin/bash .
COPY --from=golang /bin/curl .
COPY --from=golang /bin/sleep .
EXPOSE 6443
EXPOSE 8200
CMD ["sleep","100"]
CMD ["/app"]
