redis:
	docker run -d -p 6379:6379 --name redis redis

postgres:
	docker run -d -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=postgres

serve:
	go run .