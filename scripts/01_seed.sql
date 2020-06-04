INSERT INTO users (id, display_name, updated_at, created_at) VALUES
 (NULL, '', 0, 0),
 (NULL, '', 0, 0),
 (NULL, '', 0, 0),
 (NULL, '', 0, 0),
 (NULL, '', 0, 0),
 (NULL, '', 0, 0),
 (NULL, '', 0, 0),
 (NULL, '', 0, 0),
 (NULL, '', 0, 0),
 (NULL, '', 0, 0);


INSERT INTO `users` (`display_name`, `created_at`, `updated_at`) SELECT T1.`display_name`, T1.`created_at`, T1.`updated_at` from `users` as T1 CROSS JOIN `users` as T2 ON T2.id IN (1, 2, 3, 4, 5, 6, 7, 8, 9);
-- 100
INSERT INTO `users` (`display_name`, `created_at`, `updated_at`) SELECT T1.`display_name`, T1.`created_at`, T1.`updated_at` from `users` as T1 CROSS JOIN `users` as T2 ON T2.id IN (1, 2, 3, 4, 5, 6, 7, 8, 9);
-- 1,000
INSERT INTO `users` (`display_name`, `created_at`, `updated_at`) SELECT T1.`display_name`, T1.`created_at`, T1.`updated_at` from `users` as T1 CROSS JOIN `users` as T2 ON T2.id IN (1, 2, 3, 4, 5, 6, 7, 8, 9);
-- 10,000
INSERT INTO `users` (`display_name`, `created_at`, `updated_at`) SELECT T1.`display_name`, T1.`created_at`, T1.`updated_at` from `users` as T1 CROSS JOIN `users` as T2 ON T2.id IN (1, 2, 3, 4, 5, 6, 7, 8, 9);
-- 100,000

UPDATE users SET
updated_at = CEIL(RAND() * 2678400 + 1561939200),
created_at = CEIL(RAND() * 2678400 + 1561939200);

INSERT INTO user_login_names (login_name, password_hash, salt, user_id, updated_at, created_at) SELECT CONCAT(SUBSTRING(MD5(RAND()), 1, 50), id) , SUBSTRING(MD5(RAND()), 1, 10), SUBSTRING(MD5(RAND()), 1, 10), id, CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM users;


INSERT INTO channels (name, owner, updated_at, created_at) SELECT SUBSTRING(MD5(RAND()), 1, 10), id, CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM users WHERE id < 1000;

INSERT INTO channels (name, owner, updated_at, created_at) SELECT SUBSTRING(MD5(RAND()), 1, 10), owner, CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM channels;
INSERT INTO channels (name, owner, updated_at, created_at) SELECT SUBSTRING(MD5(RAND()), 1, 10), owner, CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM channels;

INSERT INTO messages (channel_id, user_id, type, body, updated_at, created_at) SELECT CEIL(RAND() * 2500), id, SUBSTRING(MD5(RAND()), 1, 1), SUBSTRING(MD5(RAND()), 1, 50), CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM users;
INSERT INTO messages (channel_id, user_id, type, body, updated_at, created_at) SELECT CEIL(RAND() * 2500), id, SUBSTRING(MD5(RAND()), 1, 1), SUBSTRING(MD5(RAND()), 1, 50), CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM users;
INSERT INTO messages (channel_id, user_id, type, body, updated_at, created_at) SELECT CEIL(RAND() * 2500), id, SUBSTRING(MD5(RAND()), 1, 1), SUBSTRING(MD5(RAND()), 1, 50), CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM users;
INSERT INTO messages (channel_id, user_id, type, body, updated_at, created_at) SELECT CEIL(RAND() * 2500), id, SUBSTRING(MD5(RAND()), 1, 1), SUBSTRING(MD5(RAND()), 1, 50), CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM users;
INSERT INTO messages (channel_id, user_id, type, body, updated_at, created_at) SELECT CEIL(RAND() * 2500), id, SUBSTRING(MD5(RAND()), 1, 1), SUBSTRING(MD5(RAND()), 1, 50), CEIL(RAND() * 2678400 + 1561939200), CEIL(RAND() * 2678400 + 1561939200)  FROM users;
