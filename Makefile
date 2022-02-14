build:
	docker-compose build 
run:
	docker-compose up -d --no-build 
clean:
	docker-compose down --volumes  --remove-orphans --rmi 'all'