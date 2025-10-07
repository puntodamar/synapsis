SHELL 				:= /bin/bash
PROTO_DIR 			:= proto
ORDER_DIR 			:= order-service
INVENTORY_DIR   	:= inventory-service
COMPOSE 			:= docker compose
ORDER_SERVICE 		:= order-service
INVENTORY_SERVICE   := inventory-service
TARGET 				?= dev

.PHONY: up
up:
	TARGET=$(TARGET) docker compose up --build

up-d:
	TARGET=$(TARGET) docker compose up --build -d

down:
	docker compose down -v

logs:
	docker compose logs -f --tail=200

#prod:
#	TARGET=release docker compose -f docker-compose.yml -f docker-compose.prod.yml up --build
#
#prod-d:
#	TARGET=release docker compose -f docker-compose.yml -f docker-compose.prod.yml up --build -d

.PHONY: restart
restart:
	docker compose restart (ORDER_SERVICE) (INVENTORY_SERVICE)

.PHONY: build
build:
	$(COMPOSE) build


.PHONY: migrate-order
migrate-order:
	$(COMPOSE) run --rm $(ORDER_SERVICE) goose -dir ./migrations postgres "$$DATABASE_URL" up

.PHONY: migrate-inventory
migrate-inventory: 
	$(COMPOSE) run --rm $(INVENTORY_SERVICE) goose -dir ./migrations postgres "$$DATABASE_URL" up

.PHONY: migrate-order-down
migrate-order-down:
	$(COMPOSE) run --rm $(ORDER_SERVICE) goose -dir ./migrations postgres "$$DATABASE_URL" down

.PHONY: migrate-inventory-down
migrate-inventory-down:
	$(COMPOSE) run --rm $(INVENTORY_SERVICE) goose -dir ./migrations postgres "$$DATABASE_URL" down

.PHONY: tidy
tidy:
	cd $(ORDER_DIR) && go mod tidy
	cd $(INVENTORY_DIR)   && go mod tidy

.PHONY: sh-order
sh-order:
	$(COMPOSE) exec $(ORDER_SERVICE) sh

.PHONY: sh-inventory
sh-inventory:
	$(COMPOSE) exec $(INVENTORY_SERVICE) sh

.PHONY: build-order build-inventory up-order up-inventory \
        rebuild-order rebuild-inventory restart-order restart-inventory

build-order:
	docker compose build order-service

build-inventory:
	docker compose build inventory-service

up-order:
	docker compose up -d order-service

up-inventory:
	docker compose up -d inventory-service

rebuild-order:
	docker compose up --build --no-deps -d order-service

rebuild-inventory:
	docker compose up --build --no-deps -d inventory-service

restart-order:
	docker compose restart order-service

restart-inventory:
	docker compose restart inventory-service


