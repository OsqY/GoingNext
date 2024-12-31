-- name: CreateRole :one
INSERT INTO roles (
    name,
    description,
    created_by
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: CreateUser :one
INSERT INTO users (
    email,
    username,
    password,
    role_id,
    image_url,
    created_by
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUserByID :one
SELECT u.*, r.name as role_name 
FROM users u
JOIN roles r ON u.role_id = r.id
WHERE u.id = $1 AND u.deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT u.*, r.name as role_name 
FROM users u
JOIN roles r ON u.role_id = r.id
WHERE u.email = $1 AND u.deleted_at IS NULL;

-- name: ListUsers :many
SELECT u.*, r.name as role_name 
FROM users u
JOIN roles r ON u.role_id = r.id
WHERE u.deleted_at IS NULL
ORDER BY u.created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET 
    email = COALESCE(sqlc.narg(email), email),
    username = COALESCE(sqlc.narg(username), username),
    role_id = COALESCE(sqlc.narg(role_id), role_id),
    password = COALESCE(sqlc.narg(password), password),
    image_url = COALESCE(sqlc.narg(image_url), image_url),
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteUser :exec
UPDATE users
SET 
    deleted_at = CURRENT_TIMESTAMP,
    deleted_by = $2
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListRoles :many
SELECT * FROM roles
WHERE deleted_at IS NULL
ORDER BY name;
