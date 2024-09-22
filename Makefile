.PHONY: boil

boil:
	@cd ./migrations && \
		$(MAKE) boil $(ARGS)