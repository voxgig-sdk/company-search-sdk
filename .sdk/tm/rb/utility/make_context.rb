# CompanySearch SDK utility: make_context
require_relative '../core/context'
module CompanySearchUtilities
  MakeContext = ->(ctxmap, basectx) {
    CompanySearchContext.new(ctxmap, basectx)
  }
end
