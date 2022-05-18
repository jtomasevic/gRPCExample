# v2
# users 
protoc -I . \
--proto_path=pkg/services/users/proto \
--go_out pkg \
--go-grpc_out pkg \
pkg/services/users/proto/users-service.proto

# tasks 
protoc -I . \
--proto_path=pkg/services/tasks/proto \
--go_out pkg \
--go-grpc_out pkg \
pkg/services/tasks/proto/tasks-service.proto

protoc -I . \
--proto_path=pkg/services/tasks/proto \
--grpc-gateway_out pkg \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt generate_unbound_methods=true \
pkg/services/tasks/proto/tasks-service.proto

# facade 
protoc -I . \
--proto_path=pkg/services/facade/proto \
--go_out pkg \
--go-grpc_out pkg \
pkg/services/facade/proto/facade-service.proto

protoc -I . \
--proto_path=pkg/services/facade/proto \
--grpc-gateway_out pkg \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt generate_unbound_methods=true \
pkg/services/facade/proto/facade-service.proto

