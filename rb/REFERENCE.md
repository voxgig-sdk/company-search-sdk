# CompanySearch Ruby SDK Reference

Complete API reference for the CompanySearch Ruby SDK.


## CompanySearchSDK

### Constructor

```ruby
require_relative 'company-search_sdk'

client = CompanySearchSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CompanySearchSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = CompanySearchSDK.test
```


### Instance Methods

#### `NearPoint(data = nil)`

Create a new `NearPoint` entity instance. Pass `nil` for no initial data.

#### `Search(data = nil)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash, err`

#### `prepare(fetchargs = {}) -> Hash, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Hash, err`


---

## NearPointEntity

```ruby
near_point = client.NearPoint
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

#### `list(reqmatch, ctrl = nil) -> result, err`

List entities matching the given criteria. Returns an array.

```ruby
results, err = client.NearPoint.list(nil)
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NearPointEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SearchEntity

```ruby
search = client.Search
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

#### `list(reqmatch, ctrl = nil) -> result, err`

List entities matching the given criteria. Returns an array.

```ruby
results, err = client.Search.list(nil)
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = CompanySearchSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

