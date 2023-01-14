tidy:
	@go mod tidy

run: tidy
	@go run main.go

testing:
	@go test -v ./test

build:
	@rm -rf ./bin
	@sh -c 'CGO_LDFLAGS_ALLOW=${CGO_LDFLAGS_ALLOW} \
	go build -tags "osusergo netgo" -v -ldflags="-s -w \
	" -o ./bin/app'

run-docker: build
	@docker container run -t -i --rm \
	--memory="0.5g" \
	--cpus="1" \
	--net host \
	-v $(PWD):/data \
	-u $$(id -u ${USER}):$$(id -g ${USER}) \
	-e GODEBUG="madvdontneed=1" \
	-e TZ="Asia/Jakarta" \
	-w /data \
	redhat/ubi8-micro:8.5-437 ./bin/app 

docker:
	@docker build --tag lutfi/sequis:1.0.0 --force-rm -f Dockerfile .