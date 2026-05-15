<?php
declare(strict_types=1);

// CompanySearch SDK utility: feature_add

class CompanySearchFeatureAdd
{
    public static function call(CompanySearchContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
