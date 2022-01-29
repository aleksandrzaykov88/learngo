-- function: public.fn_image_ins(text, text)

-- drop function public.fn_image_ins(text, text);

create or replace function public.fn_image_ins(
    arg_link text,
    arg_alt text
)
returns integer as
$body$
declare
    v_id int;
begin
    perform z_exception_if(arg_link is null, 'Не передано значение ссылки');

    insert into t_image as i (
        link,
        alt
    ) values (
        arg_link,
        arg_alt
    ) returning i.id into v_id;

    return v_id;
end;
$body$
    language plpgsql volatile
    cost 100;

alter function public.fn_image_ins(text, text) owner to postgres;
grant execute on function public.fn_image_ins(text, text) to postgres;
comment on function public.fn_image_ins(text, text) is 'Изображения. Создание';
