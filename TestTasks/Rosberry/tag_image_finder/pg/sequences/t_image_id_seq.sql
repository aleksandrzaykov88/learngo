-- sequence: public.t_image_id_seq
-- drop sequence if exists public.t_image_id_seq;

create sequence if not exists public.t_image_id_seq
    increment 1
    start 1
    minvalue 1
    maxvalue 2147483647
    cache 1
    owned by t_image.id;

alter sequence public.t_image_id_seq owner to postgres;
grant all on table public.t_image_id_seq to postgres;