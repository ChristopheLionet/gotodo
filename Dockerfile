FROM golang:1.24-bullseye
WORKDIR /app
RUN mkdir "publish" 
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls -al
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux go build -o ./publish -v ./...
CMD ["./publish/gotodosrv"]
HEALTHCHECK --interval=10s --timeout=3s CMD curl -f http://localhost:8080/diagnostics/healthcheck || exit 1