redis:
	docker run -d -p 6379:6379 --name redis redis

serve:
	go run .