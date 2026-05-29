# TicketApp — Entrega Parcial

Sistema de gestión y venta de entradas para eventos. Esta entrega parcial cubre el flujo completo del usuario **Cliente**: explorar eventos, comprar entradas, cancelarlas y transferirlas a otros usuarios.

---

## Tecnologías Utilizadas

| Capa          | Tecnología                              |
|---------------|-----------------------------------------|
| Backend       | Go 1.22 + Gin                           |
| ORM           | GORM                                    |
| Base de datos | MySQL 8                                 |
| Autenticación | JWT (`github.com/golang-jwt/jwt/v5`)    |
| Frontend      | React 18 + Vite + React Router v6       |
| HTTP Client   | Axios                                   |

---

## Requisitos Previos

- **Go** 1.22 o superior → https://go.dev/dl/
- **Node.js** 20 o superior → https://nodejs.org/
- **MySQL** 8 corriendo localmente (o vía Docker)

Verificar instalaciones:

```bash
go version
node --version
mysql --version
```

---

## Instalación y Uso

### 1. Clonar el repositorio

```bash
git clone <url-del-repositorio>
cd ticketapp
```

### 2. Configurar el backend

```bash
cd backend
cp .env.example .env
```

Editar `backend/.env` con los datos reales de conexión a MySQL y el secreto JWT:

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=tu_password
DB_NAME=ticketapp
JWT_SECRET=un_secreto_seguro
JWT_EXPIRATION_HOURS=24
```

Crear la base de datos en MySQL:

```sql
CREATE DATABASE ticketapp CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. Correr el backend

```bash
cd backend
go run main.go
```

El servidor queda disponible en `http://localhost:8080`.  
Las tablas se crean automáticamente via `AutoMigrate` al iniciar.

### 4. Configurar y correr el frontend

```bash
cd frontend
cp .env.example .env
npm install
npm run dev
```

El frontend queda disponible en `http://localhost:5173`.

### 5. Correr los tests

```bash
cd backend
go test ./tests/... -v -cover
```

---

## Endpoints disponibles (Entrega Parcial)

| Método   | Ruta                        | Auth | Descripción                        |
|----------|-----------------------------|------|------------------------------------|
| POST     | /api/auth/register          | No   | Registrar nuevo usuario            |
| POST     | /api/auth/login             | No   | Iniciar sesión, retorna JWT        |
| GET      | /api/events                 | No   | Listar eventos (filtro por categoría) |
| GET      | /api/events/:id             | No   | Detalle de un evento               |
| POST     | /api/tickets                | JWT  | Comprar entrada para un evento     |
| GET      | /api/tickets/my-tickets     | JWT  | Ver mis entradas                   |
| DELETE   | /api/tickets/:id            | JWT  | Cancelar una entrada               |
| PUT      | /api/tickets/:id/transfer   | JWT  | Transferir entrada a otro usuario  |

---

## Diagrama de Base de Datos

<!-- TODO: insertar imagen del diagrama ER -->
![Diagrama BD](docs/db-diagram.png)

---

## Decisiones de Diseño

<!-- TODO: documentar decisiones arquitectónicas relevantes -->

Algunas decisiones tomadas en esta entrega:

- **Arquitectura en capas**: `domain → dao → service → controller`. Cada capa depende solo de la inmediata inferior, facilitando los tests unitarios.
- **JWT stateless**: el token incluye `user_id` y `role` en el payload. En esta entrega parcial el middleware solo verifica firma y expiración; la validación de rol queda para la entrega final.
- **Hash de contraseñas**: SHA-256 vía `crypto/sha256` de la librería estándar de Go.
- **AutoMigrate**: GORM crea y actualiza las tablas automáticamente al iniciar el servidor.
