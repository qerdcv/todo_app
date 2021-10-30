-- name: ListTodos :many
SELECT * FROM todos ORDER BY id;

-- name: CreateTodo :exec
INSERT INTO todos ( text ) VALUES ($1);

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;
