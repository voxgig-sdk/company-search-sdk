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
# @!attribute [rw] include
#   @return [String, nil]
#
# @!attribute [rw] lat
#   @return [Float]
#
# @!attribute [rw] limite_matching_etablissement
#   @return [Integer, nil]
#
# @!attribute [rw] long
#   @return [Float]
#
# @!attribute [rw] minimal
#   @return [Boolean, nil]
#
# @!attribute [rw] page
#   @return [Integer, nil]
#
# @!attribute [rw] page_etablissement
#   @return [Integer, nil]
#
# @!attribute [rw] per_page
#   @return [Integer, nil]
#
# @!attribute [rw] radius
#   @return [Float, nil]
#
# @!attribute [rw] section_activite_principale
#   @return [String, nil]
#
# @!attribute [rw] sort_by_size
#   @return [Boolean, nil]
NearPointListMatch = Struct.new(
  :activite_principale,
  :include,
  :lat,
  :limite_matching_etablissement,
  :long,
  :minimal,
  :page,
  :page_etablissement,
  :per_page,
  :radius,
  :section_activite_principale,
  :sort_by_size,
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
# @!attribute [rw] ca_max
#   @return [Integer, nil]
#
# @!attribute [rw] ca_min
#   @return [Integer, nil]
#
# @!attribute [rw] categorie_entreprise
#   @return [String, nil]
#
# @!attribute [rw] code_collectivite_territoriale
#   @return [String, nil]
#
# @!attribute [rw] code_commune
#   @return [String, nil]
#
# @!attribute [rw] code_postal
#   @return [String, nil]
#
# @!attribute [rw] convention_collective_renseignee
#   @return [Boolean, nil]
#
# @!attribute [rw] date_naissance_personne_max
#   @return [String, nil]
#
# @!attribute [rw] date_naissance_personne_min
#   @return [String, nil]
#
# @!attribute [rw] departement
#   @return [String, nil]
#
# @!attribute [rw] egapro_renseignee
#   @return [Boolean, nil]
#
# @!attribute [rw] epci
#   @return [String, nil]
#
# @!attribute [rw] est_achats_responsable
#   @return [Boolean, nil]
#
# @!attribute [rw] est_alim_confiance
#   @return [Boolean, nil]
#
# @!attribute [rw] est_association
#   @return [Boolean, nil]
#
# @!attribute [rw] est_bio
#   @return [Boolean, nil]
#
# @!attribute [rw] est_collectivite_territoriale
#   @return [Boolean, nil]
#
# @!attribute [rw] est_entrepreneur_individuel
#   @return [Boolean, nil]
#
# @!attribute [rw] est_entrepreneur_spectacle
#   @return [Boolean, nil]
#
# @!attribute [rw] est_ess
#   @return [Boolean, nil]
#
# @!attribute [rw] est_finess
#   @return [Boolean, nil]
#
# @!attribute [rw] est_l100_3
#   @return [Boolean, nil]
#
# @!attribute [rw] est_organisme_formation
#   @return [Boolean, nil]
#
# @!attribute [rw] est_patrimoine_vivant
#   @return [Boolean, nil]
#
# @!attribute [rw] est_qualiopi
#   @return [Boolean, nil]
#
# @!attribute [rw] est_rge
#   @return [Boolean, nil]
#
# @!attribute [rw] est_service_public
#   @return [Boolean, nil]
#
# @!attribute [rw] est_siae
#   @return [Boolean, nil]
#
# @!attribute [rw] est_societe_mission
#   @return [Boolean, nil]
#
# @!attribute [rw] est_uai
#   @return [Boolean, nil]
#
# @!attribute [rw] etat_administratif
#   @return [String, nil]
#
# @!attribute [rw] id_convention_collective
#   @return [String, nil]
#
# @!attribute [rw] id_finess
#   @return [String, nil]
#
# @!attribute [rw] id_rge
#   @return [String, nil]
#
# @!attribute [rw] id_uai
#   @return [String, nil]
#
# @!attribute [rw] include
#   @return [String, nil]
#
# @!attribute [rw] limite_matching_etablissement
#   @return [Integer, nil]
#
# @!attribute [rw] minimal
#   @return [Boolean, nil]
#
# @!attribute [rw] nature_juridique
#   @return [String, nil]
#
# @!attribute [rw] nom_personne
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [Integer, nil]
#
# @!attribute [rw] page_etablissement
#   @return [Integer, nil]
#
# @!attribute [rw] per_page
#   @return [Integer, nil]
#
# @!attribute [rw] prenoms_personne
#   @return [String, nil]
#
# @!attribute [rw] q
#   @return [String, nil]
#
# @!attribute [rw] region
#   @return [String, nil]
#
# @!attribute [rw] resultat_net_max
#   @return [Integer, nil]
#
# @!attribute [rw] resultat_net_min
#   @return [Integer, nil]
#
# @!attribute [rw] section_activite_principale
#   @return [String, nil]
#
# @!attribute [rw] sort_by_size
#   @return [Boolean, nil]
#
# @!attribute [rw] tranche_effectif_salarie
#   @return [String, nil]
#
# @!attribute [rw] type_personne
#   @return [String, nil]
SearchListMatch = Struct.new(
  :activite_principale,
  :ca_max,
  :ca_min,
  :categorie_entreprise,
  :code_collectivite_territoriale,
  :code_commune,
  :code_postal,
  :convention_collective_renseignee,
  :date_naissance_personne_max,
  :date_naissance_personne_min,
  :departement,
  :egapro_renseignee,
  :epci,
  :est_achats_responsable,
  :est_alim_confiance,
  :est_association,
  :est_bio,
  :est_collectivite_territoriale,
  :est_entrepreneur_individuel,
  :est_entrepreneur_spectacle,
  :est_ess,
  :est_finess,
  :est_l100_3,
  :est_organisme_formation,
  :est_patrimoine_vivant,
  :est_qualiopi,
  :est_rge,
  :est_service_public,
  :est_siae,
  :est_societe_mission,
  :est_uai,
  :etat_administratif,
  :id_convention_collective,
  :id_finess,
  :id_rge,
  :id_uai,
  :include,
  :limite_matching_etablissement,
  :minimal,
  :nature_juridique,
  :nom_personne,
  :page,
  :page_etablissement,
  :per_page,
  :prenoms_personne,
  :q,
  :region,
  :resultat_net_max,
  :resultat_net_min,
  :section_activite_principale,
  :sort_by_size,
  :tranche_effectif_salarie,
  :type_personne,
  keyword_init: true
)

