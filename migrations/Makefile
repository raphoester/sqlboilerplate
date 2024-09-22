.PHONY: boil

CTX ?= comments
DSN ?= postgres://root@localhost:9999/root?sslmode=disable

boil:
	docker compose up --detach --wait

	cd ../ && POSTGRES_DSN=$(DSN) \
	MIGRATIONS_PATH=./migrations/$(CTX) \
	go run ./cmd/migrator

	cd ./$(CTX) && \
		sqlboiler \
			--config ../boiler.toml \
			--wipe \
			--output ../../generated/sqlboiler/$(CTX) \
			--pkgname $(CTX)b \
			psql
	 docker compose down