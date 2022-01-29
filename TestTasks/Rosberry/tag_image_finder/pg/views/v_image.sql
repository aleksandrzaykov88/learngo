-- view: public.v_image
-- drop view public.v_image;

create or replace view public.v_image as
     select
        i.id,
        i.link,
        i.alt,
        i.created_at,
        i.updated_at,
        coalesce((
            select json_agg(itags)
            from (
                select ti.name
                from t_image_tag ti
                where ti.image_id = i.id
                ) itags
            ), '[{}]') as tags
    from t_image i
    order by i.updated_at desc;

alter table public.v_image owner to postgres;
grant all on table public.v_image to postgres;

comment on view public.v_image is 'Изображения';
