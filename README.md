# CompanySearch SDK

Search French companies, associations, and public services by name, address, or leaders

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About API Recherche d’entreprises

The [API Recherche d'entreprises](https://recherche-entreprises.api.gouv.fr) is a public service operated by the French government ([api.gouv.fr](https://api.gouv.fr)) for searching the registry of French companies, associations, and public services. It is built on top of public datasets maintained by [INSEE](https://www.insee.fr) (the national statistics institute) and [INPI](https://www.inpi.fr) (the industrial property institute).

What you get from the API:

- Full-text search by company name, address, executives, or elected officials
- Geographic proximity search by latitude/longitude and radius
- Filtering by NAF/APE activity code, postal code, employee headcount bracket, and entity type
- Core identifiers and metadata such as company name (`dénomination`), `SIREN`, `SIRET`, and NAF activity code

The API is fully open: no API key, no authentication, and CORS is enabled. The published rate limit is 7 requests per second per IP. Data on predecessor/successor establishments, non-diffusible enterprises, and rejected RCS registrations is not exposed; full SIRENE database access requires a separate channel.

## Try it

**TypeScript**
```bash
npm install company-search
```

**Python**
```bash
pip install company-search-sdk
```

**PHP**
```bash
composer require voxgig/company-search-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/company-search-sdk/go
```

**Ruby**
```bash
gem install company-search-sdk
```

**Lua**
```bash
luarocks install company-search-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { CompanySearchSDK } from 'company-search'

const client = new CompanySearchSDK({})

// List all nearpoints
const nearpoints = await client.NearPoint().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o company-search-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "company-search": {
      "command": "/abs/path/to/company-search-mcp"
    }
  }
}
```

## Entities

The API exposes 2 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **NearPoint** | Geographic proximity search returning entities near a coordinate via `GET /near_point` with `lat`, `long`, `radius`, `page`, and `per_page` parameters. | `/near_point` |
| **Search** | Full-text search over French companies, associations, and public services via `GET /search` with parameters like `q`, `page`, `per_page`, and filters for NAF code, postal code, and employee count. | `/search` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from companysearch_sdk import CompanySearchSDK

client = CompanySearchSDK({})

# List all nearpoints
nearpoints, err = client.NearPoint(None).list(None, None)
```

### PHP

```php
<?php
require_once 'companysearch_sdk.php';

$client = new CompanySearchSDK([]);

// List all nearpoints
[$nearpoints, $err] = $client->NearPoint(null)->list(null, null);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/company-search-sdk/go"

client := sdk.NewCompanySearchSDK(map[string]any{})

// List all nearpoints
nearpoints, err := client.NearPoint(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "CompanySearch_sdk"

client = CompanySearchSDK.new({})

# List all nearpoints
nearpoints, err = client.NearPoint(nil).list(nil, nil)
```

### Lua

```lua
local sdk = require("company-search_sdk")

local client = sdk.new({})

-- List all nearpoints
local nearpoints, err = client:NearPoint(nil):list(nil, nil)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = CompanySearchSDK.test()
const result = await client.NearPoint().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = CompanySearchSDK.test(None, None)
result, err = client.NearPoint(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = CompanySearchSDK::test(null, null);
[$result, $err] = $client->NearPoint(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.NearPoint(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = CompanySearchSDK.test(nil, nil)
result, err = client.NearPoint(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:NearPoint(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the API Recherche d’entreprises

- Upstream: [https://recherche-entreprises.api.gouv.fr](https://recherche-entreprises.api.gouv.fr)
- API docs: [https://recherche-entreprises.api.gouv.fr/docs/](https://recherche-entreprises.api.gouv.fr/docs/)

- The SDK is distributed under the MIT License.
- The underlying API and its data are provided by the French government under Open Licence 2.0.
- No attribution is required for general use, but crediting the data sources (INSEE, INPI) is encouraged.
- The administration reserves the right to throttle access during server overload.

---

Generated from the API Recherche d’entreprises OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
