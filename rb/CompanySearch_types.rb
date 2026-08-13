# frozen_string_literal: true

# Typed models for the CompanySearch SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# NearPoint entity data model.
#
# @!attribute [rw] activite_principale
#   @return [String, nil]
#
# @!attribute [rw] activite_principale_naf25
#   @return [String, nil]
#
# @!attribute [rw] annee_categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] annee_tranche_effectif_salarie
#   @return [String, nil]
#
# @!attribute [rw] caractere_employeur
#   @return [String, nil]
#
# @!attribute [rw] categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] complements
#   @return [Hash, nil]
#
# @!attribute [rw] date_creation
#   @return [String, nil]
#
# @!attribute [rw] date_fermeture
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour_insee
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour_rne
#   @return [String, nil]
#
# @!attribute [rw] dirigeants
#   @return [Array, nil]
#
# @!attribute [rw] etat_administratif
#   @return [String, nil]
#
# @!attribute [rw] finances
#   @return [Hash, nil]
#
# @!attribute [rw] matching_etablissements
#   @return [Array, nil]
#
# @!attribute [rw] nature_juridique
#   @return [String, nil]
#
# @!attribute [rw] nom_complet
#   @return [String, nil]
#
# @!attribute [rw] nom_raison_sociale
#   @return [String, nil]
#
# @!attribute [rw] nombre_etablissements
#   @return [Integer, nil]
#
# @!attribute [rw] nombre_etablissements_ouverts
#   @return [Integer, nil]
#
# @!attribute [rw] section_activite_principale
#   @return [String, nil]
#
# @!attribute [rw] siege
#   @return [Hash, nil]
#
# @!attribute [rw] sigle
#   @return [String, nil]
#
# @!attribute [rw] siren
#   @return [String, nil]
#
# @!attribute [rw] statut_diffusion
#   @return [String, nil]
#
# @!attribute [rw] tranche_effectif_salarie
#   @return [String, nil]
NearPoint = Struct.new(
  :activite_principale,
  :activite_principale_naf25,
  :annee_categorie_entreprise,
  :annee_tranche_effectif_salarie,
  :caractere_employeur,
  :categorie_entreprise,
  :complements,
  :date_creation,
  :date_fermeture,
  :date_mise_a_jour,
  :date_mise_a_jour_insee,
  :date_mise_a_jour_rne,
  :dirigeants,
  :etat_administratif,
  :finances,
  :matching_etablissements,
  :nature_juridique,
  :nom_complet,
  :nom_raison_sociale,
  :nombre_etablissements,
  :nombre_etablissements_ouverts,
  :section_activite_principale,
  :siege,
  :sigle,
  :siren,
  :statut_diffusion,
  :tranche_effectif_salarie,
  keyword_init: true
)

# Request payload for NearPoint#list.
#
# @!attribute [rw] activite_principale
#   @return [String, nil]
#
# @!attribute [rw] activite_principale_naf25
#   @return [String, nil]
#
# @!attribute [rw] annee_categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] annee_tranche_effectif_salarie
#   @return [String, nil]
#
# @!attribute [rw] caractere_employeur
#   @return [String, nil]
#
# @!attribute [rw] categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] complements
#   @return [Hash, nil]
#
# @!attribute [rw] date_creation
#   @return [String, nil]
#
# @!attribute [rw] date_fermeture
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour_insee
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour_rne
#   @return [String, nil]
#
# @!attribute [rw] dirigeants
#   @return [Array, nil]
#
# @!attribute [rw] etat_administratif
#   @return [String, nil]
#
# @!attribute [rw] finances
#   @return [Hash, nil]
#
# @!attribute [rw] matching_etablissements
#   @return [Array, nil]
#
# @!attribute [rw] nature_juridique
#   @return [String, nil]
#
# @!attribute [rw] nom_complet
#   @return [String, nil]
#
# @!attribute [rw] nom_raison_sociale
#   @return [String, nil]
#
# @!attribute [rw] nombre_etablissements
#   @return [Integer, nil]
#
# @!attribute [rw] nombre_etablissements_ouverts
#   @return [Integer, nil]
#
# @!attribute [rw] section_activite_principale
#   @return [String, nil]
#
# @!attribute [rw] siege
#   @return [Hash, nil]
#
# @!attribute [rw] sigle
#   @return [String, nil]
#
# @!attribute [rw] siren
#   @return [String, nil]
#
# @!attribute [rw] statut_diffusion
#   @return [String, nil]
#
# @!attribute [rw] tranche_effectif_salarie
#   @return [String, nil]
NearPointListMatch = Struct.new(
  :activite_principale,
  :activite_principale_naf25,
  :annee_categorie_entreprise,
  :annee_tranche_effectif_salarie,
  :caractere_employeur,
  :categorie_entreprise,
  :complements,
  :date_creation,
  :date_fermeture,
  :date_mise_a_jour,
  :date_mise_a_jour_insee,
  :date_mise_a_jour_rne,
  :dirigeants,
  :etat_administratif,
  :finances,
  :matching_etablissements,
  :nature_juridique,
  :nom_complet,
  :nom_raison_sociale,
  :nombre_etablissements,
  :nombre_etablissements_ouverts,
  :section_activite_principale,
  :siege,
  :sigle,
  :siren,
  :statut_diffusion,
  :tranche_effectif_salarie,
  keyword_init: true
)

# Search entity data model.
#
# @!attribute [rw] activite_principale
#   @return [String, nil]
#
# @!attribute [rw] activite_principale_naf25
#   @return [String, nil]
#
# @!attribute [rw] annee_categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] annee_tranche_effectif_salarie
#   @return [String, nil]
#
# @!attribute [rw] caractere_employeur
#   @return [String, nil]
#
# @!attribute [rw] categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] complements
#   @return [Hash, nil]
#
# @!attribute [rw] date_creation
#   @return [String, nil]
#
# @!attribute [rw] date_fermeture
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour_insee
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour_rne
#   @return [String, nil]
#
# @!attribute [rw] dirigeants
#   @return [Array, nil]
#
# @!attribute [rw] etat_administratif
#   @return [String, nil]
#
# @!attribute [rw] finances
#   @return [Hash, nil]
#
# @!attribute [rw] matching_etablissements
#   @return [Array, nil]
#
# @!attribute [rw] nature_juridique
#   @return [String, nil]
#
# @!attribute [rw] nom_complet
#   @return [String, nil]
#
# @!attribute [rw] nom_raison_sociale
#   @return [String, nil]
#
# @!attribute [rw] nombre_etablissements
#   @return [Integer, nil]
#
# @!attribute [rw] nombre_etablissements_ouverts
#   @return [Integer, nil]
#
# @!attribute [rw] section_activite_principale
#   @return [String, nil]
#
# @!attribute [rw] siege
#   @return [Hash, nil]
#
# @!attribute [rw] sigle
#   @return [String, nil]
#
# @!attribute [rw] siren
#   @return [String, nil]
#
# @!attribute [rw] statut_diffusion
#   @return [String, nil]
#
# @!attribute [rw] tranche_effectif_salarie
#   @return [String, nil]
Search = Struct.new(
  :activite_principale,
  :activite_principale_naf25,
  :annee_categorie_entreprise,
  :annee_tranche_effectif_salarie,
  :caractere_employeur,
  :categorie_entreprise,
  :complements,
  :date_creation,
  :date_fermeture,
  :date_mise_a_jour,
  :date_mise_a_jour_insee,
  :date_mise_a_jour_rne,
  :dirigeants,
  :etat_administratif,
  :finances,
  :matching_etablissements,
  :nature_juridique,
  :nom_complet,
  :nom_raison_sociale,
  :nombre_etablissements,
  :nombre_etablissements_ouverts,
  :section_activite_principale,
  :siege,
  :sigle,
  :siren,
  :statut_diffusion,
  :tranche_effectif_salarie,
  keyword_init: true
)

# Request payload for Search#list.
#
# @!attribute [rw] activite_principale
#   @return [String, nil]
#
# @!attribute [rw] activite_principale_naf25
#   @return [String, nil]
#
# @!attribute [rw] annee_categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] annee_tranche_effectif_salarie
#   @return [String, nil]
#
# @!attribute [rw] caractere_employeur
#   @return [String, nil]
#
# @!attribute [rw] categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] complements
#   @return [Hash, nil]
#
# @!attribute [rw] date_creation
#   @return [String, nil]
#
# @!attribute [rw] date_fermeture
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour_insee
#   @return [String, nil]
#
# @!attribute [rw] date_mise_a_jour_rne
#   @return [String, nil]
#
# @!attribute [rw] dirigeants
#   @return [Array, nil]
#
# @!attribute [rw] etat_administratif
#   @return [String, nil]
#
# @!attribute [rw] finances
#   @return [Hash, nil]
#
# @!attribute [rw] matching_etablissements
#   @return [Array, nil]
#
# @!attribute [rw] nature_juridique
#   @return [String, nil]
#
# @!attribute [rw] nom_complet
#   @return [String, nil]
#
# @!attribute [rw] nom_raison_sociale
#   @return [String, nil]
#
# @!attribute [rw] nombre_etablissements
#   @return [Integer, nil]
#
# @!attribute [rw] nombre_etablissements_ouverts
#   @return [Integer, nil]
#
# @!attribute [rw] section_activite_principale
#   @return [String, nil]
#
# @!attribute [rw] siege
#   @return [Hash, nil]
#
# @!attribute [rw] sigle
#   @return [String, nil]
#
# @!attribute [rw] siren
#   @return [String, nil]
#
# @!attribute [rw] statut_diffusion
#   @return [String, nil]
#
# @!attribute [rw] tranche_effectif_salarie
#   @return [String, nil]
SearchListMatch = Struct.new(
  :activite_principale,
  :activite_principale_naf25,
  :annee_categorie_entreprise,
  :annee_tranche_effectif_salarie,
  :caractere_employeur,
  :categorie_entreprise,
  :complements,
  :date_creation,
  :date_fermeture,
  :date_mise_a_jour,
  :date_mise_a_jour_insee,
  :date_mise_a_jour_rne,
  :dirigeants,
  :etat_administratif,
  :finances,
  :matching_etablissements,
  :nature_juridique,
  :nom_complet,
  :nom_raison_sociale,
  :nombre_etablissements,
  :nombre_etablissements_ouverts,
  :section_activite_principale,
  :siege,
  :sigle,
  :siren,
  :statut_diffusion,
  :tranche_effectif_salarie,
  keyword_init: true
)

