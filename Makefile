build: amd arm
amd:
	@docker buildx build --push --platform linux/amd64 -t dinhlockt02/cs_chat_app_server:dev-bulleye-slim-amd .
arm:
	@docker buildx build --push --platform linux/arm64 -t dinhlockt02/cs_chat_app_server:dev-bulleye-slim .
compose:
	@docker compose up -d
compose-down:
	@docker compose down