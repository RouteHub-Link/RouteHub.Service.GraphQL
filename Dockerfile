# Build.
FROM golang:alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian12 AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/config/authorization/rbac_model.conf /config/authorization/rbac.model.conf
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]