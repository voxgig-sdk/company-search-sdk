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
| `activite_principale` | `string` | No | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `string` | No | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `string` | No | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `string` | No | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `string` | No | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `string` | No | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `Record<string, any>` | No |  |
| `date_creation` | `string` | No | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | No | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | No | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | No | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | No | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `any[]` | No |  |
| `etat_administratif` | `string` | No | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `Record<string, any>` | No | Bilans financiers par année |
| `matching_etablissements` | `any[]` | No | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | No | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | No | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | No | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `number` | No |  |
| `nombre_etablissements_ouverts` | `number` | No |  |
| `section_activite_principale` | `string` | No | Calculée à partir de l'activité principale. |
| `siege` | `Record<string, any>` | No |  |
| `sigle` | `string` | No | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | No | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | No | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | No | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NearPoint().list({ lat: 1, long: 1 })
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
| `activite_principale` | `string` | No | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `string` | No | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `string` | No | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `string` | No | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `string` | No | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `string` | No | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `Record<string, any>` | No |  |
| `date_creation` | `string` | No | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | No | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | No | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | No | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | No | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `any[]` | No |  |
| `etat_administratif` | `string` | No | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `Record<string, any>` | No | Bilans financiers par année |
| `matching_etablissements` | `any[]` | No | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | No | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | No | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | No | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `number` | No |  |
| `nombre_etablissements_ouverts` | `number` | No |  |
| `section_activite_principale` | `string` | No | Calculée à partir de l'activité principale. |
| `siege` | `Record<string, any>` | No |  |
| `sigle` | `string` | No | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | No | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | No | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | No | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

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


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

