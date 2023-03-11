-- Drop table users
DROP TABLE IF EXISTS users;

-- Create table users
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name varchar(200) DEFAULT NULL
);

-- Insert data into table users
INSERT INTO users (name)
VALUES
  ('Ari'),
  ('Yanto');
