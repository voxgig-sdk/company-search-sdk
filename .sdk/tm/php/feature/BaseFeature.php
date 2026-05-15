<?php
declare(strict_types=1);

// CompanySearch SDK base feature

class CompanySearchBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(CompanySearchContext $ctx, array $options): void {}
    public function PostConstruct(CompanySearchContext $ctx): void {}
    public function PostConstructEntity(CompanySearchContext $ctx): void {}
    public function SetData(CompanySearchContext $ctx): void {}
    public function GetData(CompanySearchContext $ctx): void {}
    public function GetMatch(CompanySearchContext $ctx): void {}
    public function SetMatch(CompanySearchContext $ctx): void {}
    public function PrePoint(CompanySearchContext $ctx): void {}
    public function PreSpec(CompanySearchContext $ctx): void {}
    public function PreRequest(CompanySearchContext $ctx): void {}
    public function PreResponse(CompanySearchContext $ctx): void {}
    public function PreResult(CompanySearchContext $ctx): void {}
    public function PreDone(CompanySearchContext $ctx): void {}
    public function PreUnexpected(CompanySearchContext $ctx): void {}
}
