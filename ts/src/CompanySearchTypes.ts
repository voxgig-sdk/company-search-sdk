// Typed models for the CompanySearch SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface NearPoint {
  activite_principale?: string
  activite_principale_naf25?: string
  annee_categorie_entreprise?: string
  annee_tranche_effectif_salarie?: string
  caractere_employeur?: string
  categorie_entreprise?: string
  complements?: Record<string, any>
  date_creation?: string
  date_fermeture?: string
  date_mise_a_jour?: string
  date_mise_a_jour_insee?: string
  date_mise_a_jour_rne?: string
  dirigeants?: any[]
  etat_administratif?: string
  finances?: Record<string, any>
  matching_etablissements?: any[]
  nature_juridique?: string
  nom_complet?: string
  nom_raison_sociale?: string
  nombre_etablissements?: number
  nombre_etablissements_ouverts?: number
  section_activite_principale?: string
  siege?: Record<string, any>
  sigle?: string
  siren?: string
  statut_diffusion?: string
  tranche_effectif_salarie?: string
}

export interface NearPointListMatch {
  activite_principale?: string
  include?: string
  lat: number
  limite_matching_etablissement?: number
  long: number
  minimal?: boolean
  page?: number
  page_etablissement?: number
  per_page?: number
  radius?: number
  section_activite_principale?: string
  sort_by_size?: boolean
}

export interface Search {
  activite_principale?: string
  activite_principale_naf25?: string
  annee_categorie_entreprise?: string
  annee_tranche_effectif_salarie?: string
  caractere_employeur?: string
  categorie_entreprise?: string
  complements?: Record<string, any>
  date_creation?: string
  date_fermeture?: string
  date_mise_a_jour?: string
  date_mise_a_jour_insee?: string
  date_mise_a_jour_rne?: string
  dirigeants?: any[]
  etat_administratif?: string
  finances?: Record<string, any>
  matching_etablissements?: any[]
  nature_juridique?: string
  nom_complet?: string
  nom_raison_sociale?: string
  nombre_etablissements?: number
  nombre_etablissements_ouverts?: number
  section_activite_principale?: string
  siege?: Record<string, any>
  sigle?: string
  siren?: string
  statut_diffusion?: string
  tranche_effectif_salarie?: string
}

export interface SearchListMatch {
  activite_principale?: string
  ca_max?: number
  ca_min?: number
  categorie_entreprise?: string
  code_collectivite_territoriale?: string
  code_commune?: string
  code_postal?: string
  convention_collective_renseignee?: boolean
  date_naissance_personne_max?: string
  date_naissance_personne_min?: string
  departement?: string
  egapro_renseignee?: boolean
  epci?: string
  est_achats_responsable?: boolean
  est_alim_confiance?: boolean
  est_association?: boolean
  est_bio?: boolean
  est_collectivite_territoriale?: boolean
  est_entrepreneur_individuel?: boolean
  est_entrepreneur_spectacle?: boolean
  est_ess?: boolean
  est_finess?: boolean
  est_l100_3?: boolean
  est_organisme_formation?: boolean
  est_patrimoine_vivant?: boolean
  est_qualiopi?: boolean
  est_rge?: boolean
  est_service_public?: boolean
  est_siae?: boolean
  est_societe_mission?: boolean
  est_uai?: boolean
  etat_administratif?: string
  id_convention_collective?: string
  id_finess?: string
  id_rge?: string
  id_uai?: string
  include?: string
  limite_matching_etablissement?: number
  minimal?: boolean
  nature_juridique?: string
  nom_personne?: string
  page?: number
  page_etablissement?: number
  per_page?: number
  prenoms_personne?: string
  q?: string
  region?: string
  resultat_net_max?: number
  resultat_net_min?: number
  section_activite_principale?: string
  sort_by_size?: boolean
  tranche_effectif_salarie?: string
  type_personne?: string
}

