Yes! `pgx` has a dedicated package called `pgtype` that provides nullable-aware types for every PostgreSQL type. These are far better than raw Go types or `any`/`interface{}`.

## How `pgtype` Works

Each `pgtype` type has two fields:
- A typed **value** field (e.g., `Int4`, `Text`, `Bool`)
- A **`Valid bool`** field — `true` if the value is present, `false` if NULL

---

## Common `pgtype` Mappings

| PostgreSQL Type | pgtype Type | Value Field |
|---|---|---|
| `serial` / `integer` | `pgtype.Int4` | `.Int32` |
| `bigserial` / `bigint` | `pgtype.Int8` | `.Int64` |
| `text` / `varchar` | `pgtype.Text` | `.String` |
| `boolean` | `pgtype.Bool` | `.Bool` |
| `numeric` / `decimal` | `pgtype.Numeric` | `.Int`, `.Exp` |
| `float4` | `pgtype.Float4` | `.Float32` |
| `float8` | `pgtype.Float8` | `.Float64` |
| `timestamptz` | `pgtype.Timestamptz` | `.Time` |
| `timestamp` | `pgtype.Timestamp` | `.Time` |
| `date` | `pgtype.Date` | `.Time` |
| `uuid` | `pgtype.UUID` | `.Bytes` |
| `jsonb` / `json` | `pgtype.Text` or raw `[]byte` | — |
| `bytea` | `pgtype.Bytea` | `.Bytes` |

---

## Example Usage

```go
package main

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
)

type User struct {
    ID        pgtype.Int4
    Username  pgtype.Text
    Email     pgtype.Text
    Score     pgtype.Float8   // nullable float
    CreatedAt pgtype.Timestamptz
    DeletedAt pgtype.Timestamptz // nullable timestamp
}

func getUser(conn *pgx.Conn, id int) (*User, error) {
    query := `
        SELECT id, username, email, score, created_at, deleted_at
        FROM users
        WHERE id = $1
    `

    var u User
    err := conn.QueryRow(context.Background(), query, id).Scan(
        &u.ID,
        &u.Username,
        &u.Email,
        &u.Score,
        &u.CreatedAt,
        &u.DeletedAt,
    )
    if err != nil {
        // pgx.ErrNoRows if not found — row-level null vs no row at all
        return nil, err
    }

    return &u, nil
}

func printUser(u *User) {
    // Check Valid before using the value
    if u.DeletedAt.Valid {
        fmt.Println("Deleted at:", u.DeletedAt.Time)
    } else {
        fmt.Println("User is not deleted")
    }

    if u.Score.Valid {
        fmt.Printf("Score: %.2f\n", u.Score.Float64)
    } else {
        fmt.Println("Score: N/A")
    }
}
```

---

## Handling "Row Not Found" vs NULL columns

These are **two different concerns** — `pgtype` handles one of them:

```go
var u User
err := conn.QueryRow(ctx, query, id).Scan(&u.ID /* ... */)

// 1. No row found at all → pgx.ErrNoRows
if errors.Is(err, pgx.ErrNoRows) {
    fmt.Println("user not found")
    return nil, nil
}

// 2. Row found but some columns are NULL → handled by pgtype.Valid
if err != nil {
    return nil, err
}

// Safe to use u here; check .Valid per nullable field
```

---

## Key Takeaway

Using `pgtype` gives you:
- **Full type safety** — no `any` or type assertions needed
- **Null awareness** via `.Valid` — no pointer gymnastics (`*int32`, `*string`, etc.)
- **Seamless scanning** — pgx knows how to scan directly into `pgtype` structs
- **Cleaner structs** — your model clearly communicates which fields can be NULL