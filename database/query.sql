--QUERIES


create table users (
	user_id serial not null primary key,
	fullname varchar(128),
	highlight varchar(32),
	mobile_number varchar(16),
	bio varchar(64),
	learners bigint default null,
	password varchar(256),
	email varchar(128) unique
);

create table teachers (
	user_id int references users (user_id),
	teacher_id int references users (user_id)
	created_at timestamp default current_timestamp
);

-- GET USER DETAILS
select
	fullname,
	highlight,
	mobile_number,
	bio,
	learners 
from users
where user_id = 1
;

select
	l.teacher_id
from users as u
join learning as l on l.user_id = u.user_id
where u.user_id = 1
;


--GET TEACHERS BY ID
select
	u.fullname,
	u.user_id
from (
	select
		l.teacher_id as user_id
	from users as u
	join learning as l on l.user_id = u.user_id
	where u.user_id = 1
) as t
join users as u on u.user_id = t.user_id
;

select
	p.post_id
	p.name,
	p.picture
from users as u
join posts as p on p.user_id = u.user_id
where u.user_id = 1
;

select
	post_id,
	name,
	picture 
from posts
where category_id = 1
;

select
	post_id,
	name,
	content,
	source,
	picture
from posts
where post_id = 1
;

select
	category_id,
	name,
	description,
	picture
from categories
where category_id = 1
;
