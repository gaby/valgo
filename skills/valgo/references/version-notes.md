# Version-risk notes

Read this file only to identify likely compatibility mistakes. Confirm every
detail against the effective module source and matching release docs resolved
by the core workflow. Do not use this reference to claim that a method exists.

## High-risk release boundaries

### v0.8.1

- Added preferred string methods `ByteLength`, `ByteLengthBetween`, `Length`,
  and `LengthBetween`.
- Kept `OfByteLength`, `OfByteLengthBetween`, `OfLength`, and
  `OfLengthBetween` as deprecated aliases until v1.0.
- Continue using the older `Of*` names when a consumer is pinned before v0.8.1.

### v0.8.0

- Raised the module's Go directive from Go 1.19 in v0.7 to Go 1.23.
- Added `OrElse()` and its cut behavior. v0.7 has `Or()` but no `OrElse()`.
- Changed failed `Or()` groups to one localized message joining alternatives.
- Added `PathValid()`, `AllValid()`, and `AnyValid()`. `IsValid()` remains as
  a deprecated compatibility method in v0.8.
- Added `IfValid`, `IfPathValid`, `IfAllValid`, `IfAnyValid` and matching
  `When*Valid` callbacks. In v0.7, use the supported `If`, `When`, `Do`, and
  `IsValid` APIs or ordinary Go control flow.
- Added `ValidatorContext.WithLocaleFallback()` for custom validator message
  defaults.

Treat OR grouping, error counts, and conditional query behavior as observable
changes; test them during an upgrade.

### v0.7.0

- Numeric validator types became generic families. Constructor call sites often
  remain similar, but code that names concrete validator types may need changes.

### v0.6.0

- String length rules changed to count runes rather than bytes. Use explicit
  byte-length methods when wire/storage size matters.

## Deprecated current APIs

- Prefer `ToError()` or `ToValgoError()` over `Validation.Error()`.
- Prefer `PathValid()` over `IsValid()` when the installed version supports it.
- Prefer the v0.8.1 short string length names only when that version supports
  them.

Do not proactively modernize these calls in a pinned consumer project unless
the replacement API exists there and the requested refactor permits the
behavior or naming change.
