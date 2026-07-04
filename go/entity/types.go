// Typed models for the CompanySearch SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// NearPoint is the typed data model for the near_point entity.
type NearPoint struct {
	ActivitePrincipale *string `json:"activite_principale,omitempty"`
	ActivitePrincipaleNaf25 *string `json:"activite_principale_naf25,omitempty"`
	AnneeCategorieEntreprise *string `json:"annee_categorie_entreprise,omitempty"`
	AnneeTrancheEffectifSalarie *string `json:"annee_tranche_effectif_salarie,omitempty"`
	CaractereEmployeur *string `json:"caractere_employeur,omitempty"`
	CategorieEntreprise *string `json:"categorie_entreprise,omitempty"`
	Complement *map[string]any `json:"complement,omitempty"`
	DateCreation *string `json:"date_creation,omitempty"`
	DateFermeture *string `json:"date_fermeture,omitempty"`
	DateMiseAJour *string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourInsee *string `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRne *string `json:"date_mise_a_jour_rne,omitempty"`
	Dirigeant *[]any `json:"dirigeant,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	Finance *map[string]any `json:"finance,omitempty"`
	MatchingEtablissement *[]any `json:"matching_etablissement,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomComplet *string `json:"nom_complet,omitempty"`
	NomRaisonSociale *string `json:"nom_raison_sociale,omitempty"`
	NombreEtablissement *int `json:"nombre_etablissement,omitempty"`
	NombreEtablissementsOuvert *int `json:"nombre_etablissements_ouvert,omitempty"`
	SectionActivitePrincipale *string `json:"section_activite_principale,omitempty"`
	Siege *map[string]any `json:"siege,omitempty"`
	Sigle *string `json:"sigle,omitempty"`
	Siren *string `json:"siren,omitempty"`
	StatutDiffusion *string `json:"statut_diffusion,omitempty"`
	TrancheEffectifSalarie *string `json:"tranche_effectif_salarie,omitempty"`
}

// NearPointListMatch mirrors the near_point fields as an all-optional match
// filter (Go analog of Partial<NearPoint>).
type NearPointListMatch struct {
	ActivitePrincipale *string `json:"activite_principale,omitempty"`
	ActivitePrincipaleNaf25 *string `json:"activite_principale_naf25,omitempty"`
	AnneeCategorieEntreprise *string `json:"annee_categorie_entreprise,omitempty"`
	AnneeTrancheEffectifSalarie *string `json:"annee_tranche_effectif_salarie,omitempty"`
	CaractereEmployeur *string `json:"caractere_employeur,omitempty"`
	CategorieEntreprise *string `json:"categorie_entreprise,omitempty"`
	Complement *map[string]any `json:"complement,omitempty"`
	DateCreation *string `json:"date_creation,omitempty"`
	DateFermeture *string `json:"date_fermeture,omitempty"`
	DateMiseAJour *string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourInsee *string `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRne *string `json:"date_mise_a_jour_rne,omitempty"`
	Dirigeant *[]any `json:"dirigeant,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	Finance *map[string]any `json:"finance,omitempty"`
	MatchingEtablissement *[]any `json:"matching_etablissement,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomComplet *string `json:"nom_complet,omitempty"`
	NomRaisonSociale *string `json:"nom_raison_sociale,omitempty"`
	NombreEtablissement *int `json:"nombre_etablissement,omitempty"`
	NombreEtablissementsOuvert *int `json:"nombre_etablissements_ouvert,omitempty"`
	SectionActivitePrincipale *string `json:"section_activite_principale,omitempty"`
	Siege *map[string]any `json:"siege,omitempty"`
	Sigle *string `json:"sigle,omitempty"`
	Siren *string `json:"siren,omitempty"`
	StatutDiffusion *string `json:"statut_diffusion,omitempty"`
	TrancheEffectifSalarie *string `json:"tranche_effectif_salarie,omitempty"`
}

// Search is the typed data model for the search entity.
type Search struct {
	ActivitePrincipale *string `json:"activite_principale,omitempty"`
	ActivitePrincipaleNaf25 *string `json:"activite_principale_naf25,omitempty"`
	AnneeCategorieEntreprise *string `json:"annee_categorie_entreprise,omitempty"`
	AnneeTrancheEffectifSalarie *string `json:"annee_tranche_effectif_salarie,omitempty"`
	CaractereEmployeur *string `json:"caractere_employeur,omitempty"`
	CategorieEntreprise *string `json:"categorie_entreprise,omitempty"`
	Complement *map[string]any `json:"complement,omitempty"`
	DateCreation *string `json:"date_creation,omitempty"`
	DateFermeture *string `json:"date_fermeture,omitempty"`
	DateMiseAJour *string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourInsee *string `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRne *string `json:"date_mise_a_jour_rne,omitempty"`
	Dirigeant *[]any `json:"dirigeant,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	Finance *map[string]any `json:"finance,omitempty"`
	MatchingEtablissement *[]any `json:"matching_etablissement,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomComplet *string `json:"nom_complet,omitempty"`
	NomRaisonSociale *string `json:"nom_raison_sociale,omitempty"`
	NombreEtablissement *int `json:"nombre_etablissement,omitempty"`
	NombreEtablissementsOuvert *int `json:"nombre_etablissements_ouvert,omitempty"`
	SectionActivitePrincipale *string `json:"section_activite_principale,omitempty"`
	Siege *map[string]any `json:"siege,omitempty"`
	Sigle *string `json:"sigle,omitempty"`
	Siren *string `json:"siren,omitempty"`
	StatutDiffusion *string `json:"statut_diffusion,omitempty"`
	TrancheEffectifSalarie *string `json:"tranche_effectif_salarie,omitempty"`
}

// SearchListMatch mirrors the search fields as an all-optional match
// filter (Go analog of Partial<Search>).
type SearchListMatch struct {
	ActivitePrincipale *string `json:"activite_principale,omitempty"`
	ActivitePrincipaleNaf25 *string `json:"activite_principale_naf25,omitempty"`
	AnneeCategorieEntreprise *string `json:"annee_categorie_entreprise,omitempty"`
	AnneeTrancheEffectifSalarie *string `json:"annee_tranche_effectif_salarie,omitempty"`
	CaractereEmployeur *string `json:"caractere_employeur,omitempty"`
	CategorieEntreprise *string `json:"categorie_entreprise,omitempty"`
	Complement *map[string]any `json:"complement,omitempty"`
	DateCreation *string `json:"date_creation,omitempty"`
	DateFermeture *string `json:"date_fermeture,omitempty"`
	DateMiseAJour *string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourInsee *string `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRne *string `json:"date_mise_a_jour_rne,omitempty"`
	Dirigeant *[]any `json:"dirigeant,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	Finance *map[string]any `json:"finance,omitempty"`
	MatchingEtablissement *[]any `json:"matching_etablissement,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomComplet *string `json:"nom_complet,omitempty"`
	NomRaisonSociale *string `json:"nom_raison_sociale,omitempty"`
	NombreEtablissement *int `json:"nombre_etablissement,omitempty"`
	NombreEtablissementsOuvert *int `json:"nombre_etablissements_ouvert,omitempty"`
	SectionActivitePrincipale *string `json:"section_activite_principale,omitempty"`
	Siege *map[string]any `json:"siege,omitempty"`
	Sigle *string `json:"sigle,omitempty"`
	Siren *string `json:"siren,omitempty"`
	StatutDiffusion *string `json:"statut_diffusion,omitempty"`
	TrancheEffectifSalarie *string `json:"tranche_effectif_salarie,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
