# Typed models for the CompanySearch SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class NearPoint(TypedDict, total=False):
    activite_principale: str
    activite_principale_naf25: str
    annee_categorie_entreprise: str
    annee_tranche_effectif_salarie: str
    caractere_employeur: str
    categorie_entreprise: str
    complements: dict
    date_creation: str
    date_fermeture: str
    date_mise_a_jour: str
    date_mise_a_jour_insee: str
    date_mise_a_jour_rne: str
    dirigeants: list
    etat_administratif: str
    finances: dict
    matching_etablissements: list
    nature_juridique: str
    nom_complet: str
    nom_raison_sociale: str
    nombre_etablissements: int
    nombre_etablissements_ouverts: int
    section_activite_principale: str
    siege: dict
    sigle: str
    siren: str
    statut_diffusion: str
    tranche_effectif_salarie: str


class NearPointListMatchRequired(TypedDict):
    lat: float
    long: float


class NearPointListMatch(NearPointListMatchRequired, total=False):
    activite_principale: str
    include: str
    limite_matching_etablissement: int
    minimal: bool
    page: int
    page_etablissement: int
    per_page: int
    radius: float
    section_activite_principale: str
    sort_by_size: bool


class Search(TypedDict, total=False):
    activite_principale: str
    activite_principale_naf25: str
    annee_categorie_entreprise: str
    annee_tranche_effectif_salarie: str
    caractere_employeur: str
    categorie_entreprise: str
    complements: dict
    date_creation: str
    date_fermeture: str
    date_mise_a_jour: str
    date_mise_a_jour_insee: str
    date_mise_a_jour_rne: str
    dirigeants: list
    etat_administratif: str
    finances: dict
    matching_etablissements: list
    nature_juridique: str
    nom_complet: str
    nom_raison_sociale: str
    nombre_etablissements: int
    nombre_etablissements_ouverts: int
    section_activite_principale: str
    siege: dict
    sigle: str
    siren: str
    statut_diffusion: str
    tranche_effectif_salarie: str


class SearchListMatch(TypedDict, total=False):
    activite_principale: str
    ca_max: int
    ca_min: int
    categorie_entreprise: str
    code_collectivite_territoriale: str
    code_commune: str
    code_postal: str
    convention_collective_renseignee: bool
    date_naissance_personne_max: str
    date_naissance_personne_min: str
    departement: str
    egapro_renseignee: bool
    epci: str
    est_achats_responsable: bool
    est_alim_confiance: bool
    est_association: bool
    est_bio: bool
    est_collectivite_territoriale: bool
    est_entrepreneur_individuel: bool
    est_entrepreneur_spectacle: bool
    est_ess: bool
    est_finess: bool
    est_l100_3: bool
    est_organisme_formation: bool
    est_patrimoine_vivant: bool
    est_qualiopi: bool
    est_rge: bool
    est_service_public: bool
    est_siae: bool
    est_societe_mission: bool
    est_uai: bool
    etat_administratif: str
    id_convention_collective: str
    id_finess: str
    id_rge: str
    id_uai: str
    include: str
    limite_matching_etablissement: int
    minimal: bool
    nature_juridique: str
    nom_personne: str
    page: int
    page_etablissement: int
    per_page: int
    prenoms_personne: str
    q: str
    region: str
    resultat_net_max: int
    resultat_net_min: int
    section_activite_principale: str
    sort_by_size: bool
    tranche_effectif_salarie: str
    type_personne: str
