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
| `activite_principale` | `str` | No | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `str` | No | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `str` | No | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `str` | No | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `str` | No | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `str` | No | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `dict` | No |  |
| `date_creation` | `str` | No | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `str` | No | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `str` | No | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `str` | No | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `str` | No | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `list` | No |  |
| `etat_administratif` | `str` | No | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `dict` | No | Bilans financiers par année |
| `matching_etablissements` | `list` | No | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `str` | No | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `str` | No | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `str` | No | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` | No |  |
| `nombre_etablissements_ouverts` | `int` | No |  |
| `section_activite_principale` | `str` | No | Calculée à partir de l'activité principale. |
| `siege` | `dict` | No |  |
| `sigle` | `str` | No | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `str` | No | le numéro unique de l'entreprise |
| `statut_diffusion` | `str` | No | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `str` | No | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NearPoint().list({"lat": 1, "long": 1})
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
| `activite_principale` | `str` | No | Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE). |
| `activite_principale_naf25` | `str` | No | Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE). |
| `annee_categorie_entreprise` | `str` | No | Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE). |
| `annee_tranche_effectif_salarie` | `str` | No | Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE). |
| `caractere_employeur` | `str` | No | Caractère employeur de l'unité légale (source : base SIRENE). |
| `categorie_entreprise` | `str` | No | Catégorie d'entreprise de l'unité légale (source : base SIRENE). |
| `complements` | `dict` | No |  |
| `date_creation` | `str` | No | Date de création de l'unité légale (source : base SIRENE). |
| `date_fermeture` | `str` | No | Date de fermeture de l'unité légale (source : base historique SIRENE). |
| `date_mise_a_jour` | `str` | No | Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE). |
| `date_mise_a_jour_insee` | `str` | No | Date de la dernière mise à jour des données INSEE pour cette unité légale. |
| `date_mise_a_jour_rne` | `str` | No | Date de la dernière mise à jour des données RNCS pour cette unité légale. |
| `dirigeants` | `list` | No |  |
| `etat_administratif` | `str` | No | État administratif de l'unité légale (source : base SIRENE). |
| `finances` | `dict` | No | Bilans financiers par année |
| `matching_etablissements` | `list` | No | Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements. |
| `nature_juridique` | `str` | No | Catégorie juridique de l'unité légale (source : base SIRENE). |
| `nom_complet` | `str` | No | Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège. |
| `nom_raison_sociale` | `str` | No | La raison sociale pour les personnes morales (source : base SIRENE). |
| `nombre_etablissements` | `int` | No |  |
| `nombre_etablissements_ouverts` | `int` | No |  |
| `section_activite_principale` | `str` | No | Calculée à partir de l'activité principale. |
| `siege` | `dict` | No |  |
| `sigle` | `str` | No | Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE). |
| `siren` | `str` | No | le numéro unique de l'entreprise |
| `statut_diffusion` | `str` | No | Statut de diffusion de l'unité légale. |
| `tranche_effectif_salarie` | `str` | No | Tranche d'effectif salarié de l'unité légale (source : base SIRENE). |

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

