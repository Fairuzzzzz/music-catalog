# music-catalog

Music Catalog adalah aplikasi yang memungkinkan pengguna untuk mencari lagi,
mendaftar, login , dan mendapatkan rekomendasi lagu berdasarkan aktivitas mereka.

## Konfigurasi

Konfigurasi aplikasi dapat diatur dalam file `config.yaml` yang berada di direktori `internal/configs/`.
Berikut adalah contoh konfigurasi:
```yaml
service:
  port: ":9999"

database:
  dataSourceName: "postgres://admin:root@localhost:5432/test_db?sslmode=disable"

spotifyConfig:
  clientID: ""
  clientSecret: ""
```

Ubah clientID dan clientSecret dengan token spotify developers anda

## Instalasi

1. Clone Repository
```sh
git clone https://github.com/Fairuzzzzz/music-catalog.git
cd music-catalog
```

2.  Instalasi Docker
```sh
docker-compose up -d
```

3. Menjalankan secara lokal
```sh
go run cmd/main.go
```

Aplikasi akan berjalan pada port yang ditentukan dalam konfigurasi (`:9999`)

## Endpoint API

Berikut adalah beberapa endpoint yang tersedia dalam aplikasi:

### Memberships
| Method | Endpoint            | Deskripsi                  | Request Body                                                                                       | Response Body                                                                                       |
|--------|---------------------|----------------------------|----------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| POST   | /memberships/sign-up | Mendaftar pengguna baru    | `{ "email": "string", "username": "string", "password": "string" }`                                 | Status Code: 201 Created                                                                            |
| POST   | /memberships/login   | Login pengguna             | `{ "email": "string", "password": "string" }`                                                      | `{ "accessToken": "string" }`                                                                       |

### Tracks

| Method | Endpoint                   | Deskripsi                        | Request Body                                                                                       | Response Body                                                                                       |
|--------|----------------------------|----------------------------------|----------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------|
| GET    | /tracks/search             | Mencari lagu                     | Query Params: `query=string&pageSize=int&pageIndex=int`                                             | `{ "limit": int, "offset": int, "items": [ { "albumType": "string", "albumTotalTracks": int, ... } ], "total": int }` |
| POST   | /tracks/track-activity     | Menambahkan aktivitas track      | `{ "spotifyID": "string", "isLiked": bool }`                                                       | Status Code: 200 OK                                                                                 |
| GET    | /tracks/recommendation     | Mendapatkan rekomendasi lagu     | Query Params: `limit=int&trackID=string`                                                            | `{ "items": [ { "albumType": "string", "albumTotalTracks": int, ... } ] }`                          |

## Struktur Direktori

Berikut struktur direktori dari project ini:
```
music-catalog/
├── cmd/
│   └── main.go
├── internal/
│   ├── configs/
│   │   ├── config.yaml
│   │   ├── configs.go
│   │   └── types.go
│   ├── handler/
│   │   ├── memberships/
│   │   │   ├── handler.go
│   │   │   ├── login.go
│   │   │   └── signup.go
│   │   ├── tracks/
│   │   │   ├── handler.go
│   │   │   ├── recommendations.go
│   │   │   ├── search.go
│   │   │   └── track_activities.go
│   ├── middleware/
│   │   └── middleware.go
│   ├── models/
│   │   ├── memberships/
│   │   │   └── user.go
│   │   ├── spotify/
│   │   │   └── spotify.go
│   │   └── trackactivities/
│   │       └── trackactivities.go
│   ├── repository/
│   │   ├── memberships/
│   │   │   ├── repository.go
│   │   │   └── users.go
│   │   ├── spotify/
│   │   │   ├── outbound.go
│   │   │   ├── recommendations.go
│   │   │   ├── response.go
│   │   │   ├── search.go
│   │   │   └── token.go
│   │   └── trackactivities/
│   │       ├── repository.go
│   │       └── trackactivites.go
│   ├── service/
│   │   ├── memberships/
│   │   │   ├── login.go
│   │   │   ├── service.go
│   │   │   └── signup.go
│   │   ├── tracks/
│   │   │   ├── recommendations.go
│   │   │   ├── search.go
│   │   │   └── track_activites.go
│   └── pkg/
│       ├── httpclient/
│       │   ├── client.go
│       └── internalsql/
│           └── sql.go
│       └── jwt/
│           └── jwt.go
└── README.md
```

## Testing

Untuk menjalankan unit test, gunakan perintah berikut:

```sh
go test ./...
```
