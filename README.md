# strcase
[![Godoc Reference](https://godoc.org/github.com/ricdeau/strcase?status.svg)](http://godoc.org/github.com/ricdeau/strcase)
[![Go Report Card](https://goreportcard.com/badge/github.com/ricdeau/strcase)](https://goreportcard.com/report/github.com/ricdeau/strcase)

strcase is a go package for converting string case to various cases (e.g. [snake case](https://en.wikipedia.org/wiki/Snake_case) or [camel case](https://en.wikipedia.org/wiki/CamelCase)) to see the full conversion table below.

## Example

```go
s := "AnyKind of_string"
```

| Function                                  | Result               |
|-------------------------------------------|----------------------|
| `ToSnake(s)`                              | `any_kind_of_string` |
| `ToSnakeWithIgnore(s, '.')`               | `any_kind.of_string` |
| `ToScreamingSnake(s)`                     | `ANY_KIND_OF_STRING` |
| `ToKebab(s)`                              | `any-kind-of-string` |
| `ToScreamingKebab(s)`                     | `ANY-KIND-OF-STRING` |
| `ToDelimited(s, '.')`                     | `any.kind.of.string` |
| `ToScreamingDelimited(s, '.', '', true)`  | `ANY.KIND.OF.STRING` |
| `ToScreamingDelimited(s, '.', ' ', true)` | `ANY.KIND OF.STRING` |
| `ToCamel(s)`                              | `AnyKindOfString`    |
| `ToLowerCamel(s)`                         | `anyKindOfString`    |
| `ToFlat(s)`                               | `anykindofstring`    |
| `GoExported(s)`                           | `AnyKindOfString`    |
| `GoUnexported(s)`                         | `anyKindOfString`    |

## Install

```bash
go get -u github.com/ricdeau/strcase
```

## Custom Acronyms for ToCamel && ToLowerCamel

Often times text can contain specific acronyms which you need to be handled a certain way.
Out of the box `strcase` treats the string "ID" as "Id" or "id" but there is no way to cater
for every case in the wild.

To configure your custom acronym globally you can use the following before running any conversion

```go
import (
    "github.com/ricdeau/strcase"
)

func init() {
    // results in "Api" using ToCamel("API")
    // results in "api" using ToLowerCamel("API")
    strcase.ConfigureAcronym("API", "api")
    
    // results in "PostgreSQL" using ToCamel("PostgreSQL")
    // results in "postgreSQL" using ToLowerCamel("PostgreSQL")
    strcase.ConfigureAcronym("PostgreSQL", "PostgreSQL")

}

```