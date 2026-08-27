# CompanySearch Golang SDK



The Golang SDK for the CompanySearch API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.NearPoint(nil)` — each with the same small set of operations (`List`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/company-search-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/company-search-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/company-search-sdk/go=../company-search-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    sdk "github.com/voxgig-sdk/company-search-sdk/go"
)

func main() {
    client := sdk.New()

    // List nearPoint records — the value is the array of records itself.
    nearPoints, err := client.NearPoint(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range nearPoints.([]any) {
        fmt.Println(item)
    }
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
nearpoints, err := client.NearPoint(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = nearpoints
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

nearPoint, err := client.NearPoint(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(nearPoint) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewCompanySearchSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
COMPANY_SEARCH_TEST_LIVE=TRUE
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewCompanySearchSDK

```go
func NewCompanySearchSDK(options map[string]any) *CompanySearchSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *CompanySearchSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CompanySearchSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `NearPoint` | `(data map[string]any) CompanySearchEntity` | Create a NearPoint entity instance. |
| `Search` | `(data map[string]any) CompanySearchEntity` | Create a Search entity instance. |

### Entity interface (CompanySearchEntity)

All entities implement the `CompanySearchEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    nearPoint, err := client.NearPoint(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // nearPoint is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### NearPoint

| Field | Description |
| --- | --- |
| `"activite_principale"` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `"activite_principale_naf25"` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `"annee_categorie_entreprise"` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `"annee_tranche_effectif_salarie"` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `"caractere_employeur"` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `"categorie_entreprise"` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `"complements"` |  |
| `"date_creation"` | Date de création de l'unité légale (source : base SIRENE). |
| `"date_fermeture"` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `"date_mise_a_jour"` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `"date_mise_a_jour_insee"` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `"date_mise_a_jour_rne"` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `"dirigeants"` |  |
| `"etat_administratif"` | État administratif de l'unité légale (source : base SIRENE). |
| `"finances"` | Bilans financiers par année |
| `"matching_etablissements"` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `"nature_juridique"` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `"nom_complet"` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `"nom_raison_sociale"` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `"nombre_etablissements"` |  |
| `"nombre_etablissements_ouverts"` |  |
| `"section_activite_principale"` | Calculée à partir de l'activité principale. |
| `"siege"` |  |
| `"sigle"` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `"siren"` | le numéro unique de l'entreprise |
| `"statut_diffusion"` | Statut de diffusion de l'unité légale. |
| `"tranche_effectif_salarie"` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

Operations: List.

API path: `/near_point`

#### Search

| Field | Description |
| --- | --- |
| `"activite_principale"` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `"activite_principale_naf25"` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `"annee_categorie_entreprise"` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `"annee_tranche_effectif_salarie"` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `"caractere_employeur"` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `"categorie_entreprise"` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `"complements"` |  |
| `"date_creation"` | Date de création de l'unité légale (source : base SIRENE). |
| `"date_fermeture"` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `"date_mise_a_jour"` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `"date_mise_a_jour_insee"` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `"date_mise_a_jour_rne"` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `"dirigeants"` |  |
| `"etat_administratif"` | État administratif de l'unité légale (source : base SIRENE). |
| `"finances"` | Bilans financiers par année |
| `"matching_etablissements"` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `"nature_juridique"` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `"nom_complet"` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `"nom_raison_sociale"` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `"nombre_etablissements"` |  |
| `"nombre_etablissements_ouverts"` |  |
| `"section_activite_principale"` | Calculée à partir de l'activité principale. |
| `"siege"` |  |
| `"sigle"` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `"siren"` | le numéro unique de l'entreprise |
| `"statut_diffusion"` | Statut de diffusion de l'unité légale. |
| `"tranche_effectif_salarie"` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

Operations: List.

API path: `/search`



## Entities


### NearPoint

Create an instance: `nearPoint := client.NearPoint(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activite_principale` | `string` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `string` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `string` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `string` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `string` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `string` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `map[string]any` |  |
| `date_creation` | `string` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `[]any` |  |
| `etat_administratif` | `string` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `map[string]any` | Bilans financiers par année |
| `matching_etablissements` | `[]any` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` |  |
| `nombre_etablissements_ouverts` | `int` |  |
| `section_activite_principale` | `string` | Calculée à partir de l'activité principale. |
| `siege` | `map[string]any` |  |
| `sigle` | `string` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

#### Example: List

```go
nearPoints, err := client.NearPoint(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(nearPoints) // the array of records
```


### Search

Create an instance: `search := client.Search(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activite_principale` | `string` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `string` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `string` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `string` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `string` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `string` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `map[string]any` |  |
| `date_creation` | `string` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `[]any` |  |
| `etat_administratif` | `string` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `map[string]any` | Bilans financiers par année |
| `matching_etablissements` | `[]any` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` |  |
| `nombre_etablissements_ouverts` | `int` |  |
| `section_activite_principale` | `string` | Calculée à partir de l'activité principale. |
| `siege` | `map[string]any` |  |
| `sigle` | `string` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

#### Example: List

```go
searchs, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(searchs) // the array of records
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | In-memory mock transport for testing without a live server |

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/company-search-sdk/go/
├── company-search.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/company-search-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
nearpoint := client.NearPoint(nil)
nearpoint.List(nil, nil)

// nearpoint.Data() now returns the nearpoint data from the last list
// nearpoint.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
