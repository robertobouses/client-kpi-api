SELECT EXISTS (
    SELECT 1 
    FROM cka.clients 
    WHERE id = $1
);
