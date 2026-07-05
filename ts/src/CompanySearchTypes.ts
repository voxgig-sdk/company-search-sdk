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
  complement?: Record<string, any>
  date_creation?: string
  date_fermeture?: string
  date_mise_a_jour?: string
  date_mise_a_jour_insee?: string
  date_mise_a_jour_rne?: string
  dirigeant?: any[]
  etat_administratif?: string
  finance?: Record<string, any>
  matching_etablissement?: any[]
  nature_juridique?: string
  nom_complet?: string
  nom_raison_sociale?: string
  nombre_etablissement?: number
  nombre_etablissements_ouvert?: number
  section_activite_principale?: string
  siege?: Record<string, any>
  sigle?: string
  siren?: string
  statut_diffusion?: string
  tranche_effectif_salarie?: string
}

export interface NearPointListMatch {
  activite_principale?: string
  activite_principale_naf25?: string
  annee_categorie_entreprise?: string
  annee_tranche_effectif_salarie?: string
  caractere_employeur?: string
  categorie_entreprise?: string
  complement?: Record<string, any>
  date_creation?: string
  date_fermeture?: string
  date_mise_a_jour?: string
  date_mise_a_jour_insee?: string
  date_mise_a_jour_rne?: string
  dirigeant?: any[]
  etat_administratif?: string
  finance?: Record<string, any>
  matching_etablissement?: any[]
  nature_juridique?: string
  nom_complet?: string
  nom_raison_sociale?: string
  nombre_etablissement?: number
  nombre_etablissements_ouvert?: number
  section_activite_principale?: string
  siege?: Record<string, any>
  sigle?: string
  siren?: string
  statut_diffusion?: string
  tranche_effectif_salarie?: string
}

export interface Search {
  activite_principale?: string
  activite_principale_naf25?: string
  annee_categorie_entreprise?: string
  annee_tranche_effectif_salarie?: string
  caractere_employeur?: string
  categorie_entreprise?: string
  complement?: Record<string, any>
  date_creation?: string
  date_fermeture?: string
  date_mise_a_jour?: string
  date_mise_a_jour_insee?: string
  date_mise_a_jour_rne?: string
  dirigeant?: any[]
  etat_administratif?: string
  finance?: Record<string, any>
  matching_etablissement?: any[]
  nature_juridique?: string
  nom_complet?: string
  nom_raison_sociale?: string
  nombre_etablissement?: number
  nombre_etablissements_ouvert?: number
  section_activite_principale?: string
  siege?: Record<string, any>
  sigle?: string
  siren?: string
  statut_diffusion?: string
  tranche_effectif_salarie?: string
}

export interface SearchListMatch {
  activite_principale?: string
  activite_principale_naf25?: string
  annee_categorie_entreprise?: string
  annee_tranche_effectif_salarie?: string
  caractere_employeur?: string
  categorie_entreprise?: string
  complement?: Record<string, any>
  date_creation?: string
  date_fermeture?: string
  date_mise_a_jour?: string
  date_mise_a_jour_insee?: string
  date_mise_a_jour_rne?: string
  dirigeant?: any[]
  etat_administratif?: string
  finance?: Record<string, any>
  matching_etablissement?: any[]
  nature_juridique?: string
  nom_complet?: string
  nom_raison_sociale?: string
  nombre_etablissement?: number
  nombre_etablissements_ouvert?: number
  section_activite_principale?: string
  siege?: Record<string, any>
  sigle?: string
  siren?: string
  statut_diffusion?: string
  tranche_effectif_salarie?: string
}

