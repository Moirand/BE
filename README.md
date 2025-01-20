# Description
This is the Back-End for Medpoint System, a platform that allows users to make various types of online medical reservations, including doctor consultations, laboratory tests, medical actions, and vaccinations.

# Technologies
- Language: Golang
- Framework: Raiden

# Project Structures

```bash
├── Medpoint
│   ├── cmd
│   │   ├── Medpoint
│   │   │   └── medpoint.go
│   │   ├── apply
│   │   │   └── main.go
│   │   └── import
│   │       └── main.go
│   ├── configs
│   │   └── app.yaml
│   ├── go.mod
│   └── internal
│       ├── bootstrap
│       │   ├── jobs.go
│       │   ├── models.go
│       │   ├── roles.go
│       │   ├── route.go
│       │   ├── rpc.go
│       │   └── storages.go
│       ├── controllers
│       │   └── hello.go
│       └── jobs
│           └── hello.go
├── configs
│   └── supabase.go
├── go.mod
├── go.sum
└── raiden-macos-amd64
```
