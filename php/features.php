<?php
declare(strict_types=1);

// CompanySearch SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class CompanySearchFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new CompanySearchBaseFeature();
            case "test":
                return new CompanySearchTestFeature();
            default:
                return new CompanySearchBaseFeature();
        }
    }
}
