// Typed models for the CompanySearch SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/company-search-sdk/go/core"
)

// NearPoint is the typed data model for the near_point entity.
type NearPoint struct {
	ActivitePrincipale *string `json:"activite_principale,omitempty"`
	ActivitePrincipaleNaf25 *string `json:"activite_principale_naf25,omitempty"`
	AnneeCategorieEntreprise *string `json:"annee_categorie_entreprise,omitempty"`
	AnneeTrancheEffectifSalarie *string `json:"annee_tranche_effectif_salarie,omitempty"`
	CaractereEmployeur *string `json:"caractere_employeur,omitempty"`
	CategorieEntreprise *string `json:"categorie_entreprise,omitempty"`
	Complements *map[string]any `json:"complements,omitempty"`
	DateCreation *string `json:"date_creation,omitempty"`
	DateFermeture *string `json:"date_fermeture,omitempty"`
	DateMiseAJour *string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourInsee *string `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRne *string `json:"date_mise_a_jour_rne,omitempty"`
	Dirigeants *[]any `json:"dirigeants,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	Finances *map[string]any `json:"finances,omitempty"`
	MatchingEtablissements *[]any `json:"matching_etablissements,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomComplet *string `json:"nom_complet,omitempty"`
	NomRaisonSociale *string `json:"nom_raison_sociale,omitempty"`
	NombreEtablissements *int `json:"nombre_etablissements,omitempty"`
	NombreEtablissementsOuverts *int `json:"nombre_etablissements_ouverts,omitempty"`
	SectionActivitePrincipale *string `json:"section_activite_principale,omitempty"`
	Siege *map[string]any `json:"siege,omitempty"`
	Sigle *string `json:"sigle,omitempty"`
	Siren *string `json:"siren,omitempty"`
	StatutDiffusion *string `json:"statut_diffusion,omitempty"`
	TrancheEffectifSalarie *string `json:"tranche_effectif_salarie,omitempty"`
}

// NearPointListMatch is the typed request payload for NearPoint.ListTyped.
type NearPointListMatch struct {
	ActivitePrincipale *string `json:"activite_principale,omitempty"`
	ActivitePrincipaleNaf25 *string `json:"activite_principale_naf25,omitempty"`
	AnneeCategorieEntreprise *string `json:"annee_categorie_entreprise,omitempty"`
	AnneeTrancheEffectifSalarie *string `json:"annee_tranche_effectif_salarie,omitempty"`
	CaractereEmployeur *string `json:"caractere_employeur,omitempty"`
	CategorieEntreprise *string `json:"categorie_entreprise,omitempty"`
	Complements *map[string]any `json:"complements,omitempty"`
	DateCreation *string `json:"date_creation,omitempty"`
	DateFermeture *string `json:"date_fermeture,omitempty"`
	DateMiseAJour *string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourInsee *string `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRne *string `json:"date_mise_a_jour_rne,omitempty"`
	Dirigeants *[]any `json:"dirigeants,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	Finances *map[string]any `json:"finances,omitempty"`
	MatchingEtablissements *[]any `json:"matching_etablissements,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomComplet *string `json:"nom_complet,omitempty"`
	NomRaisonSociale *string `json:"nom_raison_sociale,omitempty"`
	NombreEtablissements *int `json:"nombre_etablissements,omitempty"`
	NombreEtablissementsOuverts *int `json:"nombre_etablissements_ouverts,omitempty"`
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
	Complements *map[string]any `json:"complements,omitempty"`
	DateCreation *string `json:"date_creation,omitempty"`
	DateFermeture *string `json:"date_fermeture,omitempty"`
	DateMiseAJour *string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourInsee *string `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRne *string `json:"date_mise_a_jour_rne,omitempty"`
	Dirigeants *[]any `json:"dirigeants,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	Finances *map[string]any `json:"finances,omitempty"`
	MatchingEtablissements *[]any `json:"matching_etablissements,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomComplet *string `json:"nom_complet,omitempty"`
	NomRaisonSociale *string `json:"nom_raison_sociale,omitempty"`
	NombreEtablissements *int `json:"nombre_etablissements,omitempty"`
	NombreEtablissementsOuverts *int `json:"nombre_etablissements_ouverts,omitempty"`
	SectionActivitePrincipale *string `json:"section_activite_principale,omitempty"`
	Siege *map[string]any `json:"siege,omitempty"`
	Sigle *string `json:"sigle,omitempty"`
	Siren *string `json:"siren,omitempty"`
	StatutDiffusion *string `json:"statut_diffusion,omitempty"`
	TrancheEffectifSalarie *string `json:"tranche_effectif_salarie,omitempty"`
}

// SearchListMatch is the typed request payload for Search.ListTyped.
type SearchListMatch struct {
	ActivitePrincipale *string `json:"activite_principale,omitempty"`
	ActivitePrincipaleNaf25 *string `json:"activite_principale_naf25,omitempty"`
	AnneeCategorieEntreprise *string `json:"annee_categorie_entreprise,omitempty"`
	AnneeTrancheEffectifSalarie *string `json:"annee_tranche_effectif_salarie,omitempty"`
	CaractereEmployeur *string `json:"caractere_employeur,omitempty"`
	CategorieEntreprise *string `json:"categorie_entreprise,omitempty"`
	Complements *map[string]any `json:"complements,omitempty"`
	DateCreation *string `json:"date_creation,omitempty"`
	DateFermeture *string `json:"date_fermeture,omitempty"`
	DateMiseAJour *string `json:"date_mise_a_jour,omitempty"`
	DateMiseAJourInsee *string `json:"date_mise_a_jour_insee,omitempty"`
	DateMiseAJourRne *string `json:"date_mise_a_jour_rne,omitempty"`
	Dirigeants *[]any `json:"dirigeants,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	Finances *map[string]any `json:"finances,omitempty"`
	MatchingEtablissements *[]any `json:"matching_etablissements,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomComplet *string `json:"nom_complet,omitempty"`
	NomRaisonSociale *string `json:"nom_raison_sociale,omitempty"`
	NombreEtablissements *int `json:"nombre_etablissements,omitempty"`
	NombreEtablissementsOuverts *int `json:"nombre_etablissements_ouverts,omitempty"`
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
