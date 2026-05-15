<?php
declare(strict_types=1);

// CompanySearch SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

CompanySearchUtility::setRegistrar(function (CompanySearchUtility $u): void {
    $u->clean = [CompanySearchClean::class, 'call'];
    $u->done = [CompanySearchDone::class, 'call'];
    $u->make_error = [CompanySearchMakeError::class, 'call'];
    $u->feature_add = [CompanySearchFeatureAdd::class, 'call'];
    $u->feature_hook = [CompanySearchFeatureHook::class, 'call'];
    $u->feature_init = [CompanySearchFeatureInit::class, 'call'];
    $u->fetcher = [CompanySearchFetcher::class, 'call'];
    $u->make_fetch_def = [CompanySearchMakeFetchDef::class, 'call'];
    $u->make_context = [CompanySearchMakeContext::class, 'call'];
    $u->make_options = [CompanySearchMakeOptions::class, 'call'];
    $u->make_request = [CompanySearchMakeRequest::class, 'call'];
    $u->make_response = [CompanySearchMakeResponse::class, 'call'];
    $u->make_result = [CompanySearchMakeResult::class, 'call'];
    $u->make_point = [CompanySearchMakePoint::class, 'call'];
    $u->make_spec = [CompanySearchMakeSpec::class, 'call'];
    $u->make_url = [CompanySearchMakeUrl::class, 'call'];
    $u->param = [CompanySearchParam::class, 'call'];
    $u->prepare_auth = [CompanySearchPrepareAuth::class, 'call'];
    $u->prepare_body = [CompanySearchPrepareBody::class, 'call'];
    $u->prepare_headers = [CompanySearchPrepareHeaders::class, 'call'];
    $u->prepare_method = [CompanySearchPrepareMethod::class, 'call'];
    $u->prepare_params = [CompanySearchPrepareParams::class, 'call'];
    $u->prepare_path = [CompanySearchPreparePath::class, 'call'];
    $u->prepare_query = [CompanySearchPrepareQuery::class, 'call'];
    $u->result_basic = [CompanySearchResultBasic::class, 'call'];
    $u->result_body = [CompanySearchResultBody::class, 'call'];
    $u->result_headers = [CompanySearchResultHeaders::class, 'call'];
    $u->transform_request = [CompanySearchTransformRequest::class, 'call'];
    $u->transform_response = [CompanySearchTransformResponse::class, 'call'];
});
