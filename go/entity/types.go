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
	Include *string `json:"include,omitempty"`
	Lat float64 `json:"lat"`
	LimiteMatchingEtablissement *int `json:"limite_matching_etablissement,omitempty"`
	Long float64 `json:"long"`
	Minimal *bool `json:"minimal,omitempty"`
	Page *int `json:"page,omitempty"`
	PageEtablissement *int `json:"page_etablissement,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	Radius *float64 `json:"radius,omitempty"`
	SectionActivitePrincipale *string `json:"section_activite_principale,omitempty"`
	SortBySize *bool `json:"sort_by_size,omitempty"`
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
	CaMax *int `json:"ca_max,omitempty"`
	CaMin *int `json:"ca_min,omitempty"`
	CategorieEntreprise *string `json:"categorie_entreprise,omitempty"`
	CodeCollectiviteTerritoriale *string `json:"code_collectivite_territoriale,omitempty"`
	CodeCommune *string `json:"code_commune,omitempty"`
	CodePostal *string `json:"code_postal,omitempty"`
	ConventionCollectiveRenseignee *bool `json:"convention_collective_renseignee,omitempty"`
	DateNaissancePersonneMax *string `json:"date_naissance_personne_max,omitempty"`
	DateNaissancePersonneMin *string `json:"date_naissance_personne_min,omitempty"`
	Departement *string `json:"departement,omitempty"`
	EgaproRenseignee *bool `json:"egapro_renseignee,omitempty"`
	Epci *string `json:"epci,omitempty"`
	EstAchatsResponsable *bool `json:"est_achats_responsable,omitempty"`
	EstAlimConfiance *bool `json:"est_alim_confiance,omitempty"`
	EstAssociation *bool `json:"est_association,omitempty"`
	EstBio *bool `json:"est_bio,omitempty"`
	EstCollectiviteTerritoriale *bool `json:"est_collectivite_territoriale,omitempty"`
	EstEntrepreneurIndividuel *bool `json:"est_entrepreneur_individuel,omitempty"`
	EstEntrepreneurSpectacle *bool `json:"est_entrepreneur_spectacle,omitempty"`
	EstEss *bool `json:"est_ess,omitempty"`
	EstFiness *bool `json:"est_finess,omitempty"`
	EstL1003 *bool `json:"est_l100_3,omitempty"`
	EstOrganismeFormation *bool `json:"est_organisme_formation,omitempty"`
	EstPatrimoineVivant *bool `json:"est_patrimoine_vivant,omitempty"`
	EstQualiopi *bool `json:"est_qualiopi,omitempty"`
	EstRge *bool `json:"est_rge,omitempty"`
	EstServicePublic *bool `json:"est_service_public,omitempty"`
	EstSiae *bool `json:"est_siae,omitempty"`
	EstSocieteMission *bool `json:"est_societe_mission,omitempty"`
	EstUai *bool `json:"est_uai,omitempty"`
	EtatAdministratif *string `json:"etat_administratif,omitempty"`
	IdConventionCollective *string `json:"id_convention_collective,omitempty"`
	IdFiness *string `json:"id_finess,omitempty"`
	IdRge *string `json:"id_rge,omitempty"`
	IdUai *string `json:"id_uai,omitempty"`
	Include *string `json:"include,omitempty"`
	LimiteMatchingEtablissement *int `json:"limite_matching_etablissement,omitempty"`
	Minimal *bool `json:"minimal,omitempty"`
	NatureJuridique *string `json:"nature_juridique,omitempty"`
	NomPersonne *string `json:"nom_personne,omitempty"`
	Page *int `json:"page,omitempty"`
	PageEtablissement *int `json:"page_etablissement,omitempty"`
	PerPage *int `json:"per_page,omitempty"`
	PrenomsPersonne *string `json:"prenoms_personne,omitempty"`
	Q *string `json:"q,omitempty"`
	Region *string `json:"region,omitempty"`
	ResultatNetMax *int `json:"resultat_net_max,omitempty"`
	ResultatNetMin *int `json:"resultat_net_min,omitempty"`
	SectionActivitePrincipale *string `json:"section_activite_principale,omitempty"`
	SortBySize *bool `json:"sort_by_size,omitempty"`
	TrancheEffectifSalarie *string `json:"tranche_effectif_salarie,omitempty"`
	TypePersonne *string `json:"type_personne,omitempty"`
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
