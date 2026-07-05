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
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
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
| `activite_principale` | `string` | No |  |
| `activite_principale_naf25` | `string` | No |  |
| `annee_categorie_entreprise` | `string` | No |  |
| `annee_tranche_effectif_salarie` | `string` | No |  |
| `caractere_employeur` | `string` | No |  |
| `categorie_entreprise` | `string` | No |  |
| `complement` | `table` | No |  |
| `date_creation` | `string` | No |  |
| `date_fermeture` | `string` | No |  |
| `date_mise_a_jour` | `string` | No |  |
| `date_mise_a_jour_insee` | `string` | No |  |
| `date_mise_a_jour_rne` | `string` | No |  |
| `dirigeant` | `table` | No |  |
| `etat_administratif` | `string` | No |  |
| `finance` | `table` | No |  |
| `matching_etablissement` | `table` | No |  |
| `nature_juridique` | `string` | No |  |
| `nom_complet` | `string` | No |  |
| `nom_raison_sociale` | `string` | No |  |
| `nombre_etablissement` | `number` | No |  |
| `nombre_etablissements_ouvert` | `number` | No |  |
| `section_activite_principale` | `string` | No |  |
| `siege` | `table` | No |  |
| `sigle` | `string` | No |  |
| `siren` | `string` | No |  |
| `statut_diffusion` | `string` | No |  |
| `tranche_effectif_salarie` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NearPoint():list()
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
| `activite_principale` | `string` | No |  |
| `activite_principale_naf25` | `string` | No |  |
| `annee_categorie_entreprise` | `string` | No |  |
| `annee_tranche_effectif_salarie` | `string` | No |  |
| `caractere_employeur` | `string` | No |  |
| `categorie_entreprise` | `string` | No |  |
| `complement` | `table` | No |  |
| `date_creation` | `string` | No |  |
| `date_fermeture` | `string` | No |  |
| `date_mise_a_jour` | `string` | No |  |
| `date_mise_a_jour_insee` | `string` | No |  |
| `date_mise_a_jour_rne` | `string` | No |  |
| `dirigeant` | `table` | No |  |
| `etat_administratif` | `string` | No |  |
| `finance` | `table` | No |  |
| `matching_etablissement` | `table` | No |  |
| `nature_juridique` | `string` | No |  |
| `nom_complet` | `string` | No |  |
| `nom_raison_sociale` | `string` | No |  |
| `nombre_etablissement` | `number` | No |  |
| `nombre_etablissements_ouvert` | `number` | No |  |
| `section_activite_principale` | `string` | No |  |
| `siege` | `table` | No |  |
| `sigle` | `string` | No |  |
| `siren` | `string` | No |  |
| `statut_diffusion` | `string` | No |  |
| `tranche_effectif_salarie` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search():list()
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

