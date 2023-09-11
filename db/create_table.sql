create table if not exists pressure (
id serial primary key,
station text,
pressure decimal,
date date,
hour integer
)