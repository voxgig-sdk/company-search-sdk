# CompanySearch PHP SDK Reference

Complete API reference for the CompanySearch PHP SDK.


## CompanySearchSDK

### Constructor

```php
require_once __DIR__ . '/company-search_sdk.php';

$client = new CompanySearchSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
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

#### `optionsMap(): array`

Return a deep copy of the current SDK options.

#### `getUtility(): ProjectNameUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. Returns `[$result, $err]`.

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

**Returns:** `array [$result, $err]`

#### `prepare(array $fetchargs = []): array`

Prepare a fetch definition without sending the request. Returns `[$fetchdef, $err]`.


---

## NearPointEntity

```php
$near_point = $client->NearPoint();
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

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->NearPoint()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): NearPointEntity`

Create a new `NearPointEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
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

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->Search()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `getName(): string`

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

