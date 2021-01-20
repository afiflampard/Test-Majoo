install:
	go get -u github.com/gin-gonic/gin
dev:
	nodemon --exec go run main.go --signal SIGTERM