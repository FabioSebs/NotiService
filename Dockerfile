# Build Stage
FROM golang:1.20 as build

WORKDIR /app

# Dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notifications

# Execute Stage
FROM alpine:3.13

WORKDIR /app

# Copy only necessary files from the build stage
COPY --from=build /app/notifications .
COPY --from=build /app/.env .
COPY --from=build /app/stringer.opml .

EXPOSE 6000
ENTRYPOINT [ "/app/notifications" ]
