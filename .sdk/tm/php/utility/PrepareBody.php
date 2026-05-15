<?php
declare(strict_types=1);

// CompanySearch SDK utility: prepare_body

class CompanySearchPrepareBody
{
    public static function call(CompanySearchContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
