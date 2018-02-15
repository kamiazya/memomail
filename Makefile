.PHONY: run mail-server-start mail-server-stop

run:
	go run main.go config.go

mail-server-start:
	docker run \
		-p 25:25 -p 1080:1080 \
		--rm -d \
		--name memomail-server \
		kamiazya/mailcatcher

mail-server-stop:
	docker stop memomail-server
