# client-kpi-api


  - GET /clients: Obtiene la lista de todos los clientes desde la base de datos en SQLite o en memoria.
   - GET /clients/{id}: Obtiene los detalles de un cliente específico.
   - POST /clients: Crea un nuevo cliente. El cuerpo del request debe incluir el nombre, apellido, email, número de teléfono, edad y fecha de nacimiento.
   - PUT /clients/{id}: Actualiza un cliente existente.
   - DELETE /clients/{id}: Elimina un cliente específico.
   - GET /clients/kpi: Devuelve KPI de los clientes, tales como:
       Promedio de edad.
       Desviación estándar de la edad.

Validaciones de Entrada:
   - Los campos name, last_name, email, edad y fecha de nacimiento son obligatorios para crear un cliente.
   - Validar que el email tenga un formato correcto.
   - El número de teléfono debe ser numérico y tener un mínimo de 7 dígitos.
   - La fecha de nacimiento debe ser válida y coherente con la edad provista.
   - Manejo de errores claros y consistentes para casos como cliente no encontrado, datos inválidos, etc.

   Crear un endpoint GET /clients/kpi que calcule y devuelva los siguientes indicadores clave:
      - Promedio de edad: El cálculo del promedio de edad entre todos los clientes registrados.
  - Desviación estándar de edad: Calcular la variación de las edades respecto al promedio.
  - Los KPI deben calcularse en tiempo real basados en los clientes almacenados.