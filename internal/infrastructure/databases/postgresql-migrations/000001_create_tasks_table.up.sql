CREATE TABLE IF NOT EXISTS tasks (
  id bigserial primary key,
  title varchar(255),
  description text,
  status varchar(50),
  expires_at timestamp,
  created_at timestamp,
  updated_at timestamp
);