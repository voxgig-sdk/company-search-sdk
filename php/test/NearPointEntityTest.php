<?php
declare(strict_types=1);

// NearPoint entity test

require_once __DIR__ . '/../companysearch_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class NearPointEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = CompanySearchSDK::test(null, null);
        $ent = $testsdk->NearPoint(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = near_point_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "near_point." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set COMPANYSEARCH_TEST_NEAR_POINT_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $near_point_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.near_point")));
        $near_point_ref01_data = null;
        if (count($near_point_ref01_data_raw) > 0) {
            $near_point_ref01_data = Helpers::to_map($near_point_ref01_data_raw[0][1]);
        }

        // LIST
        $near_point_ref01_ent = $client->NearPoint(null);
        $near_point_ref01_match = [];

        [$near_point_ref01_list_result, $err] = $near_point_ref01_ent->list($near_point_ref01_match, null);
        $this->assertNull($err);
        $this->assertIsArray($near_point_ref01_list_result);

    }
}

function near_point_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/near_point/NearPointTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = CompanySearchSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["near_point01", "near_point02", "near_point03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("COMPANYSEARCH_TEST_NEAR_POINT_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "COMPANYSEARCH_TEST_NEAR_POINT_ENTID" => $idmap,
        "COMPANYSEARCH_TEST_LIVE" => "FALSE",
        "COMPANYSEARCH_TEST_EXPLAIN" => "FALSE",
        "COMPANYSEARCH_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["COMPANYSEARCH_TEST_NEAR_POINT_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["COMPANYSEARCH_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["COMPANYSEARCH_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new CompanySearchSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["COMPANYSEARCH_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["COMPANYSEARCH_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
