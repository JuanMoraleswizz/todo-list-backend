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
- Pipeline CI/CD con Jenkins

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
├── Jenkinsfile         # Pipeline CI/CD
└── main.go             # Punto de entrada
```

## Tests

La aplicación incluye tests unitarios para todas las capas:

### Ejecutar tests
```bash
go test ./...
```

### Ejecutar tests con cobertura
```bash
go test ./... -cover
```

### Ejecutar tests en modo verbose
```bash
go test ./... -v
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

## Pipeline CI/CD con Jenkins

El proyecto incluye un `Jenkinsfile` que define un pipeline completo de CI/CD con las siguientes etapas:

1. **Checkout** - Descarga el código fuente
2. **Setup Go** - Configura el entorno Go y descarga dependencias
3. **Test** - Ejecuta las pruebas unitarias
4. **Build** - Compila la aplicación
5. **Docker Build** - Construye las imágenes Docker
6. **Deploy** - Despliega la aplicación (solo en rama main)

### Configuración de Jenkins

1. Crear un nuevo job tipo "Pipeline"
2. Configurar el repositorio Git
3. Seleccionar "Pipeline script from SCM"
4. Asegurar que Jenkins tenga acceso a Docker
5. Configurar las credenciales necesarias

El pipeline se ejecutará automáticamente en cada push al repositorio.