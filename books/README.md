## Books REST API
### with Go

#### to run
> go build && ./books

#### client
- postman
- insomnia

#### get all books
*access https://localhost:8000/api/books in any browser or use a client*

#### get a book
*send a get request using a client to https://localhost:8000/api/books/{id}*

#### create a book
*send a post request using a client to https://localhost:8000/api/books using header Content-Type application/json*

```
{
	"isbn": "31913",
	"title": "Book Three",
	"author": {
		"firstname": "Jack",
		"lastname": "Doe"
	}
}
```