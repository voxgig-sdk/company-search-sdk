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
    public ?string $include = null;
    public float $lat;
    public ?int $limite_matching_etablissement = null;
    public float $long;
    public ?bool $minimal = null;
    public ?int $page = null;
    public ?int $page_etablissement = null;
    public ?int $per_page = null;
    public ?float $radius = null;
    public ?string $section_activite_principale = null;
    public ?bool $sort_by_size = null;
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
    public ?int $ca_max = null;
    public ?int $ca_min = null;
    public ?string $categorie_entreprise = null;
    public ?string $code_collectivite_territoriale = null;
    public ?string $code_commune = null;
    public ?string $code_postal = null;
    public ?bool $convention_collective_renseignee = null;
    public ?string $date_naissance_personne_max = null;
    public ?string $date_naissance_personne_min = null;
    public ?string $departement = null;
    public ?bool $egapro_renseignee = null;
    public ?string $epci = null;
    public ?bool $est_achats_responsable = null;
    public ?bool $est_alim_confiance = null;
    public ?bool $est_association = null;
    public ?bool $est_bio = null;
    public ?bool $est_collectivite_territoriale = null;
    public ?bool $est_entrepreneur_individuel = null;
    public ?bool $est_entrepreneur_spectacle = null;
    public ?bool $est_ess = null;
    public ?bool $est_finess = null;
    public ?bool $est_l100_3 = null;
    public ?bool $est_organisme_formation = null;
    public ?bool $est_patrimoine_vivant = null;
    public ?bool $est_qualiopi = null;
    public ?bool $est_rge = null;
    public ?bool $est_service_public = null;
    public ?bool $est_siae = null;
    public ?bool $est_societe_mission = null;
    public ?bool $est_uai = null;
    public ?string $etat_administratif = null;
    public ?string $id_convention_collective = null;
    public ?string $id_finess = null;
    public ?string $id_rge = null;
    public ?string $id_uai = null;
    public ?string $include = null;
    public ?int $limite_matching_etablissement = null;
    public ?bool $minimal = null;
    public ?string $nature_juridique = null;
    public ?string $nom_personne = null;
    public ?int $page = null;
    public ?int $page_etablissement = null;
    public ?int $per_page = null;
    public ?string $prenoms_personne = null;
    public ?string $q = null;
    public ?string $region = null;
    public ?int $resultat_net_max = null;
    public ?int $resultat_net_min = null;
    public ?string $section_activite_principale = null;
    public ?bool $sort_by_size = null;
    public ?string $tranche_effectif_salarie = null;
    public ?string $type_personne = null;
}

