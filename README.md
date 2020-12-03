# Seminario GoLang
## Una sencilla API REST con una temática elegida (Música)
- CRUD de música
- SQLite para persistencia
- Framework Gin Gonic
- Yaml

## Para correr esta aplicación, deben seguirse estos pasos/requisitos:
1. Tener GoLang instalado, de lo contrario, instalarlo: [golang](https://golang.org/)
2. Clonar el repositorio en el $GOPATH/src. (Crear el directorio *src* si es necesario. (En Windows, $GOPATH está ppor defecto en C:/Users/$YOURUSER/go)
3. Abrirlo con algún IDE, como por ejemplo [VS CODE](https://code.visualstudio.com/) 
4. Abrir la consola e insertar este comando:
```
go run cmd/music/music.go -config ./config/config.yaml
```
5. Si aparece una alerta de Firewall, darle a permitir.
6. ¡Listo! Se puede testear la API con [Postman](https://mongusteam.postman.co/home) o similar. En este ejemplo, la base de datos comienza vacía.


## Servicios REST:

### Leer toda la música
- Método: *GET*
- Endpoint: *localhost:8080/music*
- Cuerpo: No se requiere.

### Leer sólo una
- Método: *GET*
- Endpoint: *localhost:8080/music/:id*
- Cuerpo: No se requiere.

### Agregar una música (ej. canción)
- Método: *POST*
- Endpoint: *localhost:8080/music*
- Cuerpo: 
```
{
  "name": "I want it that way",
  "artist": "Back Street Boys",
  "year": 1999
}
```

### Eliminar una música
- Método: *DELETE*
- Endpoint: *localhost:8080/music/:id*
- Cuerpo: No se requiere.

### Editar una música
- Método: *PUT*
- Endpoint: *localhost:8080/music*
- Cuerpo: 
```
{
  "name": "La mano de Dios",
  "artist": "Rodrigo",
  "year": 2000,
  "ID": 1
}
```