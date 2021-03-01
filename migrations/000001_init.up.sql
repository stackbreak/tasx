create table person (
  id integer primary key generated always as identity,
  name varchar(200) not null,
  username varchar(200) not null unique,
  password_hash char(60) not null
);

create table task (
  id integer primary key generated always as identity,
  title varchar(255) not null,
  description varchar(1000),
  is_done boolean not null default false
);

create table task_list(
  id integer primary key generated always as identity,
  title varchar(255) not null,
  description varchar(500),
  person_id integer references person(id) on delete cascade
);

create table task_to_list(
  id integer primary key generated always as identity,
  task_id integer references task(id) on delete cascade,
  task_list_id integer references task_list(id) on delete cascade
);
