# Postgres and Migrations

This folder contains SQL migrations for the Postgres database used by the services.

Files in `migrations/` are intended to be applied with a migration tool such as `migrate`.

How to apply migrations (recommended):

1. Start Postgres with Docker Compose:

```powershell
cd 'c:\Users\darre\Documents\GitHub\personal-devsecops-observability-stack\services'
docker-compose up -d db
```

2. Run the migrate container to apply migrations:

```powershell
# from services/
docker-compose run --rm migrate
```

This will run the `migrate` service which calls `migrate -path=/migrations -database postgres://postgres:postgres@db:5432/coffee?sslmode=disable up`.

You can also use the `migrate` CLI locally or add Makefile targets to run migrations from the host.