# CompanySearch SDK feature factory

from companysearch_sdk.feature.base_feature import CompanySearchBaseFeature
from companysearch_sdk.feature.test_feature import CompanySearchTestFeature


def _make_feature(name):
    features = {
        "base": lambda: CompanySearchBaseFeature(),
        "test": lambda: CompanySearchTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
