# Repository Guidelines

## Project Structure & Module Organization
- `Todo_layered/`: Layered Go API (Gin + GORM/sqlite). Folders: `handler/`, `service/`, `repository/`, `model/`.
- `todo-api/`, `service_practice/`, `OSS-Golang/`: Additional Go services/examples.
- `db_practice/`: Go app with MySQL via Docker Compose (`init.sql`, `Dockerfile`).
- `go-openapi-sample/`, `OpenAPI/`: OpenAPI specs and a sample server (`openapi.yaml`, `internal/`).
- `todo-frontend/`: Next.js frontend (`src/`, `public/`).

Code and tests live inside each submodule directory. Work within one module at a time.

## Build, Test, and Development Commands
- Go (run/build/test):
  - `cd <module>` then `go run ./` (run), `go build ./...` (build), `go test ./...` (tests), `go test -cover ./...` (coverage).
- Docker (db_practice):
  - `cd db_practice && docker-compose up -d` (MySQL + app), `docker-compose down` (stop).
- Frontend (todo-frontend):
  - `cd todo-frontend && npm install`
  - `npm run dev` (local dev), `npm run build` (prod build), `npm start` (serve), `npm run lint` (lint).

## Coding Style & Naming Conventions
- Go: use `go fmt ./...` and `go vet ./...` before committing. Packages lowercase, no underscores. Exported names `CamelCase`; unexported `camelCase`. Errors: wrap with `%w` when propagating. Prefer table-driven tests.
- TypeScript/React: components `PascalCase.tsx` under `src/`. Run `npm run lint` in `todo-frontend` and fix warnings.

## Testing Guidelines
- Frameworks: Go standard `testing` only; no frontend test runner configured. Add `_test.go` files next to source; tests named `TestXxx`. Run from module root: `go test ./...`.
- Aim for meaningful coverage of handlers, services, and repositories. Use small seeded DBs (e.g., sqlite in-memory) for repository tests.

## Commit & Pull Request Guidelines
- Commits: short, imperative, and scoped. Example: `todo-api: add PUT /todos` or `Todo_layered: fix repository error wrapping`.
- PRs: include summary, affected module(s) and paths, run instructions, screenshots for UI changes, and linked issues. Ensure `go vet`, `go fmt`, and tests pass locally.

## Security & Configuration Tips
- Do not commit secrets. Use local `.env` files where applicable (e.g., `go-openapi-sample/.env`, `db_practice/.env`).
- Default ports: Go APIs `:8080`, MySQL `3306`, Next.js `3000`. Adjust via env vars and update docs when changed.
