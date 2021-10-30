COMPOSE ?= docker-compose -f docker-compose.yml
COMPOSE_LOCAL ?= $(COMPOSE) -f docker-compose.local.yml
COMPOSE_PROD ?= $(COMPOSE) -f docker-compose.prod.yml

run-local: build-local
	$(COMPOSE_LOCAL) up -d

logs-local:
	$(COMPOSE_LOCAL) logs -f

rm-local:
	$(COMPOSE_LOCAL) stop
	$(COMPOSE_LOCAL) rm -f

build-local:
	$(COMPOSE_LOCAL) build


run-prod: build-prod
	$(COMPOSE_PROD) up -d

logs-prod:
	$(COMPOSE_PROD) logs -f

rm-prod:
	$(COMPOSE_PROD) stop
	$(COMPOSE_PROD) rm -f

build-prod:
	$(COMPOSE_PROD) build
