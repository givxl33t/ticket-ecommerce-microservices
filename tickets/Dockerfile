FROM golang:1.23 AS common-build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM common-build-stage AS development-build-stage

CMD ["go", "run", "./cmd/web/main.go"]

FROM common-build-stage AS production-build-stage

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fiber-app ./cmd/web

FROM gcr.io/distroless/static:nonroot AS production-runtime-stage

WORKDIR /app

COPY --from=production-build-stage /app/fiber-app .

EXPOSE 3000

ENTRYPOINT ["./fiber-app"]
