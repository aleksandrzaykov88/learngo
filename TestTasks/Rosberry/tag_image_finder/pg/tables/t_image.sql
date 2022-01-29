-- table: public.t_image
-- drop table if exists public.t_image;

create table if not exists public.t_image
(
    id serial not null default,
    link text collate pg_catalog."default" not null,
    alt text collate pg_catalog."default",
    created_at timestamp without time zone default curent_timestamp,
    updated_at timestamp without time zone default curent_timestamp,
    constraint t_image_pkey primary key (id)
)

tablespace pg_default;

alter table if exists public.t_image owner to postgres;
grant all on table public.t_image to postgres;

comment on table public.t_image is 'Изображения';