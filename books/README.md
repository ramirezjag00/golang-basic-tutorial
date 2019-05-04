## Books REST API
### with Go

#### to run
> clone project golang-basic-tutorial
> cd golang-basic-tutorial/books
> go get -u github.com/gorilla/mux
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

#### update a book
*send a put request using a client to https://localhost:8000/api/books/{id} using header Content-Type application/json*

```
{
	"isbn": "31913",
	"title": "Book Two",
	"author": {
		"firstname": "Jim",
		"lastname": "Doe"
	}
}
```

#### delete a book
*send a delete request using a client to https://localhost:8000/api/books/{id} using header Content-Type application/json*
