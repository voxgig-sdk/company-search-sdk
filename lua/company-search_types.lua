-- Typed models for the CompanySearch SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class NearPoint
---@field activite_principale? string
---@field activite_principale_naf25? string
---@field annee_categorie_entreprise? string
---@field annee_tranche_effectif_salarie? string
---@field caractere_employeur? string
---@field categorie_entreprise? string
---@field complement? table
---@field date_creation? string
---@field date_fermeture? string
---@field date_mise_a_jour? string
---@field date_mise_a_jour_insee? string
---@field date_mise_a_jour_rne? string
---@field dirigeant? table
---@field etat_administratif? string
---@field finance? table
---@field matching_etablissement? table
---@field nature_juridique? string
---@field nom_complet? string
---@field nom_raison_sociale? string
---@field nombre_etablissement? number
---@field nombre_etablissements_ouvert? number
---@field section_activite_principale? string
---@field siege? table
---@field sigle? string
---@field siren? string
---@field statut_diffusion? string
---@field tranche_effectif_salarie? string

---@class NearPointListMatch

---@class Search
---@field activite_principale? string
---@field activite_principale_naf25? string
---@field annee_categorie_entreprise? string
---@field annee_tranche_effectif_salarie? string
---@field caractere_employeur? string
---@field categorie_entreprise? string
---@field complement? table
---@field date_creation? string
---@field date_fermeture? string
---@field date_mise_a_jour? string
---@field date_mise_a_jour_insee? string
---@field date_mise_a_jour_rne? string
---@field dirigeant? table
---@field etat_administratif? string
---@field finance? table
---@field matching_etablissement? table
---@field nature_juridique? string
---@field nom_complet? string
---@field nom_raison_sociale? string
---@field nombre_etablissement? number
---@field nombre_etablissements_ouvert? number
---@field section_activite_principale? string
---@field siege? table
---@field sigle? string
---@field siren? string
---@field statut_diffusion? string
---@field tranche_effectif_salarie? string

---@class SearchListMatch

local M = {}

return M
