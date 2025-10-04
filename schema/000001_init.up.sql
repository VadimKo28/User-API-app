create table users 
(
  id   uuid DEFAULT gen_random_uuid(),
  name varchar(255) not null,
  age  integer not null
);