CREATE TABLE IF NOT EXISTS tasks (
  id varchar(255) primary key,
  title varchar(255),
  description text,
  status varchar(50),
  expires_at timestamp with time zone not null,
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null
);