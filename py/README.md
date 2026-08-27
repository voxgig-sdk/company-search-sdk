# CompanySearch Python SDK



The Python SDK for the CompanySearch API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.NearPoint()` — each
carrying a small, uniform set of operations (`list`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/company-search-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
from companysearch_sdk import CompanySearchSDK

client = CompanySearchSDK()
```

### 2. List nearpoint records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    nearpoints = client.NearPoint().list({"lat": 1, "long": 1})
    for nearpoint in nearpoints:
        print(nearpoint)
except Exception as err:
    print(f"list failed: {err}")
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    nearpoints = client.NearPoint().list()
    print(nearpoints)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = CompanySearchSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
nearpoint = client.NearPoint().list()
# nearpoint contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = CompanySearchSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### CompanySearchSDK

```python
from companysearch_sdk import CompanySearchSDK

client = CompanySearchSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = CompanySearchSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### CompanySearchSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `NearPoint` | `(data) -> NearPointEntity` | Create a NearPoint entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `near_point = client.NearPoint()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activite_principale` | `str` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `str` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `str` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `str` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `str` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `str` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `dict` |  |
| `date_creation` | `str` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `str` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `str` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `str` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `str` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `list` |  |
| `etat_administratif` | `str` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `dict` | Bilans financiers par année |
| `matching_etablissements` | `list` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `str` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `str` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `str` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` |  |
| `nombre_etablissements_ouverts` | `int` |  |
| `section_activite_principale` | `str` | Calculée à partir de l'activité principale. |
| `siege` | `dict` |  |
| `sigle` | `str` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `str` | le numéro unique de l'entreprise |
| `statut_diffusion` | `str` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `str` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

#### Example: List

```python
near_points = client.NearPoint().list({"lat": 1, "long": 1})
```


### Search

Create an instance: `search = client.Search()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activite_principale` | `str` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `str` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `str` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `str` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `str` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `str` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `dict` |  |
| `date_creation` | `str` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `str` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `str` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `str` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `str` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `list` |  |
| `etat_administratif` | `str` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `dict` | Bilans financiers par année |
| `matching_etablissements` | `list` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `str` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `str` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `str` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` |  |
| `nombre_etablissements_ouverts` | `int` |  |
| `section_activite_principale` | `str` | Calculée à partir de l'activité principale. |
| `siege` | `dict` |  |
| `sigle` | `str` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `str` | le numéro unique de l'entreprise |
| `statut_diffusion` | `str` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `str` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

#### Example: List

```python
searchs = client.Search().list()
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── companysearch_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`companysearch_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
nearpoint = client.NearPoint()
nearpoint.list()

# nearpoint.data_get() now returns the nearpoint data from the last list
# nearpoint.match_get() returns the last match criteria
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
