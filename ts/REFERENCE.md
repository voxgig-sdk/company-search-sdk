# CompanySearch TypeScript SDK Reference

Complete API reference for the CompanySearch TypeScript SDK.


## CompanySearchSDK

### Constructor

```ts
new CompanySearchSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CompanySearchSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = CompanySearchSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `CompanySearchSDK` instance in test mode.


### Instance Methods

#### `NearPoint(data?: object)`

Create a new `NearPoint` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NearPointEntity` instance.

#### `Search(data?: object)`

Create a new `Search` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SearchEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `CompanySearchSDK.test()`.

**Returns:** `CompanySearchSDK` instance in test mode.


---

## NearPointEntity

```ts
const near_point = client.NearPoint()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NearPoint().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NearPointEntity` instance with the same client and
options.

#### `client()`

Return the parent `CompanySearchSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SearchEntity

```ts
const search = client.Search()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Search().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `CompanySearchSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new CompanySearchSDK({
  feature: {
    test: { active: true },
  }
})
```

