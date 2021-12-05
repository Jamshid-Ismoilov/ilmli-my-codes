--MOCK DATA


insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa1', 'highlight', 'numberr', 'bio bio', 34, 'picture1', 'addasd', 'email mail');
insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa2', 'highlight', 'numberr', 'bio bio', 34, 'picture2', 'addasd', 'email mail1');
insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa3', 'highlight', 'numberr', 'bio bio', 34, 'picture3', 'addasd', 'email mail2');
insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa4', 'highlight', 'numberr', 'bio bio', 34, 'picture4', 'addasd', 'email mai11l');
insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa5', 'highlight', 'numberr', 'bio bio', 34, 'picture5', 'addasd', 'email ma1i11l');
insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa6', 'highlight', 'numberr', 'bio bio', 34, 'picture6', 'addasd', 'email ma111il');
insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa7', 'highlight', 'numberr', 'bio bio', 34, 'picture7', 'addasd', 'email ma11111il');
insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa8', 'highlight', 'numberr', 'bio bio', 34, 'picture8', 'addasd', 'email 1mail');
insert into users (fullname, highlight, mobile_number, bio, learners, picture, password, email)
	values ('aaaa9', 'highlight', 'numberr', 'bio bio', 34, 'picture9', 'addasd', 'email ma1il');


insert into categories (name, description, picture, user_id)
	values ('category 1', 'description', 'asnd.img', 1);
insert into categories (name, description, picture, user_id)
	values ('category 2', 'description', 'asnd.img', 1);
insert into categories (name, description, picture, user_id)
	values ('category 3', 'description', 'asnd.img', 2);

insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 1', 'content', 'google.com', 'ilm.png', 1, 1);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 2', 'content', 'google.com', 'ilm.png', 1, 1);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 3', 'content', 'google.com', 'ilm.png', 1, 1);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 4', 'content', 'google.com', 'ilm.png', 1, 1);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 5', 'content', 'google.com', 'ilm.png', 1, 1);


insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 1', 'content', 'google.com', 'ilm.png', 1, 2);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 2', 'content', 'google.com', 'ilm.png', 1, 2);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 3', 'content', 'google.com', 'ilm.png', 1, 2);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 4', 'content', 'google.com', 'ilm.png', 1, 2);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 5', 'content', 'google.com', 'ilm.png', 1, 2);


insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 1', 'content', 'google.com', 'ilm.png', 2, 3);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 2', 'content', 'google.com', 'ilm.png', 2, 3);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 3', 'content', 'google.com', 'ilm.png', 2, 3);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 4', 'content', 'google.com', 'ilm.png', 2, 3);
insert into posts (name, content, source, picture, user_id, category_id)
	values ('post 5', 'content', 'google.com', 'ilm.png', 2, 3);

insert into learning(user_id, teacher_id)
	values (1,2);
insert into learning(user_id, teacher_id)
	values (1,3);
insert into learning(user_id, teacher_id)
	values (1,4);
insert into learning(user_id, teacher_id)
	values (1,5);


insert into learning(user_id, teacher_id)
	values (2,1);
insert into learning(user_id, teacher_id)
	values (2,2);
insert into learning(user_id, teacher_id)
	values (2,4);
insert into learning(user_id, teacher_id)
	values (2,5);