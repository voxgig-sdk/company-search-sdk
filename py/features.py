# CompanySearch SDK feature factory

from feature.base_feature import CompanySearchBaseFeature
from feature.test_feature import CompanySearchTestFeature


def _make_feature(name):
    features = {
        "base": lambda: CompanySearchBaseFeature(),
        "test": lambda: CompanySearchTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
