
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'CompanySearch',
        slug: "company-search",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     test:     {
      "options": {
        "active": false
      }
    },

  }


  options = {
    base: "https://recherche-entreprises.api.gouv.fr",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      near_point: {
      },

      search: {
      },

    }
  }


  entity = {
    "near_point": {
      "fields": [
        {
          "name": "activite_principale",
          "short": "Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "activite_principale_naf25",
          "short": "Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "annee_categorie_entreprise",
          "short": "Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "annee_tranche_effectif_salarie",
          "short": "Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "caractere_employeur",
          "short": "Caractère employeur de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "categorie_entreprise",
          "short": "Catégorie d'entreprise de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "complements",
          "type": "`$OBJECT`"
        },
        {
          "name": "date_creation",
          "short": "Date de création de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "date_fermeture",
          "short": "Date de fermeture de l'unité légale (source : base historique SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "date_mise_a_jour",
          "short": "Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "date_mise_a_jour_insee",
          "short": "Date de la dernière mise à jour des données INSEE pour cette unité légale.",
          "type": "`$STRING`"
        },
        {
          "name": "date_mise_a_jour_rne",
          "short": "Date de la dernière mise à jour des données RNCS pour cette unité légale.",
          "type": "`$STRING`"
        },
        {
          "name": "dirigeants",
          "type": "`$ARRAY`",
          "union": {
            "branches": 2,
            "count": 1,
            "depth": 1
          }
        },
        {
          "name": "etat_administratif",
          "short": "État administratif de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "finances",
          "short": "Bilans financiers par année",
          "type": "`$OBJECT`"
        },
        {
          "name": "matching_etablissements",
          "short": "Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements.",
          "type": "`$ARRAY`"
        },
        {
          "name": "nature_juridique",
          "short": "Catégorie juridique de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "nom_complet",
          "short": "Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège.",
          "type": "`$STRING`"
        },
        {
          "name": "nom_raison_sociale",
          "short": "La raison sociale pour les personnes morales (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "nombre_etablissements",
          "type": "`$INTEGER`"
        },
        {
          "name": "nombre_etablissements_ouverts",
          "type": "`$INTEGER`"
        },
        {
          "name": "section_activite_principale",
          "short": "Calculée à partir de l'activité principale.",
          "type": "`$STRING`"
        },
        {
          "name": "siege",
          "type": "`$OBJECT`"
        },
        {
          "name": "sigle",
          "short": "Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "siren",
          "short": "le numéro unique de l'entreprise",
          "type": "`$STRING`"
        },
        {
          "name": "statut_diffusion",
          "short": "Statut de diffusion de l'unité légale.",
          "type": "`$STRING`"
        },
        {
          "name": "tranche_effectif_salarie",
          "short": "Tranche d'effectif salarié de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        }
      ],
      "name": "near_point",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": "01.12Z,28.15Z",
                    "kind": "query",
                    "name": "activite_principale",
                    "orig": "activite_principale",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "siege,complements",
                    "kind": "query",
                    "name": "include",
                    "orig": "include",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "lat",
                    "orig": "lat",
                    "reqd": true,
                    "type": "`$NUMBER`"
                  },
                  {
                    "example": 10,
                    "kind": "query",
                    "name": "limite_matching_etablissement",
                    "orig": "limite_matching_etablissement",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "long",
                    "orig": "long",
                    "reqd": true,
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "minimal",
                    "orig": "minimal",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page_etablissement",
                    "orig": "page_etablissement",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 10,
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 5,
                    "kind": "query",
                    "name": "radius",
                    "orig": "radius",
                    "type": "`$NUMBER`"
                  },
                  {
                    "example": "A,J,U",
                    "kind": "query",
                    "name": "section_activite_principale",
                    "orig": "section_activite_principale",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort_by_size",
                    "orig": "sort_by_size",
                    "type": "`$BOOLEAN`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/near_point",
              "parts": [
                "near_point"
              ],
              "select": {
                "exist": [
                  "activite_principale",
                  "include",
                  "lat",
                  "limite_matching_etablissement",
                  "long",
                  "minimal",
                  "page",
                  "page_etablissement",
                  "per_page",
                  "radius",
                  "section_activite_principale",
                  "sort_by_size"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.results`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "search": {
      "fields": [
        {
          "name": "activite_principale",
          "short": "Code de l'activité principale exercée (APE) par l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "activite_principale_naf25",
          "short": "Activité principale de l'unité légale selon la nomenclature NAF 2025 (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "annee_categorie_entreprise",
          "short": "Année de validité correspondant à la catégorie d'entreprise diffusée (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "annee_tranche_effectif_salarie",
          "short": "Année de validité de la tranche d'effectif salarié de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "caractere_employeur",
          "short": "Caractère employeur de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "categorie_entreprise",
          "short": "Catégorie d'entreprise de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "complements",
          "type": "`$OBJECT`"
        },
        {
          "name": "date_creation",
          "short": "Date de création de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "date_fermeture",
          "short": "Date de fermeture de l'unité légale (source : base historique SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "date_mise_a_jour",
          "short": "Date de la dernière modification d'une variable de niveau unité légale, qu'elle soit historisée ou non (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "date_mise_a_jour_insee",
          "short": "Date de la dernière mise à jour des données INSEE pour cette unité légale.",
          "type": "`$STRING`"
        },
        {
          "name": "date_mise_a_jour_rne",
          "short": "Date de la dernière mise à jour des données RNCS pour cette unité légale.",
          "type": "`$STRING`"
        },
        {
          "name": "dirigeants",
          "type": "`$ARRAY`",
          "union": {
            "branches": 2,
            "count": 1,
            "depth": 1
          }
        },
        {
          "name": "etat_administratif",
          "short": "État administratif de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "finances",
          "short": "Bilans financiers par année",
          "type": "`$OBJECT`"
        },
        {
          "name": "matching_etablissements",
          "short": "Liste des établissements ayant contribué au résultat de la recherche : ceux qui ont « matché » la recherche textuelle ou un filtre sur les établissements.",
          "type": "`$ARRAY`"
        },
        {
          "name": "nature_juridique",
          "short": "Catégorie juridique de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "nom_complet",
          "short": "Champ construit depuis les champs de dénomination : dénomination de l'unité légale | Nom et prénom | Nom inconnu (dénomination usuelle : Construite en priorité à partir de la dénomination usuelle de l'établissement siège.",
          "type": "`$STRING`"
        },
        {
          "name": "nom_raison_sociale",
          "short": "La raison sociale pour les personnes morales (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "nombre_etablissements",
          "type": "`$INTEGER`"
        },
        {
          "name": "nombre_etablissements_ouverts",
          "type": "`$INTEGER`"
        },
        {
          "name": "section_activite_principale",
          "short": "Calculée à partir de l'activité principale.",
          "type": "`$STRING`"
        },
        {
          "name": "siege",
          "type": "`$OBJECT`"
        },
        {
          "name": "sigle",
          "short": "Forme réduite de la raison sociale ou de la dénomination d'une personne morale ou d'un organisme public (source : base SIRENE).",
          "type": "`$STRING`"
        },
        {
          "name": "siren",
          "short": "le numéro unique de l'entreprise",
          "type": "`$STRING`"
        },
        {
          "name": "statut_diffusion",
          "short": "Statut de diffusion de l'unité légale.",
          "type": "`$STRING`"
        },
        {
          "name": "tranche_effectif_salarie",
          "short": "Tranche d'effectif salarié de l'unité légale (source : base SIRENE).",
          "type": "`$STRING`"
        }
      ],
      "name": "search",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": "01.12Z,28.15Z",
                    "kind": "query",
                    "name": "activite_principale",
                    "orig": "activite_principale",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 100000,
                    "kind": "query",
                    "name": "ca_max",
                    "orig": "ca_max",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 100000,
                    "kind": "query",
                    "name": "ca_min",
                    "orig": "ca_min",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": "PME",
                    "kind": "query",
                    "name": "categorie_entreprise",
                    "orig": "categorie_entreprise",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "75C",
                    "kind": "query",
                    "name": "code_collectivite_territoriale",
                    "orig": "code_collectivite_territoriale",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "01247,01111",
                    "kind": "query",
                    "name": "code_commune",
                    "orig": "code_commune",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "38540,38189",
                    "kind": "query",
                    "name": "code_postal",
                    "orig": "code_postal",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "convention_collective_renseignee",
                    "orig": "convention_collective_renseignee",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": "1990-01-01",
                    "kind": "query",
                    "name": "date_naissance_personne_max",
                    "orig": "date_naissance_personne_max",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "1960-01-01",
                    "kind": "query",
                    "name": "date_naissance_personne_min",
                    "orig": "date_naissance_personne_min",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "02,89",
                    "kind": "query",
                    "name": "departement",
                    "orig": "departement",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "egapro_renseignee",
                    "orig": "egapro_renseignee",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": "200058519,248100737",
                    "kind": "query",
                    "name": "epci",
                    "orig": "epci",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "est_achats_responsable",
                    "orig": "est_achats_responsable",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_alim_confiance",
                    "orig": "est_alim_confiance",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_association",
                    "orig": "est_association",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_bio",
                    "orig": "est_bio",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_collectivite_territoriale",
                    "orig": "est_collectivite_territoriale",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_entrepreneur_individuel",
                    "orig": "est_entrepreneur_individuel",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_entrepreneur_spectacle",
                    "orig": "est_entrepreneur_spectacle",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_ess",
                    "orig": "est_ess",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_finess",
                    "orig": "est_finess",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_l100_3",
                    "orig": "est_l100_3",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_organisme_formation",
                    "orig": "est_organisme_formation",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_patrimoine_vivant",
                    "orig": "est_patrimoine_vivant",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_qualiopi",
                    "orig": "est_qualiopi",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_rge",
                    "orig": "est_rge",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_service_public",
                    "orig": "est_service_public",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_siae",
                    "orig": "est_siae",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_societe_mission",
                    "orig": "est_societe_mission",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "est_uai",
                    "orig": "est_uai",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "etat_administratif",
                    "orig": "etat_administratif",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "1090",
                    "kind": "query",
                    "name": "id_convention_collective",
                    "orig": "id_convention_collective",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "010003853",
                    "kind": "query",
                    "name": "id_finess",
                    "orig": "id_finess",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "8611M10D109",
                    "kind": "query",
                    "name": "id_rge",
                    "orig": "id_rge",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "0022004T",
                    "kind": "query",
                    "name": "id_uai",
                    "orig": "id_uai",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "siege,complements",
                    "kind": "query",
                    "name": "include",
                    "orig": "include",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 10,
                    "kind": "query",
                    "name": "limite_matching_etablissement",
                    "orig": "limite_matching_etablissement",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "minimal",
                    "orig": "minimal",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": "7344,6544",
                    "kind": "query",
                    "name": "nature_juridique",
                    "orig": "nature_juridique",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "Dupont",
                    "kind": "query",
                    "name": "nom_personne",
                    "orig": "nom_personne",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page_etablissement",
                    "orig": "page_etablissement",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 10,
                    "kind": "query",
                    "name": "per_page",
                    "orig": "per_page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": "Monsieur",
                    "kind": "query",
                    "name": "prenoms_personne",
                    "orig": "prenoms_personne",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "11,76",
                    "kind": "query",
                    "name": "region",
                    "orig": "region",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 100000,
                    "kind": "query",
                    "name": "resultat_net_max",
                    "orig": "resultat_net_max",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 100000,
                    "kind": "query",
                    "name": "resultat_net_min",
                    "orig": "resultat_net_min",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": "A,J,U",
                    "kind": "query",
                    "name": "section_activite_principale",
                    "orig": "section_activite_principale",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort_by_size",
                    "orig": "sort_by_size",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": "NN,00,01",
                    "kind": "query",
                    "name": "tranche_effectif_salarie",
                    "orig": "tranche_effectif_salarie",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "type_personne",
                    "orig": "type_personne",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/search",
              "parts": [
                "search"
              ],
              "select": {
                "exist": [
                  "activite_principale",
                  "ca_max",
                  "ca_min",
                  "categorie_entreprise",
                  "code_collectivite_territoriale",
                  "code_commune",
                  "code_postal",
                  "convention_collective_renseignee",
                  "date_naissance_personne_max",
                  "date_naissance_personne_min",
                  "departement",
                  "egapro_renseignee",
                  "epci",
                  "est_achats_responsable",
                  "est_alim_confiance",
                  "est_association",
                  "est_bio",
                  "est_collectivite_territoriale",
                  "est_entrepreneur_individuel",
                  "est_entrepreneur_spectacle",
                  "est_ess",
                  "est_finess",
                  "est_l100_3",
                  "est_organisme_formation",
                  "est_patrimoine_vivant",
                  "est_qualiopi",
                  "est_rge",
                  "est_service_public",
                  "est_siae",
                  "est_societe_mission",
                  "est_uai",
                  "etat_administratif",
                  "id_convention_collective",
                  "id_finess",
                  "id_rge",
                  "id_uai",
                  "include",
                  "limite_matching_etablissement",
                  "minimal",
                  "nature_juridique",
                  "nom_personne",
                  "page",
                  "page_etablissement",
                  "per_page",
                  "prenoms_personne",
                  "q",
                  "region",
                  "resultat_net_max",
                  "resultat_net_min",
                  "section_activite_principale",
                  "sort_by_size",
                  "tranche_effectif_salarie",
                  "type_personne"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.results`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config
}

