# CompanySearch SDK feature factory

from companysearch_sdk.feature.base_feature import CompanySearchBaseFeature
from companysearch_sdk.feature.test_feature import CompanySearchTestFeature


_FEATURES = {
    "base": lambda: CompanySearchBaseFeature(),
    "test": lambda: CompanySearchTestFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
