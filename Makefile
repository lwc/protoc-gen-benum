example: example/out
	go install && protoc example/example.proto -I=.  -I=benum/  --benum_out=example/out --go_out=example/out
