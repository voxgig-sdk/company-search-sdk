package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewNearPointEntityFunc func(client *CompanySearchSDK, entopts map[string]any) CompanySearchEntity

var NewSearchEntityFunc func(client *CompanySearchSDK, entopts map[string]any) CompanySearchEntity

