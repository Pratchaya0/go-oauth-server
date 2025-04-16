#OAuth & User service microservice design in mono repo 

##Regenerate gRPC code 
```bash
# auth
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./modules/auth/authProtobuf/auth.proto

# user
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./modules/user/userProtobuf/user.proto
```