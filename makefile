# Starting the server and the client
serve:
	echo "Running server ..."
	go run cmd/main.go

dev:
	echo "Running client ..."
	npm run serve --prefix "./front-end/telegraph/" -- --port 3000
