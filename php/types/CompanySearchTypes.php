<?php
declare(strict_types=1);

// Typed models for the CompanySearch SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** NearPoint entity data model. */
class NearPoint
{
    public ?string $activite_principale = null;
    public ?string $activite_principale_naf25 = null;
    public ?string $annee_categorie_entreprise = null;
    public ?string $annee_tranche_effectif_salarie = null;
    public ?string $caractere_employeur = null;
    public ?string $categorie_entreprise = null;
    public ?array $complements = null;
    public ?string $date_creation = null;
    public ?string $date_fermeture = null;
    public ?string $date_mise_a_jour = null;
    public ?string $date_mise_a_jour_insee = null;
    public ?string $date_mise_a_jour_rne = null;
    public ?array $dirigeants = null;
    public ?string $etat_administratif = null;
    public ?array $finances = null;
    public ?array $matching_etablissements = null;
    public ?string $nature_juridique = null;
    public ?string $nom_complet = null;
    public ?string $nom_raison_sociale = null;
    public ?int $nombre_etablissements = null;
    public ?int $nombre_etablissements_ouverts = null;
    public ?string $section_activite_principale = null;
    public ?array $siege = null;
    public ?string $sigle = null;
    public ?string $siren = null;
    public ?string $statut_diffusion = null;
    public ?string $tranche_effectif_salarie = null;
}

/** Request payload for NearPoint#list. */
class NearPointListMatch
{
    public ?string $activite_principale = null;
    public ?string $activite_principale_naf25 = null;
    public ?string $annee_categorie_entreprise = null;
    public ?string $annee_tranche_effectif_salarie = null;
    public ?string $caractere_employeur = null;
    public ?string $categorie_entreprise = null;
    public ?array $complements = null;
    public ?string $date_creation = null;
    public ?string $date_fermeture = null;
    public ?string $date_mise_a_jour = null;
    public ?string $date_mise_a_jour_insee = null;
    public ?string $date_mise_a_jour_rne = null;
    public ?array $dirigeants = null;
    public ?string $etat_administratif = null;
    public ?array $finances = null;
    public ?array $matching_etablissements = null;
    public ?string $nature_juridique = null;
    public ?string $nom_complet = null;
    public ?string $nom_raison_sociale = null;
    public ?int $nombre_etablissements = null;
    public ?int $nombre_etablissements_ouverts = null;
    public ?string $section_activite_principale = null;
    public ?array $siege = null;
    public ?string $sigle = null;
    public ?string $siren = null;
    public ?string $statut_diffusion = null;
    public ?string $tranche_effectif_salarie = null;
}

/** Search entity data model. */
class Search
{
    public ?string $activite_principale = null;
    public ?string $activite_principale_naf25 = null;
    public ?string $annee_categorie_entreprise = null;
    public ?string $annee_tranche_effectif_salarie = null;
    public ?string $caractere_employeur = null;
    public ?string $categorie_entreprise = null;
    public ?array $complements = null;
    public ?string $date_creation = null;
    public ?string $date_fermeture = null;
    public ?string $date_mise_a_jour = null;
    public ?string $date_mise_a_jour_insee = null;
    public ?string $date_mise_a_jour_rne = null;
    public ?array $dirigeants = null;
    public ?string $etat_administratif = null;
    public ?array $finances = null;
    public ?array $matching_etablissements = null;
    public ?string $nature_juridique = null;
    public ?string $nom_complet = null;
    public ?string $nom_raison_sociale = null;
    public ?int $nombre_etablissements = null;
    public ?int $nombre_etablissements_ouverts = null;
    public ?string $section_activite_principale = null;
    public ?array $siege = null;
    public ?string $sigle = null;
    public ?string $siren = null;
    public ?string $statut_diffusion = null;
    public ?string $tranche_effectif_salarie = null;
}

/** Request payload for Search#list. */
class SearchListMatch
{
    public ?string $activite_principale = null;
    public ?string $activite_principale_naf25 = null;
    public ?string $annee_categorie_entreprise = null;
    public ?string $annee_tranche_effectif_salarie = null;
    public ?string $caractere_employeur = null;
    public ?string $categorie_entreprise = null;
    public ?array $complements = null;
    public ?string $date_creation = null;
    public ?string $date_fermeture = null;
    public ?string $date_mise_a_jour = null;
    public ?string $date_mise_a_jour_insee = null;
    public ?string $date_mise_a_jour_rne = null;
    public ?array $dirigeants = null;
    public ?string $etat_administratif = null;
    public ?array $finances = null;
    public ?array $matching_etablissements = null;
    public ?string $nature_juridique = null;
    public ?string $nom_complet = null;
    public ?string $nom_raison_sociale = null;
    public ?int $nombre_etablissements = null;
    public ?int $nombre_etablissements_ouverts = null;
    public ?string $section_activite_principale = null;
    public ?array $siege = null;
    public ?string $sigle = null;
    public ?string $siren = null;
    public ?string $statut_diffusion = null;
    public ?string $tranche_effectif_salarie = null;
}

