/*
Since general full-text search is Hard™, I'm just doing a simple attempt at basic flexible
exact matching using trigram indexes and ILIKE '%query%'. This doesn't get us tolerance for
misspelled words, matching on similar words or lexemes, etc., but it is more flexible than
full-text search using @@/tsquery/tsvector and gives us prefix, suffix, and substring matching.
*/

CREATE EXTENSION pg_trgm;

CREATE INDEX CONCURRENTLY services_title_trgm_idx
ON services
USING GIN (title gin_trgm_ops);

CREATE INDEX CONCURRENTLY services_summary_trgm_idx
ON services
USING GIN (summary gin_trgm_ops);
