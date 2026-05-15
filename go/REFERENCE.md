# CompanySearch Golang SDK Reference

Complete API reference for the CompanySearch Golang SDK.


## CompanySearchSDK

### Constructor

```go
func NewCompanySearchSDK(options map[string]any) *CompanySearchSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `TestSDK(testopts, sdkopts map[string]any) *CompanySearchSDK`

Create a test client with mock features active. Both arguments may be `nil`.

```go
client := sdk.TestSDK(nil, nil)
```


### Instance Methods

#### `NearPoint(data map[string]any) CompanySearchEntity`

Create a new `NearPoint` entity instance. Pass `nil` for no initial data.

#### `Search(data map[string]any) CompanySearchEntity`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## NearPointEntity

```go
near_point := client.NearPoint(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activite_principale` | ``$STRING`` | No |  |
| `activite_principale_naf25` | ``$STRING`` | No |  |
| `annee_categorie_entreprise` | ``$STRING`` | No |  |
| `annee_tranche_effectif_salarie` | ``$STRING`` | No |  |
| `caractere_employeur` | ``$STRING`` | No |  |
| `categorie_entreprise` | ``$STRING`` | No |  |
| `complement` | ``$OBJECT`` | No |  |
| `date_creation` | ``$STRING`` | No |  |
| `date_fermeture` | ``$STRING`` | No |  |
| `date_mise_a_jour` | ``$STRING`` | No |  |
| `date_mise_a_jour_insee` | ``$STRING`` | No |  |
| `date_mise_a_jour_rne` | ``$STRING`` | No |  |
| `dirigeant` | ``$ARRAY`` | No |  |
| `etat_administratif` | ``$STRING`` | No |  |
| `finance` | ``$OBJECT`` | No |  |
| `matching_etablissement` | ``$ARRAY`` | No |  |
| `nature_juridique` | ``$STRING`` | No |  |
| `nom_complet` | ``$STRING`` | No |  |
| `nom_raison_sociale` | ``$STRING`` | No |  |
| `nombre_etablissement` | ``$INTEGER`` | No |  |
| `nombre_etablissements_ouvert` | ``$INTEGER`` | No |  |
| `section_activite_principale` | ``$STRING`` | No |  |
| `siege` | ``$OBJECT`` | No |  |
| `sigle` | ``$STRING`` | No |  |
| `siren` | ``$STRING`` | No |  |
| `statut_diffusion` | ``$STRING`` | No |  |
| `tranche_effectif_salarie` | ``$STRING`` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NearPoint(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NearPointEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SearchEntity

```go
search := client.Search(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activite_principale` | ``$STRING`` | No |  |
| `activite_principale_naf25` | ``$STRING`` | No |  |
| `annee_categorie_entreprise` | ``$STRING`` | No |  |
| `annee_tranche_effectif_salarie` | ``$STRING`` | No |  |
| `caractere_employeur` | ``$STRING`` | No |  |
| `categorie_entreprise` | ``$STRING`` | No |  |
| `complement` | ``$OBJECT`` | No |  |
| `date_creation` | ``$STRING`` | No |  |
| `date_fermeture` | ``$STRING`` | No |  |
| `date_mise_a_jour` | ``$STRING`` | No |  |
| `date_mise_a_jour_insee` | ``$STRING`` | No |  |
| `date_mise_a_jour_rne` | ``$STRING`` | No |  |
| `dirigeant` | ``$ARRAY`` | No |  |
| `etat_administratif` | ``$STRING`` | No |  |
| `finance` | ``$OBJECT`` | No |  |
| `matching_etablissement` | ``$ARRAY`` | No |  |
| `nature_juridique` | ``$STRING`` | No |  |
| `nom_complet` | ``$STRING`` | No |  |
| `nom_raison_sociale` | ``$STRING`` | No |  |
| `nombre_etablissement` | ``$INTEGER`` | No |  |
| `nombre_etablissements_ouvert` | ``$INTEGER`` | No |  |
| `section_activite_principale` | ``$STRING`` | No |  |
| `siege` | ``$OBJECT`` | No |  |
| `sigle` | ``$STRING`` | No |  |
| `siren` | ``$STRING`` | No |  |
| `statut_diffusion` | ``$STRING`` | No |  |
| `tranche_effectif_salarie` | ``$STRING`` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Search(nil).List(nil, nil)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewCompanySearchSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

