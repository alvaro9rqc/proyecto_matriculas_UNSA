go install goa.design/goa/v3/cmd/goa@latest

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

rm -rf gen

goa gen github.com/enrollment/design/api
sqlc generate

echo "Code generation completed successfully."
