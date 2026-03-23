select email, count(*)
from users
group by email
having count(*) > 1
