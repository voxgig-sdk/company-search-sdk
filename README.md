# CompanySearch SDK

API Recherche d’entreprises client, generated from the OpenAPI spec.

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

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

## Quickstart

### TypeScript

```ts
import { CompanySearchSDK } from 'company-search'

const client = new CompanySearchSDK({
  apikey: process.env.COMPANY-SEARCH_APIKEY,
})

// List all nearpoints
const nearpoints = await client.NearPoint().list()
console.log(nearpoints.data)
```

See the [TypeScript README](ts/README.md) for the full guide.

## Surfaces

| Surface | Path |
| --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | `go-cli/` |
| **MCP server** | `go-mcp/` |

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
| **NearPoint** |  | `/near_point` |
| **Search** |  | `/search` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
import os
from companysearch_sdk import CompanySearchSDK

client = CompanySearchSDK({
    "apikey": os.environ.get("COMPANY-SEARCH_APIKEY"),
})

# List all nearpoints
nearpoints, err = client.NearPoint().list()
print(nearpoints)
```

### PHP

```php
<?php
require_once 'companysearch_sdk.php';

$client = new CompanySearchSDK([
    "apikey" => getenv("COMPANY-SEARCH_APIKEY"),
]);

// List all nearpoints
[$nearpoints, $err] = $client->NearPoint()->list();
print_r($nearpoints);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/company-search-sdk/go"

client := sdk.NewCompanySearchSDK(map[string]any{
    "apikey": os.Getenv("COMPANY-SEARCH_APIKEY"),
})

// List all nearpoints
nearpoints, err := client.NearPoint(nil).List(nil, nil)
fmt.Println(nearpoints)
```

### Ruby

```ruby
require_relative "CompanySearch_sdk"

client = CompanySearchSDK.new({
  "apikey" => ENV["COMPANY-SEARCH_APIKEY"],
})

# List all nearpoints
nearpoints, err = client.NearPoint().list
puts nearpoints
```

### Lua

```lua
local sdk = require("company-search_sdk")

local client = sdk.new({
  apikey = os.getenv("COMPANY-SEARCH_APIKEY"),
})

-- List all nearpoints
local nearpoints, err = client:NearPoint():list()
print(nearpoints)
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
client = CompanySearchSDK.test()
result, err = client.NearPoint().load({"id": "test01"})
```

### PHP

```php
$client = CompanySearchSDK::test();
[$result, $err] = $client->NearPoint()->load(["id" => "test01"]);
```

### Golang

```go
client := sdk.Test()
result, err := client.NearPoint(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = CompanySearchSDK.test
result, err = client.NearPoint().load({ "id" => "test01" })
```

### Lua

```lua
local client = sdk.test()
local result, err = client:NearPoint():load({ id = "test01" })
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

---

Generated from the API Recherche d’entreprises OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
