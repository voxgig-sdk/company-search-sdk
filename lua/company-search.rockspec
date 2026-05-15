package = "voxgig-sdk-company-search"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/company-search-sdk.git"
}
description = {
  summary = "CompanySearch SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["company-search_sdk"] = "company-search_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
