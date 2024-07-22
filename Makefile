redis:
	docker run -d -p 6379:6379 --name redis redis

postgres:
	docker run -d -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=postgres

mongo:
	docker run -d -p 27017:27017 --name mongo mongo

serve:
	go run .

gqlgen:
	go run -mod=mod github.com/99designs/gqlgen