.PHONY: boil

boil:
	@cd ./migrations && \
		$(MAKE) boil $(ARGS)

migration:
	@cd ./migrations/$(CTX) && \
		migrate create -ext sql -dir ./ -seq $(NAME)