-- table: public.t_image_tag
-- drop table if exists public.t_image_tag;

create table if not exists public.t_image_tag
(
    id serial not null default,
    name text collate pg_catalog."default" not null,
    image_id integer not null,
    constraint t_image_tag_pkey primary key (id),
    constraint t_image_tag_fkey foreign key (image_id)
        references public.t_image (id) match simple
        on update no action
        on delete cascade
        not valid
)

tablespace pg_default;

alter table if exists public.t_image_tag owner to postgres;
grant all on table public.t_image_tag to postgres;

comment on table public.t_image_tag is 'Теги изображений';