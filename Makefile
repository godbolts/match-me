.PHONY: start

start:
# Start frontend and backend at the same time
	./frontend/node_modules/.bin/concurrently --kill-others-on-fail \
		"make start-db"\
		"sh -c 'cd backend && go run main.go'"\
		"cd frontend && npm start" 

make-executable:
	chmod +x reset_db.sh

# Resets the database, meaning it deletes all the data and initialises it again with the default data from init.sql
reset-db: make-executable
	./reset_db.sh

# Starts the database // HS removed the hyphen to make it work in WSL2 
start-db:
	docker compose up -d

# Stops the database // HS removed the hyphen to make it work in WSL2 
stop-db:
	docker compose down

backend:
	cd backend && go run .
