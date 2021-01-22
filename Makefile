
build:
	docker build -t cirri/echo:dev .


# assuming cirri infra is running, https://echo_dev.<STACKDOMAIN>/metrics will show the prometheus endpoint, and the rest will be the echo code
run:
	docker run -dit --network cirri_proxy \
		--name echo_dev \
			cirri/echo:dev

stop:
	docker rm -f echo_dev