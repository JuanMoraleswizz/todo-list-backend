# Todo List API

Una API REST para gestionar tareas (todos) construida con Go usando Clean Architecture.

## Características

- Crear, leer, actualizar y eliminar tareas
- API REST con endpoints JSON
- Base de datos PostgreSQL con GORM
- Clean Architecture (Domain, UseCase, Repository, Handler)
- Migración automática de base de datos
- Soporte CORS
- Dockerizado con Docker Compose

## API Endpoints

- `GET /todos` - Obtener todas las tareas
- `POST /todos` - Crear una nueva tarea
- `GET /todos/{id}` - Obtener una tarea específica
- `PUT /todos/{id}` - Actualizar una tarea
- `DELETE /todos/{id}` - Eliminar una tarea

## Estructura de datos

```json
{
  "id": 1,
  "title": "Título de la tarea",
  "description": "Descripción de la tarea",
  "completed": false,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Estructura del proyecto

```
├── internal/
│   ├── domain/          # Entidades y interfaces
│   ├── usecase/         # Lógica de negocio
│   ├── repository/      # Acceso a datos
│   └── handler/         # Controladores HTTP
├── pkg/
│   └── database/        # Configuración de BD
└── main.go             # Punto de entrada
```

## Ejecución local

### Con Docker Compose (recomendado)
```bash
docker-compose up --build
```

### Con Go (requiere PostgreSQL local)
```bash
# Configurar variables de entorno
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=todolist

go mod tidy
go run main.go
```

## Ejemplos de uso

### Crear una tarea
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Mi primera tarea", "description": "Descripción de la tarea"}'
```

### Obtener todas las tareas
```bash
curl http://localhost:8080/todos
```

### Actualizar una tarea
```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"completed": true}'
```

### Eliminar una tarea
```bash
curl -X DELETE http://localhost:8080/todos/1
```

La aplicación estará disponible en http://localhost:8080