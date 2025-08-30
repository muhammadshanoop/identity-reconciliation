DOCKER_COMPOSE = docker compose
SERVICE = app 

# Start containers in detached mode
up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

# Rebuild images and restart
build:
	$(DOCKER_COMPOSE) up -d --build

# View logs for all services
logs:
	$(DOCKER_COMPOSE) logs -f

# View logs for app service only
logs-app:
	$(DOCKER_COMPOSE) logs -f $(SERVICE)

# Exec into app container
shell:
	$(DOCKER_COMPOSE) exec $(SERVICE) sh

# Show running containers
ps:
	$(DOCKER_COMPOSE) ps