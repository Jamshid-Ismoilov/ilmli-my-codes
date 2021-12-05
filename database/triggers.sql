-- --TRIGGER FUNCTIONS

-- -- trigger works after inserting into table users and encrypts
-- -- users's password


-- CREATE or REPLACE FUNCTION encrypt_users()
-- RETURNS TRIGGER
-- AS
-- $$
-- BEGIN
-- 	UPDATE users SET password = crypt(new.password::varchar, gen_salt('bf')) WHERE user_id = NEW.user_id;
-- 	return new;
-- END;
-- $$
-- LANGUAGE plpgsql;



-- CREATE TRIGGER crypt_password_trigger AFTER INSERT ON users
-- FOR EACH ROW
-- EXECUTE PROCEDURE encrypt_users();

-- select * from users where password = crypt('1234', password);



