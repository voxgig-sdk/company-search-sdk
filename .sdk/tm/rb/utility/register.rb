# CompanySearch SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

CompanySearchUtility.registrar = ->(u) {
  u.clean = CompanySearchUtilities::Clean
  u.done = CompanySearchUtilities::Done
  u.make_error = CompanySearchUtilities::MakeError
  u.feature_add = CompanySearchUtilities::FeatureAdd
  u.feature_hook = CompanySearchUtilities::FeatureHook
  u.feature_init = CompanySearchUtilities::FeatureInit
  u.fetcher = CompanySearchUtilities::Fetcher
  u.make_fetch_def = CompanySearchUtilities::MakeFetchDef
  u.make_context = CompanySearchUtilities::MakeContext
  u.make_options = CompanySearchUtilities::MakeOptions
  u.make_request = CompanySearchUtilities::MakeRequest
  u.make_response = CompanySearchUtilities::MakeResponse
  u.make_result = CompanySearchUtilities::MakeResult
  u.make_point = CompanySearchUtilities::MakePoint
  u.make_spec = CompanySearchUtilities::MakeSpec
  u.make_url = CompanySearchUtilities::MakeUrl
  u.param = CompanySearchUtilities::Param
  u.prepare_auth = CompanySearchUtilities::PrepareAuth
  u.prepare_body = CompanySearchUtilities::PrepareBody
  u.prepare_headers = CompanySearchUtilities::PrepareHeaders
  u.prepare_method = CompanySearchUtilities::PrepareMethod
  u.prepare_params = CompanySearchUtilities::PrepareParams
  u.prepare_path = CompanySearchUtilities::PreparePath
  u.prepare_query = CompanySearchUtilities::PrepareQuery
  u.result_basic = CompanySearchUtilities::ResultBasic
  u.result_body = CompanySearchUtilities::ResultBody
  u.result_headers = CompanySearchUtilities::ResultHeaders
  u.transform_request = CompanySearchUtilities::TransformRequest
  u.transform_response = CompanySearchUtilities::TransformResponse
}
