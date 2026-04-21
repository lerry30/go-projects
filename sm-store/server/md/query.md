In `pgx`, scanning via `QueryRow().Scan()` is **positional** — it maps by the **order** of columns in your `SELECT`, not by name. So you just need to ensure your `Scan()` arguments match the column order.

However, if you want **name-based mapping** (like `db` struct tags in `sqlx`), pgx v5 has `pgx.RowToStructByName` using the `db` struct tag.

---

## Option 1: Positional Scan (standard pgx)
The columns in `SELECT` must match the order of fields in `Scan()`:
```go
// Column order: id, username, email, score, created_at, deleted_at
query := `SELECT id, username, email, score, created_at, deleted_at FROM users WHERE id = $1`

err := conn.QueryRow(ctx, query, id).Scan(
    &u.ID,         // id
    &u.Username,   // username
    &u.Email,      // email
    &u.Score,      // score
    &u.CreatedAt,  // created_at
    &u.DeletedAt,  // deleted_at
)
```
Simple but fragile — **adding/reordering columns breaks it silently**.

---

## Option 2: Name-based Mapping via `db` struct tag (pgx v5)

pgx v5 has built-in `pgx.RowToStructByName` that reads `db` struct tags:

```go
type User struct {
    ID        pgtype.Int4        `db:"id"`
    Username  pgtype.Text        `db:"username"`
    Email     pgtype.Text        `db:"email"`
    Score     pgtype.Float8      `db:"score"`
    CreatedAt pgtype.Timestamptz `db:"created_at"`
    DeletedAt pgtype.Timestamptz `db:"deleted_at"`
}
```

Then query using `pgx.CollectOneRow` + `pgx.RowToStructByName`:

```go
func getUser(conn *pgx.Conn, id int) (*User, error) {
    query := `SELECT id, username, email, score, created_at, deleted_at FROM users WHERE id = $1`

    rows, err := conn.Query(ctx, query, id)
    if err != nil {
        return nil, err
    }

    // Automatically maps columns to struct fields by `db` tag
    user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }

    return &user, nil
}
```

For multiple rows:
```go
func getUsers(conn *pgx.Conn) ([]User, error) {
    query := `SELECT id, username, email, score, created_at, deleted_at FROM users`

    rows, err := conn.Query(ctx, query)
    if err != nil {
        return nil, err
    }

    // Collects all rows into []User
    users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
    if err != nil {
        return nil, err
    }

    return users, nil
}
```

---

## Option 3: `RowToStructByPos` (positional but cleaner)

If you don't want struct tags but still want the `CollectRow` pattern:

```go
user, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[User])
```
Maps purely by position — no tags needed, but column order must still match struct field order.

---

## Comparison

| Approach | Tag Needed | Order Sensitive | Best For |
|---|---|---|---|
| `Scan()` manually | No | ✅ Yes | Simple/few queries |
| `RowToStructByPos` | No | ✅ Yes | Cleaner syntax, fixed queries |
| `RowToStructByName` | `db:"col"` | ❌ No | Many queries, maintainability |

**`RowToStructByName` with `db` tags is the recommended approach** for most real-world use cases — it's the closest to `sqlx` ergonomics and is the most resilient to query changes.

---------------------------------------------------------

<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>

You don't need `rows.Next()` — pgx v5 has `pgx.CollectRows` for multiple rows, which handles the loop internally for you.

---

## Side by Side

```go
// --- ONE ROW ---
rows, _ := conn.Query(ctx, query, id)
user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[UserRow])


// --- MULTIPLE ROWS ---
rows, _ := conn.Query(ctx, query)
users, err := pgx.CollectRows(rows, pgx.RowToStructByName[UserRow])
// users is []UserRow
```

---

## What `CollectRows` does internally (simplified)

```
1. Calls rows.Next() in a loop internally
2. For each row, calls RowToStructByName to scan into UserRow
3. Appends each UserRow into a []UserRow
4. Calls rows.Close() automatically when done
```

So it replaces this manual pattern:

```go
// ❌ manual way — what CollectRows replaces
var users []UserRow
for rows.Next() {
    var u UserRow
    err := rows.Scan(&u.ID, &u.Username /* ... */)
    if err != nil {
        return nil, err
    }
    users = append(users, u)
}
if err := rows.Err(); err != nil {
    return nil, err
}
rows.Close()
```

```go
// ✅ pgx v5 way
users, err := pgx.CollectRows(rows, pgx.RowToStructByName[UserRow])
```

---

## Summary

| | Method | Returns |
|---|---|---|
| One row | `pgx.CollectOneRow` | `UserRow` |
| Many rows | `pgx.CollectRows` | `[]UserRow` |
| Manual loop | `rows.Next()` + `Scan()` | you manage everything |

**Always prefer `CollectRows` and `CollectOneRow`** — they handle `rows.Close()`, `rows.Err()`, and the loop for you, which are easy to forget with the manual approach.