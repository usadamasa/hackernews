
GSQLGEN := github.com/99designs/gqlgen

graphql-gen:
	go run $(GSQLGEN) generate
