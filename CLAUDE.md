# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

Personal collection of LeetCode solutions in Go, plus a small CLI ("kit") that scaffolds new problems.

## IMPORTANT: Do not solve the problems

The repo owner solves every challenge personally; this repo is practice. AI assists with everything around the solution, never the solution itself.

**Allowed:**
- Scaffolding a problem (`gen`) and filling the test table with LeetCode's examples.
- Running tests, `go vet`, `gofmt`, and committing/pushing when asked.
- Reviewing the owner's code when asked: point out bugs, failing cases, complexity, and Go idioms, explaining the *why*, and leave the fix to the owner.
- Answering general Go or language questions (e.g. "how to convert a string to a slice").
- Hints and clues, **only when explicitly asked**, and only as many as asked for ("give me one hint" means one).

**Never:**
- Write, complete, or rewrite solution logic, even partially or "just to show".
- Suggest an algorithm, approach, or solution unprompted, including in reviews or at scaffolding time.
- Edit a solution file beyond what the owner explicitly asks for (e.g. a requested rename).

## Commands

```sh
go run . gen -p <id>          # scaffold a problem (must run from repo root)
go test ./...                 # all solutions (what CI runs on push to main)
go test ./0217.contains_duplicate/
go test -run Test -v ./0049.group_anagrams/
go vet ./<dir>/ && gofmt -l <dir>
```

## Layout

- Each problem is its own package in `NNNN.title_slug/` (zero-padded frontend ID, slug with `_`), containing `<slug>.go`, `<slug>_test.go`, and a one-line `README.md` linking the problem.
- Tests are table-driven: a package-level `var tests = []struct{...}` and a single `func Test`, using `testify/assert`.
- `structures/` holds shared types (e.g. `ListNode` with `FromInts`/`Equal`) for linked-list problems.

## Scaffolding kit (`main.go` → `cmd/` → `kit/`)

- `cmd/gen.go` wires the `gen` cobra command to `kit.Init`, `kit.Get`, `kit.Generate`.
- `kit.Init` loads `./problems.json` (gitignored cache of LeetCode's problem list; downloaded on first run). Delete it to pick up newly released problems.
- `kit.Get` fetches the Go code snippet via LeetCode's GraphQL API; `kit.Generate` renders `kit/source.tmpl`, `kit/test.tmpl`, `kit/readme.tmpl` into the new directory.
- Paths (`./kit/*.tmpl`, `./problems.json`) are relative, so the CLI only works from the repo root. `gen` panics if the problem directory already exists.
- The generated test file is a placeholder that does not compile as-is: it imports an unused `fmt`, uses `input int`/`expect int` fields, and calls `sol(...)`. It must be rewritten for the problem's real signature.

## Workflow

- When starting a problem, use the `new-problem` skill (`.claude/skills/new-problem/`): it runs `gen`, fills the test table with the official LeetCode examples, and leaves the solution function empty (see "Do not solve the problems" above).
- When the answer may be returned in any order, compare with `assert.ElementsMatch` (and normalize inner order for nested results).

## Commit messages

Conventional Commits with a custom `solution` type. Format: `<type>[(<scope>)]: <subject>`. Subject is a single lowercase line, imperative mood, no trailing period, with code identifiers in backticks. No body and no trailers; never add `Co-Authored-By` or any other AI attribution.

| Type | Use for | Example from history |
|---|---|---|
| `solution(<id>)` | A new or reworked solution for one problem | `solution(49): group anagrams` |
| `feat` | New tooling or repo capability | `feat: add script for generating commit messages` |
| `fix` | Bug fixes in tooling or templates | `fix(kit): call solution with fields defined in test template` |
| `docs` | README and other documentation | `docs: add workflow badge to README file` |
| `ci` | GitHub Actions workflows | `ci: run go test on all nested packages` |
| `chore` | Dependencies, Go version, `.gitignore`, cleanup | `chore: bump go version to 1.21` |

- **`solution` commits:** the scope is the problem's frontend ID without leading zeros (`0049.group_anagrams` → `49`), and the subject is the directory slug with `_` replaced by spaces (`n-th tribonacci number`, `jump game ii`). It's exactly what `scripts/git_commit_solutions.fish` generates. One problem per commit, containing only that problem's directory. Revisiting a solved problem uses the same form (`solution(491): ...` appears twice). The early `solution(1): solution for problem 1. two sum` format is obsolete.
- **Scope for other types:** `kit` when the change is limited to the scaffolding CLI (`feat(kit)`, `fix(kit)`); otherwise no scope.
- **Legacy:** older workflow changes used `infra:`; use `ci:` for new ones.
- One logical change per commit. Unrelated work in the tree gets split into separate commits by type.
