# Valgo: Type-Safe Validation for Go

Valgo is a type-safe, expressive, and extensible validation library for Go with built-in i18n support.

Unlike validation libraries that rely on struct tags, Valgo defines validation rules as functions. This gives you greater flexibility to validate any value, compose rules programmatically, and decide where validation belongs within your application.

Valgo can be customized to fit your application's needs, from overriding validation messages to localizing them for different languages and contexts.

**Valgo is pre-v1.0, so breaking changes can happen.**

## Quick Minimal Example

```go
package main

import (
  "encoding/json"
  "fmt"

  v "github.com/cohesivestack/valgo"
)

func main() {
  val := v.Is(
    v.String("Bob", "full_name").Not().Blank().LengthBetween(4, 20),
    v.Number(17, "age").GreaterThan(18),
  )

  if !val.Valid() {
    out, _ := json.MarshalIndent(val.ToError(), "", "  ")
    fmt.Println(string(out))
  }
}
```

Output:

```json
{
  "age": [
    "Age must be greater than \"18\""
  ],
  "full_name": [
    "Full name must have a length between \"4\" and \"20\""
  ]
}
```

## Extended Example

This account registration example shows how validation rules can be composed
without struct tags, including nested structs and validation that depends on
earlier results.

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

* Here, `Or()` accepts either supported contact preference, while `OrElse()`
short-circuits the referral-code pattern check when the optional code is empty.
* `IfPathValid()` merges the password-confirmation result only when the password
itself passes.
* The address rules are grouped with `In()`, so their errors use
paths such as `address.city` and `address.postal_code`. Within that nested
validation, `WhenAllValid()` calls the application's `verifyPostalCode`
function only after both the country and postal code pass their initial rules.
* The excerpt assumes that `verifyPostalCode` is provided by the application.

## Website and documentation

[valgo.build](https://valgo.build)

## Installing

```bash
go get github.com/cohesivestack/valgo
```

## Agent skill

This repository includes a Valgo Agent Skill installable with [`npx skills`](https://github.com/vercel-labs/skills):

```bash
npx skills add cohesivestack/valgo --skill valgo
```

## Docs

- Start Here
  - [Getting Started](https://valgo.build/getting-started/)
  - [Migration Notes](https://valgo.build/migration/)
- Using Valgo
  - [Validation Sessions](https://valgo.build/using-valgo/validation-sessions/)
  - [Namespaces](https://valgo.build/using-valgo/namespaces/)
  - [Querying Results](https://valgo.build/using-valgo/querying-results/)
  - [Conditional Flows](https://valgo.build/using-valgo/conditional-flows/)
  - [Errors & Output](https://valgo.build/using-valgo/errors/)
  - [Localization & Factory](https://valgo.build/using-valgo/localization/)
- Validators
  - [Overview](https://valgo.build/validators/overview/)
  - [String](https://valgo.build/validators/string/)
  - [Numbers](https://valgo.build/validators/numbers/)
  - [Boolean](https://valgo.build/validators/boolean/)
  - [Time](https://valgo.build/validators/time/)
  - [Comparable](https://valgo.build/validators/comparable/)
  - [Typed & Any](https://valgo.build/validators/typed-any/)
  - [OR Operators (Or / OrElse)](https://valgo.build/validators/or-operators/)
  - [Rule Index](https://valgo.build/validators/rule-index/)
  - [Stateless Predicates](https://valgo.build/validators/predicates/)
- Extending
  - [Custom Validators](https://valgo.build/extending/custom-validators/)
- Cookbook
  - [Overview](https://valgo.build/cookbook/)
  - [Sign-up Form](https://valgo.build/cookbook/signup-form/)
  - [Nested Structs](https://valgo.build/cookbook/nested-structs/)
  - [Slices & Indexed Errors](https://valgo.build/cookbook/slices/)
  - [Optional Fields (Pointers)](https://valgo.build/cookbook/optional-fields/)
  - [Conditional Rules](https://valgo.build/cookbook/conditional-rules/)
  - [Custom Messages](https://valgo.build/cookbook/custom-messages/)
  - [Localization](https://valgo.build/cookbook/localization/)
  - [Reusable Validations](https://valgo.build/cookbook/reusable-validations/)
- About
  - [License](https://valgo.build/about/license/)

# Github Code Contribution Guide

We welcome contributions to our project! To make the process smooth and efficient, please follow these guidelines when submitting code:

* **Discuss changes with the community**: We encourage contributors to discuss their proposed changes or improvements with the [community](https://github.com/cohesivestack/valgo/discussions/categories/ideas) before starting to code. This ensures that the changes align with the focus and purpose of the project, and that other contributors are aware of the work being done.

* **Make commits small and cohesive**: It is important to keep your commits focused on a single task or change. This makes it easier to review and understand your changes.

* **Check code formatting with go fmt**: Before submitting your code, please ensure that it is properly formatted using the go fmt command.

* **Make tests to cover your changes**: Please include tests that cover the changes you have made. This ensures that your code is functional and reduces the likelihood of bugs.

* **Update golang docs and README to cover your changes**: If you have made changes that affect documentation or the README file, please update them accordingly.

* **Keep a respectful language with a collaborative tune**: We value a positive and collaborative community. Please use respectful language when communicating with other contributors or maintainers.

* **Go version support:**: Valgo supports the Go versions currently supported by the Go project, plus two previous Go major versions when compatible. The minimum supported Go version is declared in `go.mod`.

# License

Copyright © 2026 Carlos Forero

Valgo is developed and maintained by [Cohesive Stack LLC](https://cohesivestack.com) and released under the [MIT License](LICENSE).
