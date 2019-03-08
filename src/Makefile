LINTERCOMMAND=gometalinter --vendor --tests --skip=mock --deadline=1500s

packages = \
	./config \
	./database \
	./entity\
	./model \
    ./services/users \
    ./services/auth  \
    ./services/common \
    ./services/trains \




setup:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	go get github.com/pressly/goose/cmd/goose

.PHONY: test
test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

.PHONY: code-quality
code-quality:
	$(LINTERCOMMAND) --checkstyle ./... > static-analysis.xml

.PHONY: code-quality-print
code-quality-print:
	$(LINTERCOMMAND) ./...

go-build:
	GOOS=linux GOARCH=amd64 go build -o cool_tasks

dc-build:
	docker-compose build
dc-up:
	docker-compose up &
