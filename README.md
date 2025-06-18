# user-api
🧠 User API - Desarrollador Backend (Go + Gin + PostgreSQL)
# 🧠 User API - Desarrollador Backend (Go + Gin + PostgreSQL)

Este proyecto es una API RESTful para la gestión de usuarios, desarrollada en Go utilizando el framework **Gin**, el ORM **GORM** y **PostgreSQL** como base de datos.

---

## 🚀 Tecnologías utilizadas

- Go 1.22+
- Gin Gonic
- GORM
- PostgreSQL 13
- Docker (opcional, si se desea contenerizar)

---

## 📂 Estructura del proyecto

```bash
user-api/
├── internal/
│   ├── config/         # Configuración DB (database.go)
│   ├── controllers/    # Handlers HTTP (user.go)
│   └── models/         # Modelo de usuario (user.go)
├── main.go             # Punto de entrada
├── go.mod / go.sum     # Dependencias
└── README.md
```

---

## ⚙️ Requisitos

- Go 1.22+
- PostgreSQL corriendo localmente o en Docker
  - Usuario: `postgres`
  - Contraseña: `secret`
  - Base de datos: `users_db`

---

## 🛠️ Instalación

1. Clona el repositorio:

```bash
git clone https://github.com/englergz/user-api.git
cd user-api
```

2. Instala dependencias:

```bash
go mod tidy
```

3. Asegúrate de tener PostgreSQL corriendo:

```bash
docker run --name postgres -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13
```

Crea la base de datos `users_db`:

```bash
docker exec -it postgres psql -U postgres
CREATE DATABASE users_db;
\q
```

4. Corre la aplicación:

```bash
go run main.go
```

La API estará disponible en:  
`http://localhost:3000`

---

## 📡 Endpoints disponibles

### ✅ Crear usuario

```http
POST /users
Content-Type: application/json
```

**Body JSON**:
```json
{
  "name": "Engler Prado",
  "email": "engler@ejemplo.com"
}
```

### ✅ Obtener todos los usuarios

```http
GET /users
```

### ✅ Obtener un usuario por ID

```http
GET /users/{id}
```

### ✅ Actualizar usuario

```http
PUT /users/{id}
Content-Type: application/json
```

**Body JSON**:
```json
{
  "name": "Nuevo Nombre",
  "email": "nuevo@email.com"
}
```

### ✅ Eliminar usuario

```http
DELETE /users/{id}
```

---

## 🧪 Ejemplos con `curl`

```bash
# Crear
curl -X POST http://localhost:3000/users \
-H "Content-Type: application/json" \
-d '{"name":"Engler","email":"engler@ejemplo.com"}'

# Listar
curl http://localhost:3000/users

# Ver uno
curl http://localhost:3000/users/1

# Actualizar
curl -X PUT http://localhost:3000/users/1 \
-H "Content-Type: application/json" \
-d '{"name":"Nuevo Nombre","email":"nuevo@ejemplo.com"}'

# Eliminar
curl -X DELETE http://localhost:3000/users/1
```

---

## 🧠 Aprendizajes

Consulta el archivo [`aprendi.md`](aprendi.md) donde se explican los desafíos, aprendizajes y posibles mejoras del proyecto.

---

## ✅ Comandos útiles

```bash
# Formatear código
go fmt ./...

# Verificar dependencias
go mod tidy

# Correr aplicación
go run main.go
```

---

## ✨ Autor

**Engler González Prado**  
[GitHub: englergz](https://github.com/englergz)
