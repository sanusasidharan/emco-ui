export GO111MODULE=on
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=amd64

all: clean
	@go build -tags netgo -o ./middleend ./main/main.go

race: clean
	@go build -race -tags netgo -o ./middleend ./main/main.go

clean:
	@find . -name "*so" -delete
	@rm -f middleend 
swagger:
	@swagger generate spec -o ./swagger.yaml --scan-models