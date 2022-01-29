-- function: public.fn_image_tag_get(text)

-- drop function public.fn_image_tag_get(text);

create or replace function public.fn_image_tag_get(
    arg_tag text
)
returns table (
    id integer,
    link text,
    alt text,
    created_at timestamp,
    updated_at timestamp,
    tags json
) as
$body$
declare
begin
    if arg_tag = '' then
        return query
            select
                i.id,
                i.link,
                i.alt,
                i.created_at,
                i.updated_at,
                i.tags
            from v_image i;
    else
        return query
            select
                i.id,
                i.link,
                i.alt,
                i.created_at,
                i.updated_at,
                i.tags
            from v_image i
            join t_image_tag it on it.image_id = i.id
            where it.name = arg_tag;
    end if;
end;
$body$
    language plpgsql volatile
    cost 100;

alter function public.fn_image_tag_get(text) owner to postgres;
grant execute on function public.fn_image_tag_get(text) to postgres;
comment on function public.fn_image_tag_get(text) is 'Изображения. Поиск по тегу';
