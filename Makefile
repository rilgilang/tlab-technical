run:
	go run ./cmd run

migration up:
	go run ./cmd migration up

migration down:
	go run ./cmd migration down

migration rollback 1:
	go run ./cmd migration rollback one-step