redis:
	podman run -d -p 6379:6379 --name redis redis

postgres:
	podman run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres

mongo:
	podman run -d -p 27017:27017 --name mongo mongo

serve:
	go run .

gqlgen:
	go run -mod=mod github.com/99designs/gqlgen