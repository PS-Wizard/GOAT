run:
	templ generate --watch --proxybind="http://0.0.0.0:8080" --cmd="go run ." & npx @tailwindcss/cli -i ./static/static.css -o ./static/output.css --watch
