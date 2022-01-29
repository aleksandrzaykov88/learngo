-- function: public.z_exception_if(boolean, text, text[])

-- drop function public.z_exception_if(boolean, text, text[]);

create or replace function public.z_exception_if(
    arg_condition boolean,
    arg_message text,
    variadic arg_params text[] default array[]::text[]
)
returns void as
$body$
begin
    if arg_condition then
        raise exception 'USER_ERROR %', format(arg_message, variadic arg_params);
    end if;
end;
$body$
  language plpgsql immutable
  cost 100;

alter function public.z_exception_if(boolean, text, text[]) owner to postgres;
grant execute on function public.z_exception_if(boolean, text, text[]) to postgres;
comment on function public.z_exception_if(boolean, text, text[]) is 'Выбрасывает исключение по условию';
