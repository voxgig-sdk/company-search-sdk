# CompanySearch SDK utility: feature_add
module CompanySearchUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
