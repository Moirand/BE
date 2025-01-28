# Description
This is the Back-End for Medpoint System, a platform that allows users to make various types of online medical reservations, including doctor consultations, laboratory tests, medical actions, and vaccinations.

# Technologies
- Language: Golang
- Framework: Raiden

# Project Structures

```bash
├── README.md
├── build
│   ├── import
│   ├── medpoint
│   └── state
├── cmd
│   ├── Medpoint
│   │   └── medpoint.go
│   ├── apply
│   │   └── main.go
│   └── import
│       └── main.go
├── configs
│   └── app.yaml
├── go.mod
├── go.sum
└── internal
    ├── bootstrap
    │   ├── jobs.go
    │   ├── models.go
    │   ├── roles.go
    │   ├── route.go
    │   ├── rpc.go
    │   ├── storages.go
    │   └── subscribers.go
    ├── controllers
    │   ├── doctor.go
    │   └── hello.go
    ├── jobs
    │   └── hello.go
    ├── models
    │   ├── addresses.go
    │   ├── appointments.go
    │   ├── doctor_schedules.go
    │   ├── doctors.go
    │   ├── healthcare_facilities.go
    │   ├── medical_records.go
    │   ├── payments.go
    │   ├── speciality_categories.go
    │   └── users.go
    ├── roles
    │   ├── admin.go
    │   ├── doctor.go
    │   ├── super_admin.go
    │   └── users.go
    ├── rpc
    ├── storages
    └── subscribers
```
