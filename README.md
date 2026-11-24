# Delete a Given Node in O(1)

> Striver SDE Sheet 030 · Day 5 Linked List — you can't unlink a node you can't reach the predecessor of, so become your successor and unlink it instead. Go, stdlib only, tested.

## Problem

Given *only* a pointer to a node inside a singly linked list — no head, no predecessor — delete that node in O(1). (Guaranteed not to be the tail in the classic statement; this version returns false for a tail instead of pretending.)

## Approach

Unlinking requires the predecessor, which is unreachable. The trick: don't unlink *this* node — overwrite its value with the successor's value, then splice the successor out (its predecessor is the node in hand). Observationally the list now looks as if the target was deleted. The honest caveats, which interviewers fish for: it cannot work on the tail, and it breaks external pointers that identified the successor node — which is why it returns false rather than guessing on tails.

**Complexity:** O(1) time and space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
