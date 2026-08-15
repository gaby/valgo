---
title: Getting Started with Go Validation
description: Install Valgo, a type-safe Go validation library without struct tags, and learn Is(), Check(), and New() with examples.
---

## Introduction

Valgo is a type-safe, expressive, and extensible validation library for Go with built-in i18n support.

Unlike validation libraries that rely on struct tags, Valgo defines validation rules as functions. This gives you greater flexibility to validate any value, compose rules programmatically, and decide where validation belongs within your application.

Valgo can be customized to fit your application's needs, from overriding validation messages to localizing them for different languages and contexts.

## Install

```bash
go get github.com/cohesivestack/valgo
```

Valgo v0.9.0 has been tested with Go 1.23 and later.

## Agent skill

This repository includes a Valgo Agent Skill installable with [`npx skills`](https://github.com/vercel-labs/skills):

```bash
npx skills add cohesivestack/valgo --skill valgo
```

## Your first validation

`Is(...)` creates a validation session. Within each validator chain, rules are
evaluated from left to right and stop after the first failed rule. OR groups and
`OrElse()` have their own control-flow behavior; see `Validators -> OR
Operators`.

```go
import (
  "encoding/json"
  "fmt"
  v "github.com/cohesivestack/valgo"
)

val := v.Is(
  v.String("Bob", "full_name").Not().Blank().LengthBetween(4, 20),
  v.Number(17, "age").GreaterThan(18),
)

if err := val.ToError(); err != nil {
  out, _ := json.MarshalIndent(err, "", "  ")
  fmt.Println(string(out))
}
```

## Extended example

This larger example composes validation without struct tags. It shows
alternative rules, optional values, validation that depends on earlier
results, and nested error paths in one flow.

```go
type Address struct {
  Line1      string
  City       string
  Country    string
  PostalCode string
}

type Registration struct {
  Name                 string
  Email                string
  Password             string
  PasswordConfirmation string
  ContactPreference    string
  ReferralCode         string
  Address              Address
}

func validateRegistration(input Registration) *v.Validation {
  return v.Is(
    v.String(input.Name, "name").Not().Blank().LengthBetween(2, 80),
    v.String(input.Email, "email").Not().Blank(),
    v.String(input.Password, "password").Not().Blank().LengthBetween(10, 64),
    v.String(input.ContactPreference, "contact_preference").
      EqualTo("email").Or().EqualTo("sms"),
    v.String(input.ReferralCode, "referral_code").
      Empty().OrElse().MatchingTo(regexp.MustCompile(`^[A-Z0-9]{6,12}$`)),
  ).IfPathValid(
    "password",
    v.Is(v.String(input.PasswordConfirmation, "password_confirmation").
      EqualTo(input.Password)),
  ).In("address",
    v.Is(
      v.String(input.Address.Line1, "line1").Not().Blank(),
      v.String(input.Address.City, "city").Not().Blank(),
      v.String(input.Address.Country, "country").Not().Blank().EqualTo("US"),
      v.String(input.Address.PostalCode, "postal_code").Not().Blank(),
    ).WhenAllValid([]string{"country", "postal_code"}, func(val *v.Validation) {
      if err := verifyPostalCode("US", input.Address.PostalCode); err != nil {
        val.AddErrorMessage("postal_code", "Postal code could not be verified")
      }
    }),
  )
}
```

The example assumes that the application provides `verifyPostalCode`.

- `Or()` accepts `email` or `sms` as the contact preference.
- `OrElse()` accepts an empty referral code without evaluating the regex.
- `IfPathValid()` merges the password-confirmation validation only after the
  password passes.
- `In("address", ...)` prefixes nested errors with `address`.
- `WhenAllValid()` calls postal-code verification only after the initial
  country and postal-code rules pass.

See [OR operators](/validators/or-operators/),
[conditional validation](/using-valgo/conditional-flows/), and
[namespaces](/using-valgo/namespaces/) for the detailed behavior.

## When to use `Is` vs `Check`

- `Is(...)`: stops a validator chain after its first failed rule.
- `Check(...)`: continues evaluating rules after failures so it can collect
  multiple messages. A successful `OrElse()` still cuts the remainder of its
  chain by design.

```go
val := v.Check(
  v.String("", "full_name").Not().Blank().LengthBetween(4, 20),
)

_ = val.Valid() // false, with 2 messages for full_name
```

## Nested models and collections

Use namespaces to build structured paths:

- `In("ns", ...)` for nested structs
- `InRow("list", i, ...)` for slices of structs
- `InCell("list", i, ...)` for slices of scalar values

See [Namespaces](/using-valgo/namespaces/) and
[Slices and indexed errors](/cookbook/slices/) for complete examples.
