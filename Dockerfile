FROM chainguard/go:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target="/root/.cache/go-build" go mod download

COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o /app

# Deploy the application binary into a lean image
FROM chainguard/wolfi-base:latest AS build-release-stage

WORKDIR /

COPY --from=build /app /app

ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app"]