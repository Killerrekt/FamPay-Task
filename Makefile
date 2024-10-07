server-up:
	docker compose up -d

server-down:
	docker compose down

server-build:
	docker compose up --build -d

database-up:
	docker compose up postgres -d

database-down:
	docker compose down postgres

