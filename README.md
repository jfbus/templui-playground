# TemplUI Playground
A playground for [TemplUI](https://github.com/jfbus/templui)
## Rebuilding the playground
```
go install github.com/jfbus/templui-playground/templui-playground-generator
go generate ./...
cd playground
templ generate
cd assets_src
npm run build
```
## Running the playground
```
PORT=8080 go run main.go
```