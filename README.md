# Kong Takehome Exercise

Prerequisites:

- Docker
- npm
- Go 1.17

I use macOS so I've included a Makefile for easy commands. If you are on Windows or another OS that doesn't play nice with make, you'll need to run commands supported by your system.

```
$ make help
build                          compile the binary
clean                          clean workspace and remove build/test artifacts
docker                         run docker containers
help                           commands help text
install                        install npm dependencies
run                            launch the server locally
test-integration               run integration tests
test-unit                      run unit tests
```

## API Design

### Endpoints

All results are sorted by Service ID ascending, then by Version ID ascending if applicable. IDs are just the integer primary keys for the corresponding database rows. In a real API, internal database identifiers would not be leaked outside the database like this, instead sharing something like a UUID.

- `GET /services`
    - Returns list of services in the user's organization
    - Supports pagination via `limit` and `offset` query parameters
    - Supports search via `q` query parameter
    - Includes the number of versions for each service
    - I didn't implement filtering or sorting, but those would be easy additions
- `GET /services/:serviceID`
    - Returns service by ID, or 404 if the service does not exist in the user's organization
    - Includes the first 5 versions for the service
- `GET /services/:serviceID/versions`
    - Returns a paginated list of versions for the given service
    - Supports pagination of the versions list via `limit` and `offset` query parameters
- `GET /services/:serviceID/versions/:versionID`
    - Returns version by ID, or 404 if the service/version does not exist in the user's organization

### Authentication

All endpoints expect the `X-User-Id` header to be a valid user ID. An actual API would not do this, but it's a simple mechanism to demonstrate the point for this exercise. Each User and Service belongs to an organization, and users only receive results for services owned by the same organization.

- Organization 1
    - User IDs `1` and `2`
    - 60 services
- Organization 2
    - User ID `3`
    - 5 services
- Organization 3
    - User ID `4`
    - No services

## Directory Structure

- `kong_takehome` - API implementation
  - `main.go` - main entrypoint, sets up DB connection and runs HTTP server
  - `middleware.go` - HTTP server middleware
  - `data.go` - database interface and SQL queries
  - `routes.go` - API endpoint handlers
- `sql` - contains SQL migration files
- `test` - contains Mocha-based integration tests

## Design Choices

After playing around with full-text search in SQLite3 using `fts5` and `spellfix1`, I chose PostgreSQL as the database because it provides nice out-of-the-box full-text search and trigram matching (and because I heard it mentioned as being used at Kong by one the interviewers). I've never implemented full-text search myself, so I played around with both regular `@@/tsvector/tsquery` search and the `pg_trgm` extension and settled on the `pg_trgm ILIKE` approach because it gave better results without putting in the effort required for actually robust search.

I haven't ever used the `database/sql` standard library before, so I took the opportunity to implement the data layer using that. It was OK, has some pros and cons.

For the server, just a simple `net/http` server using https://github.com/julienschmidt/httprouter to make accessing named path parameters easier.

## Testing

### Unit Tests

I wrote just enough unit tests to validate the first endpoint I implemented and to demonstrate some approaches to unit testing in Go that I follow. I did not implement an exhaustive suite of tests.

To run unit tests:

```
$ make test-unit
```

### Integration Tests

I wrote basic integration tests for all the endpoints, but they are not exhaustive. These use mocha, chai, and supertest. The `test-integration` make task rebuilds the Docker Compose stack and runs the tests against the Docker container.

Note: these can be flaky if they run before the database is up and populated with test data, so I popped a `sleep 1` in there, but they might still flake sometimes.

To run integration tests:

```
$ make test-integration
```
