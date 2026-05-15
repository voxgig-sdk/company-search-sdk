# CompanySearch SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module CompanySearchFeatures
  def self.make_feature(name)
    case name
    when "base"
      CompanySearchBaseFeature.new
    when "test"
      CompanySearchTestFeature.new
    else
      CompanySearchBaseFeature.new
    end
  end
end
