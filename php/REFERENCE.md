# CompanySearch PHP SDK Reference

Complete API reference for the CompanySearch PHP SDK.


## CompanySearchSDK

### Constructor

```php
require_once __DIR__ . '/companysearch_sdk.php';

$client = new CompanySearchSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CompanySearchSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = CompanySearchSDK::test();
```


### Instance Methods

#### `NearPoint($data = null)`

Create a new `NearPointEntity` instance. Pass `null` for no initial data.

#### `Search($data = null)`

Create a new `SearchEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): CompanySearchUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## NearPointEntity

```php
$near_point = $client->NearPoint();
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
| `complements` | `array` | No |  |
| `date_creation` | `string` | No |  |
| `date_fermeture` | `string` | No |  |
| `date_mise_a_jour` | `string` | No |  |
| `date_mise_a_jour_insee` | `string` | No |  |
| `date_mise_a_jour_rne` | `string` | No |  |
| `dirigeants` | `array` | No |  |
| `etat_administratif` | `string` | No |  |
| `finances` | `array` | No |  |
| `matching_etablissements` | `array` | No |  |
| `nature_juridique` | `string` | No |  |
| `nom_complet` | `string` | No |  |
| `nom_raison_sociale` | `string` | No |  |
| `nombre_etablissements` | `int` | No |  |
| `nombre_etablissements_ouverts` | `int` | No |  |
| `section_activite_principale` | `string` | No |  |
| `siege` | `array` | No |  |
| `sigle` | `string` | No |  |
| `siren` | `string` | No |  |
| `statut_diffusion` | `string` | No |  |
| `tranche_effectif_salarie` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NearPoint()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NearPointEntity`

Create a new `NearPointEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
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
| `complements` | `array` | No |  |
| `date_creation` | `string` | No |  |
| `date_fermeture` | `string` | No |  |
| `date_mise_a_jour` | `string` | No |  |
| `date_mise_a_jour_insee` | `string` | No |  |
| `date_mise_a_jour_rne` | `string` | No |  |
| `dirigeants` | `array` | No |  |
| `etat_administratif` | `string` | No |  |
| `finances` | `array` | No |  |
| `matching_etablissements` | `array` | No |  |
| `nature_juridique` | `string` | No |  |
| `nom_complet` | `string` | No |  |
| `nom_raison_sociale` | `string` | No |  |
| `nombre_etablissements` | `int` | No |  |
| `nombre_etablissements_ouverts` | `int` | No |  |
| `section_activite_principale` | `string` | No |  |
| `siege` | `array` | No |  |
| `sigle` | `string` | No |  |
| `siren` | `string` | No |  |
| `statut_diffusion` | `string` | No |  |
| `tranche_effectif_salarie` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Search()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new CompanySearchSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

