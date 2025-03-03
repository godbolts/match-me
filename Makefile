.PHONY: start make-executable reset-db start-db stop-db backend

start:
	./frontend/node_modules/.bin/concurrently --kill-others-on-fail \
		"make start-db" \
		"cd backend && go run ." \
		"cd frontend && npm start"

make-executable:
	chmod +x reset_db.sh

reset-db: make-executable
	./reset_db.sh

start-db:
	docker compose up -d

stop-db:
	docker compose down

backend:
	cd backend && go run .