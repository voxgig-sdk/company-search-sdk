# CompanySearch SDK utility: make_context

from projectname_sdk.core.context import CompanySearchContext


def make_context_util(ctxmap, basectx):
    return CompanySearchContext(ctxmap, basectx)
