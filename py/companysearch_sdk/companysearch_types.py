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


class NearPointListMatch(TypedDict, total=False):
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
