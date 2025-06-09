# Requires:
<br>

`go install google.golang.org/protobuf/cmd/protoc-gen-go
@latest`

<br>

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

<br>

`brew install bufbuild/buf/buf`

<br>

`brew install docker`

<br>

`go install sigs.k8s.io/kind@v0.29.0`

<br>

# Compile protobuf using Buf:
<br>

`buf generate`

<br>

# Run k8s with Kind
<br>

`kind create cluster --config k8s/kind.yaml`

<br>

`kubectl apply -f k8s/server.yaml`

<br>

`kubectl apply -f k8s/client.yaml`

<br>

# Run project
<br>

`go run ./server/main.go 0.0.0.0:50051`

<br>

`go run ./client/main.go 0.0.0.0:50051`

