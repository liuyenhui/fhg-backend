# fhg-backend

Go API service that powers the `flux-helm-github` stack. It lives under `cmd/server` and listens on `:8080` to respond to health checks.

## Repository layout

```
fhg-backend/
├── cmd/server/main.go   # entry point for the HTTP server
├── Dockerfile           # built by GitHub Actions to produce the production image
├── go.mod/go.sum        # dependency management
└── .github/workflows    # CI pipeline for tests, multi-arch build, and chart dispatch
```

## Local build & test

```bash
go test ./...
go build ./cmd/server
```

Those commands verify the code compiles and can create the same binary that the pipeline packages into the Docker image.

## CI/CD workflow

- GitHub Actions `docker-publish.yml` (runs on `main` pushes or manual dispatch) performs the following steps:
  1. Runs Go tests and builds the backend binary.
  2. Calculates the next `vMAJOR.MINOR.PATCH` tag (patch increments up to 99 before bumping the next digits).
  3. Uses Docker Buildx to produce `linux/amd64` and `linux/arm64` manifests, pushes both the semantic tag and `latest` to GHCR, and pushes the Git tag.
  4. Once the image is published, triggers `fhg-charts` via a repository dispatch so the Helm chart is automatically repackaged with the new `image.tag`.

Ensure `CHARTS_REPO_TOKEN` (a PAT with repo/workflow permissions for `liuyenhui/fhg-charts`) is configured under this repo's Actions secrets so chart packaging is triggered.
