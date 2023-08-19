update:
	go clean -modcache
	go get -u github.com/dzoniops/...
update-all:
	go clean -modcache
	go get -u all 
