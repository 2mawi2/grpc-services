protoc -I . \
  --openapiv2_out ./gen/openapi \
  --openapiv2_opt logtostderr=true login.proto \
  --openapiv2_opt=generate_unbound_methods=true
