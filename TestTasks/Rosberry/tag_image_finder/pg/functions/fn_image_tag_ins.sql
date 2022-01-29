-- function: public.fn_image_tag_ins(int, text)

-- drop function public.fn_image_tag_ins(int, text);

create or replace function public.fn_image_tag_ins(
    arg_image_id int,
    arg_tag text
)
returns integer as
$body$
declare
    v_id int;
begin
    perform z_exception_if(arg_image_id is null, 'Не передан код изображения');
    perform z_exception_if(arg_tag      is null, 'Не передано значение тега');

    perform z_exception_if(
        not exists(
            select 1
            from t_image i
            where i.id = arg_image_id
        ),
        'Изображеия с кодом %L не существует',
        arg_image_id::text
    );

    perform z_exception_if(
        exists(
            select 1
            from t_image_tag it
            where it.image_id = arg_image_id and
            it.name = arg_tag
        ),
        'Тег %L уже существует',
         arg_tag::text
    );

    insert into t_image_tag as it (
        name,
        image_id
    ) values (
        arg_tag,
        arg_image_id
    ) returning it.id into v_id;

    return v_id;
end;
$body$
    language plpgsql volatile
    cost 100;

alter function public.fn_image_tag_ins(int, text) owner to postgres;
grant execute on function public.fn_image_tag_ins(int, text) to postgres;
comment on function public.fn_image_tag_ins(int, text) is 'Теги изображений. Создание';
