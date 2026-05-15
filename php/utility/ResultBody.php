<?php
declare(strict_types=1);

// CompanySearch SDK utility: result_body

class CompanySearchResultBody
{
    public static function call(CompanySearchContext $ctx): ?CompanySearchResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
