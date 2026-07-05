# CompanySearch Python SDK Reference

Complete API reference for the CompanySearch Python SDK.


## CompanySearchSDK

### Constructor

```python
from companysearch_sdk import CompanySearchSDK

client = CompanySearchSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CompanySearchSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = CompanySearchSDK.test()
```


### Instance Methods

#### `NearPoint(data=None)`

Create a new `NearPointEntity` instance. Pass `None` for no initial data.

#### `Search(data=None)`

Create a new `SearchEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## NearPointEntity

```python
near_point = client.NearPoint()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activite_principale` | `str` | No |  |
| `activite_principale_naf25` | `str` | No |  |
| `annee_categorie_entreprise` | `str` | No |  |
| `annee_tranche_effectif_salarie` | `str` | No |  |
| `caractere_employeur` | `str` | No |  |
| `categorie_entreprise` | `str` | No |  |
| `complement` | `dict` | No |  |
| `date_creation` | `str` | No |  |
| `date_fermeture` | `str` | No |  |
| `date_mise_a_jour` | `str` | No |  |
| `date_mise_a_jour_insee` | `str` | No |  |
| `date_mise_a_jour_rne` | `str` | No |  |
| `dirigeant` | `list` | No |  |
| `etat_administratif` | `str` | No |  |
| `finance` | `dict` | No |  |
| `matching_etablissement` | `list` | No |  |
| `nature_juridique` | `str` | No |  |
| `nom_complet` | `str` | No |  |
| `nom_raison_sociale` | `str` | No |  |
| `nombre_etablissement` | `int` | No |  |
| `nombre_etablissements_ouvert` | `int` | No |  |
| `section_activite_principale` | `str` | No |  |
| `siege` | `dict` | No |  |
| `sigle` | `str` | No |  |
| `siren` | `str` | No |  |
| `statut_diffusion` | `str` | No |  |
| `tranche_effectif_salarie` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NearPoint().list()
for near_point in results:
    print(near_point)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NearPointEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SearchEntity

```python
search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activite_principale` | `str` | No |  |
| `activite_principale_naf25` | `str` | No |  |
| `annee_categorie_entreprise` | `str` | No |  |
| `annee_tranche_effectif_salarie` | `str` | No |  |
| `caractere_employeur` | `str` | No |  |
| `categorie_entreprise` | `str` | No |  |
| `complement` | `dict` | No |  |
| `date_creation` | `str` | No |  |
| `date_fermeture` | `str` | No |  |
| `date_mise_a_jour` | `str` | No |  |
| `date_mise_a_jour_insee` | `str` | No |  |
| `date_mise_a_jour_rne` | `str` | No |  |
| `dirigeant` | `list` | No |  |
| `etat_administratif` | `str` | No |  |
| `finance` | `dict` | No |  |
| `matching_etablissement` | `list` | No |  |
| `nature_juridique` | `str` | No |  |
| `nom_complet` | `str` | No |  |
| `nom_raison_sociale` | `str` | No |  |
| `nombre_etablissement` | `int` | No |  |
| `nombre_etablissements_ouvert` | `int` | No |  |
| `section_activite_principale` | `str` | No |  |
| `siege` | `dict` | No |  |
| `sigle` | `str` | No |  |
| `siren` | `str` | No |  |
| `statut_diffusion` | `str` | No |  |
| `tranche_effectif_salarie` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Search().list()
for search in results:
    print(search)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = CompanySearchSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

