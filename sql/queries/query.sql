-- name: ListCategories :many
select * from categories;

-- name: ListCategorie :one
select * from categories where id = ?;

-- name: CreateCategorie :exec
insert into categories (id, name, description) values (?,?,?);

-- name: UpdateCategorie :exec
update categories set name= ?, description=?  where id = ?;

-- name: DeleteCategorie :exec
delete from categories where id = ?;
