~~~sql
select coalesce(nullif(client_no,''),'0') from table

select cast(substr(coalesce(nullif(client,''),'0000'),7,8)as integer) from table
~~~

