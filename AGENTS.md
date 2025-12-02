# Repository Guidelines

## Project Structure & Module Organization
- Root `go.mod` declares module `github.com/yoshiyoshifujii/go-sample` targeting Go 1.25.4; dependencies live in `go.sum`.
- Each `sample_*` folder (`sample_context`, `sample_di`, `sample_enum`, `sample_enum2`, `sample_interface`, `sample_override`) is a self-contained `main` package demonstrating a pattern; run them individually.
- Tests currently reside in `sample_di/main_test.go` using `testify/assert`; add new tests alongside the code they cover.

## Build, Test, and Development Commands
- Format all code before review: `go fmt ./...`
- Run the full test suite: `go test ./...`
- Execute a specific sample: `go run ./sample_di` (replace with target directory); use `go build ./...` if you need binaries for all samples.
- Refresh dependencies if imports change: `go mod tidy`

## Coding Style & Naming Conventions
- Rely on `go fmt` defaults (tabs for indentation, standard import grouping); avoid manual style tweaks.
- Use clear, descriptive names; export types and functions only when they must be reused outside the package.
- Keep functions small and focused; prefer returning errors over panics in non-demo code.
- Comments should explain intent, not mechanics; keep them close to the code they clarify.

## Testing Guidelines
- Use the standard testing package with `testify/assert` for readability; name tests `TestXxx` and table-drive behavior-heavy logic.
- Aim for coverage on new logic paths; when adding examples, use `ExampleXxx` where it improves discoverability.
- For behavior dependent on interfaces (e.g., `Service`), provide lightweight fakes/mocks in the same test file.

## Commit & Pull Request Guidelines
- Write commit subjects in the imperative mood (`Add DI sample tests`), with concise bodies describing motivation when non-obvious.
- Scope commits narrowly (one feature/fix per commit) to keep history clean.
- Pull requests should include: a short summary of the change, any testing performed (`go test ./...` output is enough), and links to related issues or context.
- If the change affects runtime behavior of a sample, note how to run it (e.g., `go run ./sample_context`) and any expected output.

## Communication
- 日本語でのコミュニケーション歓迎です。PR やコメントも日本語で問題ありません。
