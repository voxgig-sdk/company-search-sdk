# CompanySearch Ruby SDK Reference

Complete API reference for the CompanySearch Ruby SDK.


## CompanySearchSDK

### Constructor

```ruby
require_relative 'CompanySearch_sdk'

client = CompanySearchSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
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

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

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

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## NearPointEntity

```ruby
near_point = client.NearPoint
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activite_principale` | `String` | No | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `String` | No | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `String` | No | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `String` | No | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `String` | No | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `String` | No | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `Hash` | No |  |
| `date_creation` | `String` | No | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `String` | No | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `String` | No | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `String` | No | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `String` | No | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `Array` | No |  |
| `etat_administratif` | `String` | No | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `Hash` | No | Bilans financiers par année |
| `matching_etablissements` | `Array` | No | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `String` | No | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `String` | No | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `String` | No | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `Integer` | No |  |
| `nombre_etablissements_ouverts` | `Integer` | No |  |
| `section_activite_principale` | `String` | No | Calculée à partir de l'activité principale. |
| `siege` | `Hash` | No |  |
| `sigle` | `String` | No | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `String` | No | le numéro unique de l'entreprise |
| `statut_diffusion` | `String` | No | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `String` | No | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NearPoint.list
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
| `activite_principale` | `String` | No | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `String` | No | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `String` | No | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `String` | No | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `String` | No | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `String` | No | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `Hash` | No |  |
| `date_creation` | `String` | No | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `String` | No | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `String` | No | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `String` | No | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `String` | No | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `Array` | No |  |
| `etat_administratif` | `String` | No | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `Hash` | No | Bilans financiers par année |
| `matching_etablissements` | `Array` | No | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `String` | No | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `String` | No | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `String` | No | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `Integer` | No |  |
| `nombre_etablissements_ouverts` | `Integer` | No |  |
| `section_activite_principale` | `String` | No | Calculée à partir de l'activité principale. |
| `siege` | `Hash` | No |  |
| `sigle` | `String` | No | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `String` | No | le numéro unique de l'entreprise |
| `statut_diffusion` | `String` | No | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `String` | No | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Search.list
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

