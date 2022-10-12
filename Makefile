.PHONY: example
example:
	go install && protoc example/*.proto --go_opt=paths=source_relative  -I=example/ -I=benum/ --benum_out=example/ --go_out=example/
