build:
	go build -o crawlerg cmd/main.go

run:
	go run crawlerg

install:
	install -d /opt/crawlerg/
	install -m 755 crawlerg /opt/crawlerg/
	install -m 755 top10milliondomains.txt /opt/crawlerg/