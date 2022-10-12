.PHONY: example
example:
	go install && protoc example/*.proto -I=example/ -I=benum/ --benum_out=example/ --benum_opt=paths=source_relative --go_out=example/ --go_opt=paths=source_relative
