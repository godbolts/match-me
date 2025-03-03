docker compose up -d

docker compose down

docker volume rm $(docker volume ls -q --filter "name=postgres_data")