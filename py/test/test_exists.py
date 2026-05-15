# ProjectName SDK exists test

import pytest
from companysearch_sdk import CompanySearchSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = CompanySearchSDK.test(None, None)
        assert testsdk is not None
