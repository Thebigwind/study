代码检查

go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

golangci-lint -D staticcheck run ./...

【代码格式化，语法检查】
go fmt ./...
go vet ./...
golangci-lint -D staticcheck run ./...

【golangci-lint下载安装】
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

brew install golangci/tap/golangci-lint
brew upgrade golangci/tap/golangci-lint

