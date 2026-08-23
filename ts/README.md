# CompanySearch TypeScript SDK



The TypeScript SDK for the CompanySearch API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.NearPoint()` — each with a small set of operations (`list`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/company-search-sdk/releases](https://github.com/voxgig-sdk/company-search-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { CompanySearchSDK } from '@voxgig-sdk/company-search'

const client = new CompanySearchSDK()
```

### 2. List nearpoint records

`list()` resolves to an array of NearPoint ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const nearpoints = await client.NearPoint().list()

for (const nearpoint of nearpoints) {
  console.log(nearpoint)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const nearpoints = await client.NearPoint().list()
  console.log(nearpoints)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = CompanySearchSDK.test()

const nearpoint = await client.NearPoint().list()
// nearpoint is the entity, populated with mock response data
// — call nearpoint.data() for the record itself
console.log(nearpoint)
```

You can also use the instance method:

```ts
const client = new CompanySearchSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.NearPoint()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new CompanySearchSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
COMPANY_SEARCH_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### CompanySearchSDK

#### Constructor

```ts
new CompanySearchSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `NearPoint(data?)` | `NearPointEntity` | Create a NearPoint entity instance. |
| `Search(data?)` | `SearchEntity` | Create a Search entity instance. |
| `tester(testopts?, sdkopts?)` | `CompanySearchSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `CompanySearchSDK.test(testopts?, sdkopts?)` | `CompanySearchSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): CompanySearchSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

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

Operations: list.

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

Operations: list.

API path: `/search`



## Entities


### NearPoint

Create an instance: `const near_point = client.NearPoint()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activite_principale` | `string` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `string` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `string` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `string` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `string` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `string` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `Record<string, any>` |  |
| `date_creation` | `string` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `any[]` |  |
| `etat_administratif` | `string` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `Record<string, any>` | Bilans financiers par année |
| `matching_etablissements` | `any[]` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `number` |  |
| `nombre_etablissements_ouverts` | `number` |  |
| `section_activite_principale` | `string` | Calculée à partir de l'activité principale. |
| `siege` | `Record<string, any>` |  |
| `sigle` | `string` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

#### Example: List

```ts
const near_points = await client.NearPoint().list()
```


### Search

Create an instance: `const search = client.Search()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activite_principale` | `string` | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `string` | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `string` | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `string` | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `string` | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `string` | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `Record<string, any>` |  |
| `date_creation` | `string` | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `any[]` |  |
| `etat_administratif` | `string` | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `Record<string, any>` | Bilans financiers par année |
| `matching_etablissements` | `any[]` | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `number` |  |
| `nombre_etablissements_ouverts` | `number` |  |
| `section_activite_principale` | `string` | Calculée à partir de l'activité principale. |
| `siege` | `Record<string, any>` |  |
| `sigle` | `string` | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

#### Example: List

```ts
const searchs = await client.Search().list()
```


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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
company-search/
├── src/
│   ├── CompanySearchSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { CompanySearchSDK } from '@voxgig-sdk/company-search'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const nearpoint = client.NearPoint()
await nearpoint.list()

// nearpoint.data() now returns the nearpoint data from the last `list`
// nearpoint.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
