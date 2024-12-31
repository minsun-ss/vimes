build:
	@docker build -t vimes .
	@docker run --rm vimes -vvv
