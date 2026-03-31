# BE-Agora

Backend API untuk platform diskusi dan forum berbasis Go. Agora adalah platform modern yang memungkinkan pengguna untuk membuat thread diskusi, berinteraksi melalui posting, dan memberikan vote pada konten.

## Tentang Proyek

BE-Agora adalah backend API yang dibangun dengan Go menggunakan arsitektur Clean Architecture. Platform ini menyediakan fitur-fitur untuk:

- Manajemen pengguna dan otentikasi
- Sistem thread dan diskusi
- Posting dan komentar
- Sistem voting (upvote/downvote)
- Kategorisasi konten
- Caching dengan Redis

## Teknologi

### Backend

- **Go 1.25.3** - Bahasa pemrograman utama
- **Gin** - Web framework untuk REST API
- **PostgreSQL** - Database utama
- **Redis** - Caching layer
- **JWT** - Authentication & Authorization
- **Docker & Docker Compose** - Containerization

### Libraries Utama

```go
github.com/gin-gonic/gin         // Web framework
github.com/jackc/pgx/v5          // PostgreSQL driver
github.com/golang-jwt/jwt/v5     // JWT implementation
github.com/jmoiron/sqlx          // SQL extensions
github.com/spf13/viper           // Configuration management
github.com/google/uuid           // UUID generation
golang.org/x/crypto              // Cryptography
```

## Arsitektur

Proyek ini menggunakan **Clean Architecture** dengan struktur:

```
be-agora/
├── cmd/
│   └── api/                 # Entry point aplikasi
├── internal/
│   ├── config/              # Konfigurasi aplikasi
│   ├── domain/              # Domain entities
│   │   ├── user.go
│   │   ├── thread.go
│   │   ├── post.go
│   │   ├── category.go
│   │   └── vote.go
│   ├── handler/             # HTTP handlers
│   │   └── http/
│   ├── repository/          # Data access layer
│   │   ├── postgres/        # PostgreSQL implementation
│   │   └── redis/           # Redis implementation
│   ├── service/             # Business services
│   ├── usecase/             # Business logic
│   └── pkg/                 # Shared packages
├── migrations/              # Database migrations
├── docker-compose.yml       # Docker configuration
└── Dockerfile              # Docker build file
```

### Lapisan Arsitektur:

1. **Handler Layer** - HTTP request/response handling
2. **Use Case Layer** - Business logic implementation
3. **Repository Layer** - Data access abstraction
4. **Domain Layer** - Core business entities

## Instalasi

### Prerequisites

- Go 1.25.3+
- Docker & Docker Compose
- PostgreSQL 15+ (opsional, jika tidak menggunakan Docker)

### Clone Repository

```bash
git clone https://github.com/srgjo27/be-agora.git
cd be-agora
```

### Install Dependencies

```bash
go mod download
```

## Konfigurasi

1. **Buat file `.env`** berdasarkan template:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5433
DB_USER=agora_user
DB_PASSWORD=your_password
DB_NAME=agora_db

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key

# Server Configuration
PORT=8080
```

2. **Konfigurasi Docker** (opsional):
   - File `docker-compose.yml` sudah dikonfigurasi
   - Sesuaikan environment variables sesuai kebutuhan

## Menjalankan Aplikasi

### Dengan Docker (Recommended)

```bash
# Build dan jalankan semua services
docker-compose up --build

# Jalankan di background
docker-compose up -d --build

# Lihat logs
docker-compose logs -f api
```

### Manual (Tanpa Docker)

```bash
# Pastikan PostgreSQL berjalan di port yang benar
# Jalankan aplikasi
go run cmd/api/main.go
```

Server akan berjalan di `http://localhost:8080`

## API Documentation

### Base URL

```
http://localhost:8080/api/v1
```

### Coding Standards

- Gunakan `gofmt` untuk formatting
- Ikuti konvensi penamaan Go
- Update dokumentasi sesuai perubahan

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Kontak

**Project Link**: [https://github.com/srgjo27/be-agora](https://github.com/srgjo27/be-agora)

---
