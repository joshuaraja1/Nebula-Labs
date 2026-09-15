# Nebula Labs

Nebula Labs is an independent educational-data platform that combines a data
collection pipeline with a Go REST API. The repository is organized as a
monorepo so ingestion, parsing, validation, API schemas, and service behavior
can evolve together.

> **Project status:** active development. The imported services retain
> UTD-specific behavior while the platform is refactored toward independently
> owned configuration, infrastructure, and product features.

## What is included

| Component | Purpose | Core technologies |
| --- | --- | --- |
| [`services/api-tools`](services/api-tools) | Scrapes source data, parses and validates records, then uploads normalized data | Go, ChromeDP, structured parsing, validation |
| [`services/nebula-api`](services/nebula-api) | Serves educational data through REST endpoints and shared schemas | Go, Gin, MongoDB, Swagger |

## Architecture

```text
External data sources
        |
        v
Scrapers -> Parsers -> Validation -> Uploaders
                                      |
                                      v
                                   MongoDB
                                      |
                                      v
                              Nebula REST API
```

The separation between collection, transformation, and serving keeps the
pipeline testable and allows each stage to fail independently without silently
corrupting downstream data.

## Recent independent work

- Added package and function-level documentation for shared API utility code.
- Added a typed scholarship schema with stable JSON and BSON field mappings.
- Added serialization tests covering required keys, MongoDB `_id` behavior,
  numeric GPA values, and slice round trips.
- Preserved both upstream histories and MIT license notices during the monorepo
  import.

## Getting started

### Prerequisites

- Git
- Go 1.26 or newer for the API service
- Make on macOS/Linux, or the included `.bat` scripts on Windows
- MongoDB access when running the API against persistent data

### Clone

```bash
git clone https://github.com/joshuaraja1/Nebula-Labs.git
cd Nebula-Labs
```

### Build and test the REST API

```bash
cd services/nebula-api
make setup
go test ./...
make build
```

To run the API locally, copy `.env.template` to `.env`, configure the
required values, and start the generated `rest-api` executable.

### Build and test the data tools

```bash
cd services/api-tools
make setup
go test ./...
make build
```

Run `./api-tools` to view the available scrape, parse, validate, and upload
options. See the component [README](services/api-tools/README.md) for the full
command reference.

## Repository layout

```text
Nebula-Labs/
├── services/
│   ├── api-tools/       # collection, parsing, validation, and upload pipeline
│   └── nebula-api/      # REST service, routes, controllers, and schemas
├── CONTRIBUTING.md
├── NOTICE.md
└── README.md
```

## Contributing

Issues and pull requests are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md)
for the local workflow, testing expectations, and pull-request checklist.

## Roadmap

- Replace upstream-specific configuration with environment-driven adapters.
- Expand schema validation and contract tests.
- Add automated CI for builds, tests, and formatting.
- Improve observability across ingestion and API request paths.
- Document local sample-data workflows.

## Attribution and licensing

This repository is an independent derivative of the MIT-licensed
[UTDNebula/api-tools](https://github.com/UTDNebula/api-tools) and
[UTDNebula/nebula-api](https://github.com/UTDNebula/nebula-api) projects. It is
not the official UTD Nebula organization or its production service.

Original histories and license notices are retained in each imported component.
See [NOTICE.md](NOTICE.md) and the component license files for details.
