#build binary :	go build -o ./bin/ph_xlxx.exe
start:
	go run ./src/

create-coverprofile:
	go test -coverprofile=tmp/profile.out ./...

watch-coverprofile:
	go tool cover -html=tmp/profile.out

test:
	go test ./src -v 
#flag : -run Testfunc