SELECT
id,
name,
last_name,
email,
age,
birthday,
telephone_number
FROM
cka.clients
WHERE id=$1;