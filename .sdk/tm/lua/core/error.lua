-- CompanySearch SDK error

local CompanySearchError = {}
CompanySearchError.__index = CompanySearchError


function CompanySearchError.new(code, msg, ctx)
  local self = setmetatable({}, CompanySearchError)
  self.is_sdk_error = true
  self.sdk = "CompanySearch"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function CompanySearchError:error()
  return self.msg
end


function CompanySearchError:__tostring()
  return self.msg
end


return CompanySearchError
