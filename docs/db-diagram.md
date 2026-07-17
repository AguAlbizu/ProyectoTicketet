# Diagrama de Base de Datos — TicketApp

```mermaid
erDiagram
    USERS {
        uint id_users PK
        varchar nombre
        varchar email UK
        varchar password
        varchar rol
        datetime created_at
        datetime updated_at
    }

    EVENTS {
        uint id_events PK
        varchar titulo
        text descripcion
        date fecha
        varchar hora
        int capacidad
        int cupo_disponible
        varchar categoria
        varchar direccion
        varchar imagen_url
        int precio
        varchar estado
        datetime created_at
        datetime updated_at
    }

    TICKETS {
        uint id_tickets PK
        uint id_users FK
        uint id_events FK
        varchar estado
        varchar origen
        datetime fecha_compra
        datetime created_at
        datetime updated_at
    }

    SORTEOS {
        uint id_sorteo PK
        uint id_events FK
        varchar nombre
        int valor_chance
        varchar estado
        uint id_ganador FK
        datetime fecha_realizado
        datetime created_at
        datetime updated_at
    }

    CHANCES {
        uint id_chance PK
        uint id_sorteo FK
        uint id_users FK
        datetime fecha_compra
        datetime created_at
    }

    NOTIFICATIONS {
        uint id_notification PK
        uint id_users FK
        varchar tipo
        varchar titulo
        text mensaje
        boolean leida
        uint id_sorteo FK
        datetime created_at
    }

    USERS       ||--o{ TICKETS       : "compra"
    EVENTS      ||--o{ TICKETS       : "genera"
    EVENTS      ||--o{ SORTEOS       : "tiene"
    USERS       |o--o{ SORTEOS       : "gana (opcional)"
    SORTEOS     ||--o{ CHANCES       : "recibe"
    USERS       ||--o{ CHANCES       : "compra"
    USERS       ||--o{ NOTIFICATIONS : "recibe"
    SORTEOS     |o--o{ NOTIFICATIONS : "genera (opcional)"
```

## Claves foráneas implementadas

| Tabla origen | Columna | Referencia | ON DELETE | ON UPDATE |
|---|---|---|---|---|
| `tickets` | `id_users` | `users.id_users` | RESTRICT | CASCADE |
| `tickets` | `id_events` | `events.id_events` | RESTRICT | CASCADE |
| `sorteos` | `id_events` | `events.id_events` | CASCADE | CASCADE |
| `sorteos` | `id_ganador` | `users.id_users` | SET NULL | CASCADE |
| `chances` | `id_sorteo` | `sorteos.id_sorteo` | CASCADE | CASCADE |
| `chances` | `id_users` | `users.id_users` | RESTRICT | CASCADE |
| `notifications` | `id_users` | `users.id_users` | CASCADE | CASCADE |
| `notifications` | `id_sorteo` | `sorteos.id_sorteo` | CASCADE | CASCADE |

`USERS`, `EVENTS` y `TICKETS` son las tres entidades principales requeridas por la consigna.
`SORTEOS`, `CHANCES` y `NOTIFICATIONS` son las entidades agregadas para el Bonus Track
(sorteo por evento y notificaciones in-app).
