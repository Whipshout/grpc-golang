# Requires:
<br>

`go install google.golang.org/protobuf/cmd/protoc-gen-go
@latest`

<br>

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

<br>

`brew install bufbuild/buf/buf`

# Compile protobuf using Buf:
<br>

`buf generate`

# Run project
<br>

`go run ./server/main.go 0.0.0.0:50051`

<br>

`go run ./client/main.go 0.0.0.0:50051`

