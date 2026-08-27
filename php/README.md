# CompanySearch PHP SDK



The PHP SDK for the CompanySearch API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->NearPoint()` — with named operations (`list`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/company-search-sdk/releases](https://github.com/voxgig-sdk/company-search-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'companysearch_sdk.php';

$client = new CompanySearchSDK();
```

### 2. List nearpoint records

```php
try {
    // list() returns an array of NearPoint records — iterate directly.
    $nearpoints = $client->NearPoint()->list();
    foreach ($nearpoints as $item) {
        echo $item["activite_principale"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $nearpoints = $client->NearPoint()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = CompanySearchSDK::test();

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$nearpoint = $client->NearPoint()->list();
print_r($nearpoint);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new CompanySearchSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
COMPANY_SEARCH_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### CompanySearchSDK

```php
require_once 'companysearch_sdk.php';
$client = new CompanySearchSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = CompanySearchSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### CompanySearchSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `NearPoint` | `($data): NearPointEntity` | Create a NearPoint entity instance. |
| `Search` | `($data): SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

### Entities

#### NearPoint

| Field | Description |
| --- | --- |
| `activite_principale` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` |  |
| `date_creation` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` |  |
| `etat_administratif` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | Bilans financiers par année |
| `matching_etablissements` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` |  |
| `nombre_etablissements_ouverts` |  |
| `section_activite_principale` | Calculée à partir de l'activité principale. |
| `siege` |  |
| `sigle` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | le numéro unique de l'entreprise |
| `statut_diffusion` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

Operations: List.

API path: `/near_point`

#### Search

| Field | Description |
| --- | --- |
| `activite_principale` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` |  |
| `date_creation` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` |  |
| `etat_administratif` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | Bilans financiers par année |
| `matching_etablissements` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` |  |
| `nombre_etablissements_ouverts` |  |
| `section_activite_principale` | Calculée à partir de l'activité principale. |
| `siege` |  |
| `sigle` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | le numéro unique de l'entreprise |
| `statut_diffusion` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

Operations: List.

API path: `/search`



## Entities


### NearPoint

Create an instance: `$near_point = $client->NearPoint();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activite_principale` | `string` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `string` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `string` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `string` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `string` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `string` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `array` |  |
| `date_creation` | `string` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `array` |  |
| `etat_administratif` | `string` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `array` | Bilans financiers par année |
| `matching_etablissements` | `array` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` |  |
| `nombre_etablissements_ouverts` | `int` |  |
| `section_activite_principale` | `string` | Calculée à partir de l'activité principale. |
| `siege` | `array` |  |
| `sigle` | `string` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

#### Example: List

```php
// list() returns an array of NearPoint records (throws on error).
$near_points = $client->NearPoint()->list();
```


### Search

Create an instance: `$search = $client->Search();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activite_principale` | `string` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `string` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `string` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `string` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `string` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `string` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `array` |  |
| `date_creation` | `string` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `array` |  |
| `etat_administratif` | `string` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `array` | Bilans financiers par année |
| `matching_etablissements` | `array` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` |  |
| `nombre_etablissements_ouverts` | `int` |  |
| `section_activite_principale` | `string` | Calculée à partir de l'activité principale. |
| `siege` | `array` |  |
| `sigle` | `string` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

#### Example: List

```php
// list() returns an array of Search records (throws on error).
$searchs = $client->Search()->list();
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── companysearch_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`companysearch_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$nearpoint = $client->NearPoint();
$nearpoint->list();

// $nearpoint->data_get() now returns the nearpoint data from the last list
// $nearpoint->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
