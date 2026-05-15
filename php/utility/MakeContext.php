<?php
declare(strict_types=1);

// CompanySearch SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class CompanySearchMakeContext
{
    public static function call(array $ctxmap, ?CompanySearchContext $basectx): CompanySearchContext
    {
        return new CompanySearchContext($ctxmap, $basectx);
    }
}
