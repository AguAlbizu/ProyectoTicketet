# TicketApp

Sistema de gestión y venta de entradas para eventos, con soporte para sorteos y dos roles de usuario (Cliente y Administrador).

---

## Descripción

TicketApp permite a los clientes explorar eventos, comprar entradas, cancelarlas o transferirlas a otros usuarios. Cada evento puede tener un sorteo asociado donde los clientes adquieren chances y el administrador ejecuta el sorteo, notificando a los participantes por email.

Los administradores gestionan el catálogo de eventos y acceden a reportes de ventas.

---

## Tecnologías

| Capa       | Tecnología                          |
|------------|-------------------------------------|
| Backend    | Go 1.22 + Gin + GORM                |
| Frontend   | React 18 + Vite + React Router      |
| Base de datos | MySQL 8                          |
| Auth       | JWT (roles: cliente / admin)        |
| Contenedores | Docker + Docker Compose           |

---

## Requisitos Previos

- [Docker](https://www.docker.com/) y Docker Compose v2+
- (Desarrollo local) Go 1.22+ y Node 20+
- Cuenta SMTP para el envío de emails (sorteos)

---

## Instalación

### Con Docker (recomendado)

```bash
# 1. Clonar el repositorio
git clone <repo-url>
cd ticketapp

# 2. Copiar y completar las variables de entorno
cp backend/.env.example backend/.env
# Editar backend/.env con tus valores reales

# 3. Levantar todos los servicios
docker compose up --build
```

La API queda disponible en `http://localhost:8080` y el frontend en `http://localhost:3000`.

### Desarrollo local (sin Docker)

```bash
# Backend
cd backend
cp .env.example .env
go mod download
go run main.go

# Frontend (en otra terminal)
cd frontend
cp .env.example .env
npm install
npm run dev
```

---

## Cómo correr los tests

```bash
cd backend
go test ./tests/... -v
```

---

## Diagrama de Base de Datos

<!-- TODO: insertar imagen del diagrama ER -->
![Diagrama BD](docs/db-diagram.png)

---

## Decisiones de Diseño

<!-- TODO: documentar decisiones arquitectónicas relevantes -->
- Arquitectura MVC en el backend con separación en capas: domain → dao → service → controller
- JWT stateless con rol embebido en el claim
- Soft delete en usuarios y entradas (GORM DeletedAt)
