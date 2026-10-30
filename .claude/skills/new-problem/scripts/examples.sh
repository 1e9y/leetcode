#!/usr/bin/env bash
# Prints the "Example N" blocks (Input / Output / Explanation) of a LeetCode problem.
# Usage: examples.sh <title-slug>     e.g. examples.sh valid-anagram
set -euo pipefail

slug="${1:?usage: examples.sh <title-slug>}"

curl -sf https://leetcode.com/graphql \
  -H 'Content-Type: application/json' \
  -H "Referer: https://leetcode.com/problems/$slug" \
  -d "{\"query\":\"query q(\$s: String!){question(titleSlug:\$s){content}}\",\"variables\":{\"s\":\"$slug\"}}" |
  jq -r --arg s "$slug" '.data.question.content // error("no such problem: \($s)")' |
  sed -e 's/<[^>]*>//g' -e 's/&quot;/"/g' -e 's/&nbsp;/ /g' -e 's/&lt;/</g' -e 's/&gt;/>/g' -e 's/&#39;/'"'"'/g' -e 's/&amp;/\&/g' |
  awk '/^Example [0-9]+/ { on = 1 } /^Constraints:/ { exit } on && NF'
