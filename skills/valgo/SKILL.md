---
name: valgo
description: Add, refactor, debug, review, explain, or migrate type-safe validation in consumer Go applications using github.com/cohesivestack/valgo. Use for new or existing Valgo integrations, validator and rule selection, nested or conditional validation, errors and localization, reusable or custom validators, migrations from tag-based validation libraries, and version-compatible Valgo code generation and verification.
---

# Valgo

Use Valgo in an application that consumes the library. Do not use this skill to
develop, maintain, or contribute to Valgo itself. For upstream Valgo work,
follow that repository's contributor instructions and treat its source and
tests as the implementation under change.

## Inspect the consumer project

Start from the project root and inspect before editing:

1. Read repository instructions and locate `go.mod`, `go.work`, `vendor/`,
   relevant packages, tests, and generated-code policies.
2. Inspect package organization, request or domain types, JSON/form field
   naming, error handling, HTTP framework conventions, localization needs, and
   existing validation helpers.
3. Determine whether Valgo is installed and resolve the code actually built:

   ```bash
   go list -m -json github.com/cohesivestack/valgo
   ```

4. Record `Version`, `Dir`, `GoMod`, and any `Replace` value. Inspect vendored
   code or a local replacement when the build uses it; a tag alone is then not
   the effective API.
5. Preserve the installed version. Do not run an upgrade, change a replacement,
   or add a second validation library unless the user explicitly requests it.
6. For a new integration, identify the latest stable Valgo release whose
   `go.mod` is compatible with the project's Go version:

   ```bash
   go list -m -versions github.com/cohesivestack/valgo
   go list -m -json github.com/cohesivestack/valgo@vX.Y.Z
   ```

   Add the selected explicit version. Do not hard-code the version listed in
   this skill.

If `go list` reports that Valgo is absent, add it only when the requested work
requires a new integration or migration.

## Resolve API details in source-of-truth order

Use this priority:

1. The API in the effective Valgo source used by the consumer project
   (`Dir`, replacement, or vendor tree).
2. Valgo documentation from the same release or minor-version snapshot.
3. This skill's references.
4. Latest Valgo documentation only for a new integration or explicit upgrade.

Never silently use `master` APIs in a project pinned to an older release. If
documentation and source disagree, investigate the version boundary and follow
the effective source. Use `go doc`, source declarations, and package tests to
settle ambiguity. Never invent a constructor, rule, session method, option, or
error helper. If matching docs are not shipped with the effective source,
retrieve them from the exact release tag rather than the latest branch.

Read only the matching documentation needed for the task:

- Setup and sessions: `getting-started.md` and
  `using-valgo/validation-sessions.md`
- Validator selection and rules: `validators/overview.md`,
  `validators/rule-index.md`, and the relevant validator page
- Stateless boolean rules: `validators/predicates.md` when that page and the
  `github.com/cohesivestack/valgo/is` package exist in the effective version
- Nested values and slices: `using-valgo/namespaces.md`
- Result queries and conditional work:
  `using-valgo/querying-results.md` and
  `using-valgo/conditional-flows.md`
- Errors, localization, and factories: `using-valgo/errors.md` and
  `using-valgo/localization.md`
- OR behavior: `validators/or-operators.md` (or the matching older page)
- Extension and reusable patterns: `extending/custom-validators.md` and the
  relevant `cookbook/` page
- Release differences: `migration.md`

When versioned docs exist, select the directory matching the installed minor
release. Use unversioned latest docs only under priority 4 above.

## Load bundled references selectively

- Read [references/rule-index.md](references/rule-index.md) before choosing an
  unfamiliar constructor or translating rules from another library. Confirm
  every selected API against the installed version.
- Read [references/patterns.md](references/patterns.md) only for the requested
  feature: pointers, namespaces, slices, conditionals, errors, localization,
  factories, reuse, custom validators, `Or()`, or `OrElse()`.
- Read [references/version-notes.md](references/version-notes.md) when the
  project is pinned, replaced, vendored, being upgraded, or uses a deprecated
  name. Treat it as a risk checklist, not a changelog.

## Design the validation

1. Preserve the project's existing field paths, titles, message language,
   response schema, and boundary between validation and transport handling.
2. Prefer the conventional alias for a new import:

   ```go
   import v "github.com/cohesivestack/valgo"
   ```

   Preserve a different established alias.
3. Keep Valgo's function-based model. Do not introduce validation struct tags.
4. Select the narrowest validator compatible with the Go value's exact type.
   On versions with generalized numeric constructors, use `Int`/`IntP` for
   signed integers, `Uint`/`UintP` for unsigned integers, and `Float`/`FloatP`
   for floating-point values regardless of width. On older versions, use the
   width-specific constructor. Do not replace pointer absence with a sentinel
   zero value.
   For strings, use `EqualTo` for exact equality. Use `EqualFold` only when the
   effective version provides it and Unicode simple case-insensitive comparison
   is intended; it does not normalize text or perform full case folding.
5. Give validators stable machine-readable names. Use titles only when display
   text must differ from the field name.
6. Use `In` for nested objects, `InRow` for slices of objects, and `InCell` for
   slices of scalar values. Match the project's dot/index path convention.
7. Keep session operations separate from validator rules:
   - Build or extend sessions with `Is`, `Check`, `New`, `In`, `InRow`,
     `InCell`, `If`, `When`, `Do`, and `Merge` only when supported.
   - Chain rules such as `Not`, `Blank`, `Between`, or `Passing` on a
     type-compatible validator.
8. Use `Is` for the usual first-failure-per-validator behavior. Use `Check`
   only when the consumer needs every failing message in a chain. Use `New`
   for sessions assembled over multiple statements.
9. Treat `Or()` grouping and `OrElse()` cut semantics as version-sensitive.
   Verify both behavior and error shape in matching docs/tests before changing
   a chain.
10. Keep a one-off validation local. Extract a helper returning
    `*v.Validation` when the pattern genuinely repeats. Create a custom
    validator only for a reusable domain rule that benefits from its own typed
    API and message key.
11. When the effective Valgo version includes the `is` subpackage, use its
    stateless predicates for reusable boolean checks that do not need field
    paths, messages, localization, or validation-session results. A predicate
    may be passed to a validator's `Passing` method when an error is needed.

## Follow the task-specific workflow

### Add or change validation

Translate requirements into exact Go value types, field paths, rules, and
expected errors. Implement the smallest change that fits existing package and
handler conventions. Do not replace project error handling merely because
Valgo offers another output form.

### Review, debug, or refactor

Trace each input value through its constructor, rule chain, namespace, session
operation, and error conversion. Check pointer nil behavior, rune-versus-byte
length, OR precedence, conditional evaluation timing, unknown-path queries,
and `Is` versus `Check`. Preserve observable behavior during refactors unless
the user requests a behavior change.

### Migrate another validation library

1. Inventory tags, aliases, registered custom rules, cross-field checks,
   conditional rules, slice traversal, translations, and error response code.
2. Capture existing valid/invalid cases and error paths in tests before
   replacing behavior.
3. Map each rule explicitly to an installed Valgo API or an ordinary Go/custom
   predicate. Do not infer a same-named Valgo method from a source-library tag.
4. Convert `required` according to the Go type and desired whitespace
   semantics; convert `omitempty` using pointers or explicit Go conditions;
   convert slice traversal with `InRow` or `InCell`; express cross-field rules
   with typed Go values and conditional session flow.
5. Preserve JSON/form names and the external error contract. Remove the old
   dependency only when the requested migration is complete and no usages
   remain.

## Verify the result

1. Format every modified Go file with `gofmt`.
2. Compile and run the narrowest relevant tests, then run `go test ./...` when
   practical.
3. Run project-specific integration or handler tests that assert status codes,
   error paths, messages, locale, and JSON shape.
4. If dependency metadata changed, inspect `go.mod` and `go.sum`; run
   `go mod tidy` only when appropriate, and verify it did not upgrade Valgo or
   unrelated dependencies unexpectedly.
5. Search edited code for guessed or deprecated APIs and compare them with the
   effective source.
6. Review the final diff for accidental abstractions, unrelated edits, changed
   error conventions, and version drift. Report commands run and any checks
   that could not be completed.
