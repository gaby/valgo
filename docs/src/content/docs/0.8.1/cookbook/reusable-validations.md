---
title: Reusable Go Validation Functions
description: Extract Valgo validations into reusable Go functions and compose
  validation sessions with Merge().
slug: 0.8.1/cookbook/reusable-validations
---

```go
validatePreStatus := func(status string) *v.Validation {
  regex, _ := regexp.Compile("pre-.+")
  return v.Is(v.String(status, "status").Not().Blank().MatchingTo(regex))
}

val := v.Is(
  v.String(r.Name, "name").Not().Blank(),
  v.String(r.Status, "status").Not().Blank(),
)

val.Merge(validatePreStatus(r.Status))
```
