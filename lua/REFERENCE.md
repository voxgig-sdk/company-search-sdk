# CompanySearch Lua SDK Reference

Complete API reference for the CompanySearch Lua SDK.


## CompanySearchSDK

### Constructor

```lua
local sdk = require("company-search_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts, sdkopts)`

Create a test client with mock features active. Both arguments may be `nil`.

```lua
local client = sdk.test(nil, nil)
```


### Instance Methods

#### `NearPoint(data)`

Create a new `NearPoint` entity instance. Pass `nil` for no initial data.

#### `Search(data)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## NearPointEntity

```lua
local near_point = client:NearPoint(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NearPoint(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NearPointEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SearchEntity

```lua
local search = client:Search(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

