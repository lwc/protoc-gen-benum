.PHONY: example
example:
	go install && protoc example/*.proto -I=example/ -I=benum/ --benum_out=example/ --go_out=example/
