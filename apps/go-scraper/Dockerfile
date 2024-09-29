FROM golang:1.23 AS BUILD
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/out/app

FROM alpine:3.19 AS FINAL
WORKDIR /app
COPY --from=BUILD /app/out/app /app/out/app

ENTRYPOINT ls
# ENTRYPOINT /app/out/app