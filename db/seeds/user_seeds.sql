INSERT INTO
  users (id, name, email)
VALUES
	(1, 'yamamoto', 'mizuki-y@syati.info')
ON DUPLICATE KEY UPDATE
  name=VALUES(name),
  email=VALUES(email);