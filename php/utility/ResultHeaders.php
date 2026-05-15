<?php
declare(strict_types=1);

// CompanySearch SDK utility: result_headers

class CompanySearchResultHeaders
{
    public static function call(CompanySearchContext $ctx): ?CompanySearchResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
