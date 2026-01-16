# Database Migrations

Folder ini berisi SQL migration files untuk database Agora.

## Quick Start

### Option 1: Jalankan Full Schema (Recommended)
Jalankan file ini untuk membuat semua tabel sekaligus:

```bash
psql -U postgres -d agora_db -f migrations/001_init_schema.sql
```

### Option 2: Jalankan Step-by-Step
Jika ingin bertahap, mulai dengan users table:

```bash
psql -U postgres -d agora_db -f migrations/001_create_users_table.sql
```

## Cara Menjalankan Migration

### 1. Menggunakan psql CLI

```bash
# Connect ke database
psql -U postgres -d agora_db

# Jalankan migration
\i migrations/001_init_schema.sql

# Atau langsung dari command line
psql -U postgres -d agora_db -f migrations/001_init_schema.sql
```

### 2. Menggunakan GUI (pgAdmin, DBeaver, dll)

1. Buka file `001_init_schema.sql`
2. Copy seluruh isi file
3. Paste di SQL query editor
4. Execute

### 3. Menggunakan Docker (jika database di container)

```bash
docker exec -i postgres_container psql -U postgres -d agora_db < migrations/001_init_schema.sql
```

## Verifikasi

Setelah menjalankan migration, verifikasi dengan:

```sql
-- List semua tables
\dt

-- Atau dengan query
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public';

-- Check struktur users table
\d users
```

## Struktur Database

### Tables:
- `users` - User accounts
- `categories` - Forum categories
- `threads` - Discussion threads
- `posts` - Thread replies/comments
- `thread_votes` - Votes on threads
- `post_votes` - Votes on posts

### Relationships:
- `threads.user_id` → `users.id`
- `threads.category_id` → `categories.id`
- `posts.user_id` → `users.id`
- `posts.thread_id` → `threads.id`
- `posts.parent_post_id` → `posts.id` (for nested replies)

## Rollback

Jika perlu rollback (hapus semua tables):

```sql
DROP TABLE IF EXISTS post_votes CASCADE;
DROP TABLE IF EXISTS thread_votes CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS threads CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP FUNCTION IF EXISTS update_updated_at_column CASCADE;
```

## Notes

- Semua ID menggunakan UUID
- Password di-hash menggunakan bcrypt
- Timestamps otomatis dengan triggers
- Foreign keys dengan CASCADE delete
- Indexes untuk performance optimization
