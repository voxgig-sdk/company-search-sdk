# CompanySearch Golang SDK Reference

Complete API reference for the CompanySearch Golang SDK.


## CompanySearchSDK

### Constructor

```go
func NewCompanySearchSDK(options map[string]any) *CompanySearchSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *CompanySearchSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *CompanySearchSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `NearPoint(data map[string]any) CompanySearchEntity`

Create a new `NearPoint` entity instance. Pass `nil` for no initial data.

#### `Search(data map[string]any) CompanySearchEntity`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## NearPointEntity

```go
nearPoint := client.NearPoint(nil)
fmt.Println(nearPoint.GetName()) // "near_point"
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
| `complements` | `map[string]any` | No |  |
| `date_creation` | `string` | No | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | No | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | No | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | No | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | No | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `[]any` | No |  |
| `etat_administratif` | `string` | No | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `map[string]any` | No | Bilans financiers par année |
| `matching_etablissements` | `[]any` | No | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | No | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | No | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | No | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` | No |  |
| `nombre_etablissements_ouverts` | `int` | No |  |
| `section_activite_principale` | `string` | No | Calculée à partir de l'activité principale. |
| `siege` | `map[string]any` | No |  |
| `sigle` | `string` | No | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | No | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | No | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | No | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NearPoint(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NearPointEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SearchEntity

```go
search := client.Search(nil)
fmt.Println(search.GetName()) // "search"
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
| `complements` | `map[string]any` | No |  |
| `date_creation` | `string` | No | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `string` | No | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `string` | No | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `string` | No | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `string` | No | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `[]any` | No |  |
| `etat_administratif` | `string` | No | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `map[string]any` | No | Bilans financiers par année |
| `matching_etablissements` | `[]any` | No | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `string` | No | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `string` | No | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `string` | No | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` | No |  |
| `nombre_etablissements_ouverts` | `int` | No |  |
| `section_activite_principale` | `string` | No | Calculée à partir de l'activité principale. |
| `siege` | `map[string]any` | No |  |
| `sigle` | `string` | No | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `string` | No | le numéro unique de l'entreprise |
| `statut_diffusion` | `string` | No | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `string` | No | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewCompanySearchSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

