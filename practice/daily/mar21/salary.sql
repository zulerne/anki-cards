select salary from (
    select salary, dense_rank() over (order by salary desc) as rn
    from employees
)t where rn = 2

select max(salary) from employees
where salary < (select max(salary) from employees);
