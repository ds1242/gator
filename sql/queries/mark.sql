-- name: MarkFeedFetched :one
UPDATE feeds
SET updated_at = $1, last_fetched_at = $2
WHERE id = $3
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT url
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
