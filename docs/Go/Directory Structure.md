# Directory Structure

```text
myapp/
├── cmd/
├── internal/
├── pkg/
├── api/
├── configs/
├── scripts/
├── build/
├── deployments/
├── test/
├── web/
├── go.mod
├── go.sum
└── README.md
```

### cmd

Application entrypoints. Contains main programs for your application.

Each subfolder produces a separate executable.

```text
cmd/
 ├── server/
 │   └── main.go
 └── worker/
     └── main.go
```

### internal

Private application code. Contains code only usable inside the module.

```text
internal/
 ├── server/
 │   └── server.go
 ├── database/
 │   └── db.go
 └── auth/
     └── auth.go
```

External projects cannot import anything from inside internal.

### pkg 

Public reusable libraries. Contains packages intended to be imported by other projects.

```text
pkg/
 ├── logger/
 │   └── logger.go
 └── cache/
     └── cache.go
```

### api

API definitions

```text
api/
 └── proto/
     └── user.proto
```

### configs
configs

```text
configs/
 ├── dev.yaml
 ├── prod.yaml
 └── config.go
```

### scripts 

automation scripts

```text
scripts/
 ├── build.sh
 ├── migrate.sh
 └── test.sh
```

### build

Build and packaging files

```text
build/
 ├── Dockerfile
 └── ci/
```

### deployment

Infrastructure configs

```
deployments/
 ├── docker-compose.yml
 ├── kubernetes/
 └── helm/
```

### test

Test files that tests the whole system. Separate test for each file should be beside it where ever they are.

```text
test/
 ├── integration_test.go
 └── e2e_test.go
```

### web

Front end bullshit

```text
web/
 ├── static/
 └── templates/
```