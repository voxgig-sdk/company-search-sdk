# Typed models for the CompanySearch SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class NearPoint:
    activite_principale: Optional[str] = None
    activite_principale_naf25: Optional[str] = None
    annee_categorie_entreprise: Optional[str] = None
    annee_tranche_effectif_salarie: Optional[str] = None
    caractere_employeur: Optional[str] = None
    categorie_entreprise: Optional[str] = None
    complement: Optional[dict] = None
    date_creation: Optional[str] = None
    date_fermeture: Optional[str] = None
    date_mise_a_jour: Optional[str] = None
    date_mise_a_jour_insee: Optional[str] = None
    date_mise_a_jour_rne: Optional[str] = None
    dirigeant: Optional[list] = None
    etat_administratif: Optional[str] = None
    finance: Optional[dict] = None
    matching_etablissement: Optional[list] = None
    nature_juridique: Optional[str] = None
    nom_complet: Optional[str] = None
    nom_raison_sociale: Optional[str] = None
    nombre_etablissement: Optional[int] = None
    nombre_etablissements_ouvert: Optional[int] = None
    section_activite_principale: Optional[str] = None
    siege: Optional[dict] = None
    sigle: Optional[str] = None
    siren: Optional[str] = None
    statut_diffusion: Optional[str] = None
    tranche_effectif_salarie: Optional[str] = None


@dataclass
class NearPointListMatch:
    activite_principale: Optional[str] = None
    activite_principale_naf25: Optional[str] = None
    annee_categorie_entreprise: Optional[str] = None
    annee_tranche_effectif_salarie: Optional[str] = None
    caractere_employeur: Optional[str] = None
    categorie_entreprise: Optional[str] = None
    complement: Optional[dict] = None
    date_creation: Optional[str] = None
    date_fermeture: Optional[str] = None
    date_mise_a_jour: Optional[str] = None
    date_mise_a_jour_insee: Optional[str] = None
    date_mise_a_jour_rne: Optional[str] = None
    dirigeant: Optional[list] = None
    etat_administratif: Optional[str] = None
    finance: Optional[dict] = None
    matching_etablissement: Optional[list] = None
    nature_juridique: Optional[str] = None
    nom_complet: Optional[str] = None
    nom_raison_sociale: Optional[str] = None
    nombre_etablissement: Optional[int] = None
    nombre_etablissements_ouvert: Optional[int] = None
    section_activite_principale: Optional[str] = None
    siege: Optional[dict] = None
    sigle: Optional[str] = None
    siren: Optional[str] = None
    statut_diffusion: Optional[str] = None
    tranche_effectif_salarie: Optional[str] = None


@dataclass
class Search:
    activite_principale: Optional[str] = None
    activite_principale_naf25: Optional[str] = None
    annee_categorie_entreprise: Optional[str] = None
    annee_tranche_effectif_salarie: Optional[str] = None
    caractere_employeur: Optional[str] = None
    categorie_entreprise: Optional[str] = None
    complement: Optional[dict] = None
    date_creation: Optional[str] = None
    date_fermeture: Optional[str] = None
    date_mise_a_jour: Optional[str] = None
    date_mise_a_jour_insee: Optional[str] = None
    date_mise_a_jour_rne: Optional[str] = None
    dirigeant: Optional[list] = None
    etat_administratif: Optional[str] = None
    finance: Optional[dict] = None
    matching_etablissement: Optional[list] = None
    nature_juridique: Optional[str] = None
    nom_complet: Optional[str] = None
    nom_raison_sociale: Optional[str] = None
    nombre_etablissement: Optional[int] = None
    nombre_etablissements_ouvert: Optional[int] = None
    section_activite_principale: Optional[str] = None
    siege: Optional[dict] = None
    sigle: Optional[str] = None
    siren: Optional[str] = None
    statut_diffusion: Optional[str] = None
    tranche_effectif_salarie: Optional[str] = None


@dataclass
class SearchListMatch:
    activite_principale: Optional[str] = None
    activite_principale_naf25: Optional[str] = None
    annee_categorie_entreprise: Optional[str] = None
    annee_tranche_effectif_salarie: Optional[str] = None
    caractere_employeur: Optional[str] = None
    categorie_entreprise: Optional[str] = None
    complement: Optional[dict] = None
    date_creation: Optional[str] = None
    date_fermeture: Optional[str] = None
    date_mise_a_jour: Optional[str] = None
    date_mise_a_jour_insee: Optional[str] = None
    date_mise_a_jour_rne: Optional[str] = None
    dirigeant: Optional[list] = None
    etat_administratif: Optional[str] = None
    finance: Optional[dict] = None
    matching_etablissement: Optional[list] = None
    nature_juridique: Optional[str] = None
    nom_complet: Optional[str] = None
    nom_raison_sociale: Optional[str] = None
    nombre_etablissement: Optional[int] = None
    nombre_etablissements_ouvert: Optional[int] = None
    section_activite_principale: Optional[str] = None
    siege: Optional[dict] = None
    sigle: Optional[str] = None
    siren: Optional[str] = None
    statut_diffusion: Optional[str] = None
    tranche_effectif_salarie: Optional[str] = None

