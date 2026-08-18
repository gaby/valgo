---
title: Reusable Go Validation Functions
description: Extract Valgo validations into reusable Go functions and compose
  validation sessions with Merge().
slug: 0.9.0/cookbook/reusable-validations
---

## Reuse a validation session

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

## Reuse a stateless predicate

When you only need a boolean result, import
`github.com/cohesivestack/valgo/is`. Predicates contain the same rule logic as
Valgo's built-in validators, without field names, error messages, localization,
or validation sessions.

```go
import "github.com/cohesivestack/valgo/is"

func validUsername(username string) bool {
  return is.StringLengthBetween(username, 3, 30) &&
    !is.StringBlank(username)
}
```

The same function can be reused inside a validator when an error message is
needed:

```go
val := v.Is(
  v.String(input.Username, "username").
    Passing(validUsername, "{{title}} must be 3-30 non-blank characters"),
)
```

See [Stateless Predicates](/0.9.0/validators/predicates/) for the available families
and pointer behavior.
