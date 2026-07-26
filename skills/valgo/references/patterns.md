# Verified consumer patterns

These patterns were checked against the repository's v0.8.1 public source.
Load only the section needed for the task, adapt names and error handling to the
consumer, and confirm version support before copying an API.

## Basic validation

Valgo v0.8.1 does not provide an `Email()` rule. Use an existing project helper
or a verified regex/predicate instead of inventing one.

```go
package validation

import (
  "regexp"

  v "github.com/cohesivestack/valgo"
)

var emailPattern = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

type Signup struct {
  Email    string
  Password string
  Age      int
}

func ValidateSignup(input Signup) error {
  return v.Is(
    v.String(input.Email, "email", "Email").
      Not().Blank().
      MatchingTo(emailPattern),
    v.String(input.Password, "password", "Password").
      Not().Blank().
      LengthBetween(10, 72),
    v.Int(input.Age, "age", "Age").GreaterOrEqualTo(18),
  ).ToError()
}
```

Use `Check(...)` instead of `Is(...)` only when multiple messages for one
validator chain are part of the required error contract.

## Optional pointer fields

Represent PATCH presence with pointers and skip absent values explicitly. Keep
the pointer constructor even when the value is known to be non-nil so the
validation matches the request type.

```go
type AccountPatch struct {
  DisplayName *string
  Age         *int
}

func ValidateAccountPatch(input AccountPatch) error {
  val := v.New()

  if input.DisplayName != nil {
    val.Is(v.StringP(input.DisplayName, "display_name").
      Not().Blank().
      LengthBetween(2, 80))
  }
  if input.Age != nil {
    val.Is(v.IntP(input.Age, "age").Between(18, 130))
  }

  return val.ToError()
}
```

For v0.8+, `Nil().OrElse().<remaining rules>` can encode “nil or validate the
rest,” but explicit Go control flow is often clearer and also works on v0.7.

## Nested objects and indexed slices

```go
type Address struct {
  City   string
  Street string
}

type LineItem struct {
  SKU      string
  Quantity int
}

type Order struct {
  Shipping Address
  Items    []LineItem
  Tags     []string
}

func ValidateOrder(input Order) error {
  val := v.New().
    In("shipping", v.Is(
      v.String(input.Shipping.City, "city").Not().Blank(),
      v.String(input.Shipping.Street, "street").Not().Blank(),
    ))

  for i, item := range input.Items {
    val.InRow("items", i, v.Is(
      v.String(item.SKU, "sku").Not().Blank(),
      v.Int(item.Quantity, "quantity").Positive(),
    ))
  }
  for i, tag := range input.Tags {
    val.InCell("tags", i, v.Is(
      v.String(tag, "tag").Not().Blank(),
    ))
  }

  return val.ToError()
}
```

This yields paths such as `shipping.city`, `items[0].sku`, and `tags[0]`.
`InCell` places all nested messages on the indexed cell path and ignores the
nested validator name in the resulting path.

## Conditional validation

Use `When` to defer work. `If` receives an already-created validation because
Go evaluates function arguments before the call.

```go
type Profile struct {
  Country    string
  PostalCode string
  Company    string
}

func ValidateProfile(input Profile, businessAccount bool) error {
  val := v.Is(
    v.String(input.Country, "country").Not().Blank(),
    v.String(input.PostalCode, "postal_code").Not().Blank(),
  )

  val.When(businessAccount, func(val *v.Validation) {
    val.Is(v.String(input.Company, "company").Not().Blank())
  })

  val.WhenAllValid([]string{"country", "postal_code"}, func(val *v.Validation) {
    if input.Country == "US" {
      val.Is(v.String(input.PostalCode, "postal_code").
        MatchingTo(regexp.MustCompile(`^[0-9]{5}(-[0-9]{4})?$`)))
    }
  })

  return val.ToError()
}
```

`PathValid` and related helpers query recorded invalid paths; they do not prove
that a path was validated. An unknown path is considered valid in v0.8.

## Errors and custom output

Keep `ToError()` for ordinary error flow. Use `ToValgoError()` only when the
caller needs structured paths and messages.

```go
func ValidationMessages(val *v.Validation) map[string][]string {
  errInfo := val.ToValgoError()
  if errInfo == nil {
    return nil
  }

  messages := make(map[string][]string, len(errInfo.Errors()))
  for path, valueErr := range errInfo.Errors() {
    messages[path] = append([]string(nil), valueErr.Messages()...)
  }
  return messages
}
```

Most rules accept a final custom message template:

```go
val := v.Is(
  v.String(input.Email, "email", "Email").
    Not().Blank("{{title}} is required"),
)
```

Preserve the consumer's established handler/status/JSON conversion rather than
returning Valgo's default JSON automatically.

## Localization and factories

Use `New(Options{...})` for one session and `Factory(FactoryOptions{...})` for
reused defaults. The APIs are `Factory` and `Options`; do not invent
`NewFactory` or option-builder functions.

```go
var frenchLocale = &v.Locale{
  v.ErrorKeyNotBlank: "{{title}} ne doit pas être vide",
}

var localizedFactory = v.Factory(v.FactoryOptions{
  LocaleCodeDefault: v.LocaleCodeEn,
  Locales: map[string]*v.Locale{
    "fr": frenchLocale,
  },
})

func ValidateLocalizedName(name, localeCode string) error {
  return localizedFactory.
    New(v.Options{LocaleCode: localeCode}).
    Is(v.String(name, "name", "Name").Not().Blank()).
    ToError()
}
```

In the current API, `ValidationFactory.Is` and `ValidationFactory.Check` each
accept one validator. Start with `factory.New(...).Is(validators...)` when a
factory-backed session needs several validators.

Use a factory-level `MarshalJsonFunc` only when the application deliberately
standardizes a custom Valgo error JSON shape.

## Reusable validations

Return `*v.Validation` from a repeated validation helper and compose it with a
namespace or `Merge`.

```go
func validateAddress(address Address) *v.Validation {
  return v.Is(
    v.String(address.City, "city").Not().Blank(),
    v.String(address.Street, "street").Not().Blank(),
  )
}

func ValidateShippingAddress(address Address) error {
  return v.In("shipping", validateAddress(address)).ToError()
}
```

Avoid a helper that only hides one call and has no repeated domain meaning.

## Custom validators

Wrap `*v.ValidatorContext`, implement `Context()`, and add typed rules. Supply a
fallback locale for a custom message key only on versions that support
`WithLocaleFallback`.

```go
const errorKeyEven = "even"

var evenLocale = &v.Locale{
  errorKeyEven: "{{title}} must be even",
}

type ValidatorEven struct {
  context *v.ValidatorContext
}

func Even(value int, nameAndTitle ...string) *ValidatorEven {
  context := v.NewContext(value, nameAndTitle...)
  context.WithLocaleFallback(evenLocale)
  return &ValidatorEven{context: context}
}

func (validator *ValidatorEven) Context() *v.ValidatorContext {
  return validator.context
}

func (validator *ValidatorEven) IsEven(template ...string) *ValidatorEven {
  value := validator.context.Value().(int)
  validator.context.Add(
    func() bool { return value%2 == 0 },
    errorKeyEven,
    template...,
  )
  return validator
}

func ValidatePairCount(count int) error {
  return v.Is(Even(count, "count", "Count").IsEven()).ToError()
}
```

Add wrapper methods delegating to `ValidatorContext.Not`, `Or`, or `OrElse`
only when users need those operators fluently on the custom validator.

## `Or()` and `OrElse()`

Use `Or()` for alternatives in one rule group:

```go
func ValidateStatus(status string) error {
  return v.Is(
    v.String(status, "status").
      EqualTo("draft").
      Or().EqualTo("published"),
  ).ToError()
}
```

In v0.8+, use `OrElse()` for “accept the left side, otherwise validate the
entire remaining chain”:

```go
var referralPattern = regexp.MustCompile(`^[A-Z0-9]{6,12}$`)

func ValidateOptionalReferral(code string) error {
  return v.Is(
    v.String(code, "referral_code").
      Empty().
      OrElse().
      MatchingTo(referralPattern),
  ).ToError()
}
```

`A.Or().B.C` means `(A OR B) AND C`. `A.OrElse().B.C` means
`A OR (B AND C)`, with a cut that skips the remainder when `A` succeeds.
Under `Check`, a successful left side still causes the `OrElse` cut.
