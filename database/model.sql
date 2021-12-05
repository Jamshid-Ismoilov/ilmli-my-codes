--MODEL 

create table users (
	user_id serial not null primary key,
	fullname varchar(128),
	highlight varchar(32),
	mobile_number varchar(16),
	bio varchar(64),
	learners bigint default null,
	picture varchar(64),
	password varchar(256),
	email varchar(128) unique
);

create table learning (
	learning_id serial not null primary key,
	user_id int references users (user_id),
	teacher_id int references users (user_id),
	created_at timestamp default current_timestamp
);

create table categories (
	category_id serial not null primary key,
	name varchar(128),
	description varchar(256),
	picture varchar(256)
);

create table posts (
	post_id serial not null primary key,
	name varchar(64),
	content text,
	source varchar(32),
	picture varchar(256),
	user_id int references users (user_id),
	category_id int references categories(category_id)
);

create table trendings (
	trending_id serial not null primary key,
	post_id int references categories (category_id),
	count bigint default null,
	created_at timestamp default current_timestamp
);


