# CompanySearch TypeScript SDK



The TypeScript SDK for the CompanySearch API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.NearPoint()` — each with a small set of operations (`list`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
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
| `activite_principale` |  |
| `activite_principale_naf25` |  |
| `annee_categorie_entreprise` |  |
| `annee_tranche_effectif_salarie` |  |
| `caractere_employeur` |  |
| `categorie_entreprise` |  |
| `complements` |  |
| `date_creation` |  |
| `date_fermeture` |  |
| `date_mise_a_jour` |  |
| `date_mise_a_jour_insee` |  |
| `date_mise_a_jour_rne` |  |
| `dirigeants` |  |
| `etat_administratif` |  |
| `finances` |  |
| `matching_etablissements` |  |
| `nature_juridique` |  |
| `nom_complet` |  |
| `nom_raison_sociale` |  |
| `nombre_etablissements` |  |
| `nombre_etablissements_ouverts` |  |
| `section_activite_principale` |  |
| `siege` |  |
| `sigle` |  |
| `siren` |  |
| `statut_diffusion` |  |
| `tranche_effectif_salarie` |  |

Operations: list.

API path: `/near_point`

#### Search

| Field | Description |
| --- | --- |
| `activite_principale` |  |
| `activite_principale_naf25` |  |
| `annee_categorie_entreprise` |  |
| `annee_tranche_effectif_salarie` |  |
| `caractere_employeur` |  |
| `categorie_entreprise` |  |
| `complements` |  |
| `date_creation` |  |
| `date_fermeture` |  |
| `date_mise_a_jour` |  |
| `date_mise_a_jour_insee` |  |
| `date_mise_a_jour_rne` |  |
| `dirigeants` |  |
| `etat_administratif` |  |
| `finances` |  |
| `matching_etablissements` |  |
| `nature_juridique` |  |
| `nom_complet` |  |
| `nom_raison_sociale` |  |
| `nombre_etablissements` |  |
| `nombre_etablissements_ouverts` |  |
| `section_activite_principale` |  |
| `siege` |  |
| `sigle` |  |
| `siren` |  |
| `statut_diffusion` |  |
| `tranche_effectif_salarie` |  |

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
| `activite_principale` | `string` |  |
| `activite_principale_naf25` | `string` |  |
| `annee_categorie_entreprise` | `string` |  |
| `annee_tranche_effectif_salarie` | `string` |  |
| `caractere_employeur` | `string` |  |
| `categorie_entreprise` | `string` |  |
| `complements` | `Record<string, any>` |  |
| `date_creation` | `string` |  |
| `date_fermeture` | `string` |  |
| `date_mise_a_jour` | `string` |  |
| `date_mise_a_jour_insee` | `string` |  |
| `date_mise_a_jour_rne` | `string` |  |
| `dirigeants` | `any[]` |  |
| `etat_administratif` | `string` |  |
| `finances` | `Record<string, any>` |  |
| `matching_etablissements` | `any[]` |  |
| `nature_juridique` | `string` |  |
| `nom_complet` | `string` |  |
| `nom_raison_sociale` | `string` |  |
| `nombre_etablissements` | `number` |  |
| `nombre_etablissements_ouverts` | `number` |  |
| `section_activite_principale` | `string` |  |
| `siege` | `Record<string, any>` |  |
| `sigle` | `string` |  |
| `siren` | `string` |  |
| `statut_diffusion` | `string` |  |
| `tranche_effectif_salarie` | `string` |  |

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
| `activite_principale` | `string` |  |
| `activite_principale_naf25` | `string` |  |
| `annee_categorie_entreprise` | `string` |  |
| `annee_tranche_effectif_salarie` | `string` |  |
| `caractere_employeur` | `string` |  |
| `categorie_entreprise` | `string` |  |
| `complements` | `Record<string, any>` |  |
| `date_creation` | `string` |  |
| `date_fermeture` | `string` |  |
| `date_mise_a_jour` | `string` |  |
| `date_mise_a_jour_insee` | `string` |  |
| `date_mise_a_jour_rne` | `string` |  |
| `dirigeants` | `any[]` |  |
| `etat_administratif` | `string` |  |
| `finances` | `Record<string, any>` |  |
| `matching_etablissements` | `any[]` |  |
| `nature_juridique` | `string` |  |
| `nom_complet` | `string` |  |
| `nom_raison_sociale` | `string` |  |
| `nombre_etablissements` | `number` |  |
| `nombre_etablissements_ouverts` | `number` |  |
| `section_activite_principale` | `string` |  |
| `siege` | `Record<string, any>` |  |
| `sigle` | `string` |  |
| `siren` | `string` |  |
| `statut_diffusion` | `string` |  |
| `tranche_effectif_salarie` | `string` |  |

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
