build:
	@docker build -t vimes .
	@docker run --rm \
		--env-file vimes.env \
		-p 24680:24680 \
		vimes -vvv
