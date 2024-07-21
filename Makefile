run: 
	air

css:
	npx tailwindcss build -i static/css/main.css -o public/style.css --watch

templ:
	templ generate --watch