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

-- name: CreateCourses :exec
insert into courses (id, category_id , name, description, price) values (?,?,?,?,?);

-- name: ListCourses :many
select c.*, ca.name as category_name 
from courses c join categories ca on c.category_id = ca.id;
