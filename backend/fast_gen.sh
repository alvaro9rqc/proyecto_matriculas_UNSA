rm -rf gen

#go mod tidy # escanea tu projecto y actualiza tu go.mod
goa gen github.com/enrollment/design/api
sqlc generate

echo "Code generation completed successfully."

