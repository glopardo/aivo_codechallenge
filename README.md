# Aivo code challenge
## Descripción
Esta API cuenta con un endpoint para obtener la dicografía de un artista a partir del nombre.
Para la autenticación cuenta con los secrets personales de mi cuenta de spotify (hardcodeados, por cuestión de tiempo).

### Endpoint
GET `"/api/v1/albums/:q"`

### Descripción
La API obtiene el access token (a partir de los secrets hardcodeados) al instanciar el controller.
Con ese access token, y el parámetro `:q` que contiene el nombre del artista (si son más de una palabra, separa por guión bajo/medio), obtiene el ID del artista.
El ID es utilizado para voler a la API de Spotify y traer la discrografía por ID de artista.
Se mapean los objetos recibidos a una estructura propia y se devuelven en el body de la respuesta del request.

_Observación_: Al buscar por nombre de artista, la API de spotify devuelve una lista de artistas, se toma el primer valor de la lista (que se asume es el que más similitud tiene con el nombre ingresado).

### Tests
Por cuestión de tiempo y practicidad sólo se agregaron tests a las instanciaciones de controller/service/repository. Agregar los mocks para testear todos los métodos de obtención de datos iba a requerir más tiempo del que contaba.
