protoc -I=/usr/local/include -I=. \
    --go_out=../cproto \
    --cpp_out=../cproto/proto_cpp \
    *.proto