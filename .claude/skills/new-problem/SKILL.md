---
name: new-problem
description: Scaffold a new LeetCode problem in this repo — run the kit generator, then fill the test table with the official LeetCode examples, leaving the solution empty for the owner to write. Use whenever the user names a problem to start, e.g. "next problem: 347", "start solving 217", "let's do 49", "new problem two sum", "/new-problem 242", even if they don't say "scaffold". Do not use for reviewing, hinting, or committing an existing solution.
argument-hint: <problem-id>
---

# New problem scaffold

Problem: `$ARGUMENTS` (if empty, take the ID from the user's message; if they gave a title, use the matching ID).

The owner solves every problem personally; this repo is practice. Your job ends at a scaffold whose tests compile against an empty function. Don't write or hint at solution logic, don't suggest an approach, and don't mention algorithms in your report. Hints happen only when the user asks for them.

## 1. Generate

From the repo root:

```sh
ls -d "$(printf '%04d' <id>)".* 2>/dev/null   # already scaffolded? then stop and tell the user
go run . gen -p <id>
```

The generator writes `NNNN.title_slug/` (zero-padded ID; slug with `-` → `_`) with `<slug>.go`, `<slug>_test.go`, `README.md`.

If it panics with a nil pointer, the problem is missing from the cached `problems.json` (usually a newly released problem), or LeetCode has no Go snippet for it. Delete `problems.json` (a gitignored cache, re-downloaded on the next run) and retry once; if it still fails, report that to the user.

Leave `<slug>.go` exactly as generated, including its empty body and 4-space indent. It's the owner's file.

## 2. Fetch the official examples

Don't type examples from memory: LeetCode edits problem statements over time, so fetch the current ones.

```sh
.claude/skills/new-problem/scripts/examples.sh <title-slug>   # dashes, e.g. valid-anagram
```

Use every example, verbatim, and nothing else; no extra edge cases unless the user asks.

## 3. Rewrite the test file

The generated `<slug>_test.go` is a placeholder that doesn't compile (unused `fmt`, `input int`/`expect int`, a call to `sol`). Replace it, matching the function signature in `<slug>.go`:

```go
package valid_anagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	s      string
	t      string
	expect bool
}{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
}

func Test(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, tt.expect, isAnagram(tt.s, tt.t))
	}
}
```

- **Fields:** one per parameter named like the parameter (`nums`, `k`), or a single `input` if there's one parameter, plus `expect`. One example per line.
- **Order-insensitive answers** ("return the answer in any order"): use `assert.ElementsMatch`. For nested results like `[][]string` where inner order is free too, add a small helper that sorts each inner slice of both sides first. Use `slices.Clone` + `slices.Sort`, since `go.mod` targets Go 1.21 (no `slices.Sorted`).
- **Linked lists:** keep the fields as `[]int` and convert with `structures.FromInts`; compare with `e.Equal(got)`. The solution file then needs `type ListNode = structures.ListNode` plus the `github.com/1e9y/leetcode/structures` import. Adding that alias is scaffolding, not solution logic, and it's the one edit to `<slug>.go` that's allowed.
- **Trees or other types not in `structures/`:** stop and ask the user before adding a shared type.

## 4. Verify

```sh
gofmt -l <dir>/<slug>_test.go   # must print nothing
go vet ./<dir>/                  # the only error should be "missing return" in the empty function
```

Any other error means the test file is wrong; fix it. Don't commit: the commit happens after the owner solves the problem (`solution(<id>): <title>`, see CLAUDE.md).

## 5. Report

Keep it short: the directory created, the examples as they appear in the test table, the assertion style if it isn't `assert.Equal` (and why), and that the function is empty and the only build error is the expected "missing return". Then hand over to the owner.
