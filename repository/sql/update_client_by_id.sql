UPDATE cka.clients
SET 
    name = COALESCE($1, name),
    last_name = COALESCE($2, last_name),
    email = COALESCE($3, email),
    age = COALESCE($4, age),
    birthday = COALESCE($5, birthday),
    telephone_number=COALESCE($6, telephone_number)
WHERE id = $7;