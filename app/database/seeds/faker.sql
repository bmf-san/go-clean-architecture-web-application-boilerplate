INSERT INTO users
  (name)
VALUES('foo');
INSERT INTO users
  (name)
VALUES('bar');

INSERT INTO posts
  (user_id, body)
VALUES(1, 'Hello World');
INSERT INTO posts
  (user_id, body)
VALUES
  (2, 'Hi');