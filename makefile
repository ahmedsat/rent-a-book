
default:
	@echo "default"

dev:
	cp example.env .env
	watchexec -e go -c -r go run main.go

setup:
	wget https://github.com/watchexec/watchexec/releases/download/v1.20.6/watchexec-1.20.6-x86_64-unknown-linux-musl.deb 
	sudo dpkg -i watchexec-1.20.6-x86_64-unknown-linux-musl.deb 
	rm watchexec-1.20.6-x86_64-unknown-linux-musl.deb 