# Assessment

## Overview
This project serves as an assessment for the role of Backend Engineer for Mr. Cornellius.

## Installation & Running
### Build 
```go
docker compose build
```
### Run
```go
docker compose up
```

## Usage & Endpoints
here is Postman Documentation
```url
https://documenter.getpostman.com/view/1357220/2sA3Qy5pAB
```
## Books
### Add new book:
Method: POST
URL: http://localhost:9005/api/v1/books/store

This endpoint allows the user to add a new book to the database.
#### Request Body
- title (string, required): The title of the book.
- author_id (integer, required): The ID of the author of the book.
- category_id (integer, required): The ID of the category to which the book belongs.
- description (string, required): The description of the book.
- is_published (boolean, required): Indicates whether the book is published.
- isbn (string, required): The ISBN of the book.
```json
{
    "title": "Future Books part 1",
    "author_id": 1,
    "category_id": 1,
    "description": "hello",
    "is_published": true,
    "isbn": "978-1-16-148410-0"
}
```

#### Response:
```json
{
    "success": true,
    "data": {
        "id": 3,
        "title": "Future Books part 2",
        "isbn": "978-3-16-148410-0",
        "author_id": 1,
        "category_id": 1,
        "author": {
            "id": 1,
            "name": "Samsono",
            "email": "samsono@mail.com",
            "created_at": "2024-06-03T22:16:06+07:00",
            "updated_at": "2024-06-03T22:16:09+07:00"
        },
        "category": {
            "id": 1,
            "name": "Science Fiction",
            "description": "This is Science Fiction",
            "is_active": true,
            "created_at": "2024-06-03T22:16:33+07:00",
            "updated_at": "2024-06-03T22:16:36+07:00"
        },
        "description": "hello",
        "is_published": true,
        "created_at": "2024-06-03T15:38:48.543851+07:00",
        "updated_at": "2024-06-03T15:38:48.543851+07:00"
    },
    "error": {}
}

```

### Update book:
Method: PUT
URL: http://localhost:9005/api/v1/books/update

This endpoint allows the user to update the details of a book.
#### Request Body
- id (integer) - The unique identifier of the book.
- title (string, required): The title of the book.
- author_id (integer, required): The ID of the author of the book.
- category_id (integer, required): The ID of the category to which the book belongs.
- description (string, required): The description of the book.
- is_published (boolean, required): Indicates whether the book is published.
- isbn (string, required): The ISBN of the book.

```json
{
    "id": 1,
    "title": "Future Books part 1",
    "author_id": 1,
    "category_id": 1,
    "description": "hello",
    "is_published": true,
    "isbn": "978-1-16-148410-0"
}
```

#### Response:
```json
{
    "success": true,
    "data": null,
    "error": {}
}
```

### Get all books:
Method: GET
URL: http://localhost:9005/api/v1/books/all?page=1&perPage=10&searchBy=isbn&searchKeyword=97

This endpoint makes an HTTP GET request to retrieve a list of books based on the provided parameters. The request includes the page number, number of items per page, search criteria, and search keyword.

#### Query Parameters
- page (integer) - page number
- perPage (integer) - total data per page
- searchBy (string) - search criteria (author, isbn, title)
- searchKeyword (string) - search keyword

#### Response:
```json
{
    "success": true,
    "data": {
        "books": [
            {
                "id": 13,
                "title": "You can select any connections,         ",
                "isbn": "-3966508-1010",
                "author_id": 38,
                "category_id": 5,
                "author": {
                    "id": 38,
                    "name": "Alan Kelly",
                    "email": "alkelly@icloud.com",
                    "created_at": "2000-07-07T07:28:28+07:00",
                    "updated_at": "2015-04-17T12:48:46+07:00"
                },
                "category": {
                    "id": 0,
                    "name": "",
                    "description": "",
                    "is_active": null,
                    "created_at": "0001-01-01T00:00:00Z",
                    "updated_at": "0001-01-01T00:00:00Z"
                },
                "description": "To start working with your server in Navicat, you should first establish a connection or several connections using the Connection window. Navicat 15 has added support for the system-wide dark mode.",
                "is_published": true,
                "created_at": "2003-11-19T03:57:11+07:00",
                "updated_at": "2016-02-13T11:27:33+07:00"
            },
            ....
        ],
        "total_rows": 8,
        "total_page": 1
    },
    "error": {}
}

```
### Delete book:
Method: DELETE

URL: http://localhost:9005/api/v1/books/delete

This HTTP DELETE request is used to delete a book by its ID. The request should include a JSON payload with the key "id" representing the ID of the book to be deleted.
#### Request Body
- id (integer) - The unique identifier of the book.

```json
{
    "id": 1,
}
```

#### Response:
```json
{
    "success": true,
    "data": null,
    "error": {}
}
```

### Get book by ID:
Method: GET

URL: http://localhost:9005/api/v1/books/get?id=1

This endpoint makes an HTTP GET request to retrieve information about a specific book identified by the provided ID. The response of this request can be documented as a JSON schema to describe the structure and data types of the returned book information.
#### Request Parameter
- id (integer) - The unique identifier of the book.

```json
{
    "id": 1,
}
```

#### Response:
```json
{
    "success": true,
    "data": null,
    "error": {}
}
```

## Authors
### Add new author:
Method: POST

URL: http://localhost:9005/api/v1/authors/store

This endpoint is used to store author information.

#### Request Body
- name (string, required): The name of the author.
- email (string, unique, required): The email of the author.

```json
{
    "name": "Author name",
    "email": "test@mail.com",
}
```

#### Response:
```json
{
    "success": true,
    "data": {
        "id": 1,
        "name": "Author name",
        "email": "test@mail.com",
        "created_at": "2024-06-04T10:28:50.690829+07:00",
        "updated_at": "2024-06-04T10:28:50.690829+07:00"
    },
    "error": {}
}

```

### Update author:
Method: PUT

URL: http://localhost:9005/api/v1/authors/update

This endpoint allows the user to update the details of a author.
#### Request Body
- id (number): The unique identifier of the author.
- name (string): The updated name of the author.
- email (string): The updated email of the author.

```json
{
    "id": 1,
    "name": "Sampurl",
    "email": "hello@mail.com"
}
```

#### Response:
```json
{
    "success": true,
    "data": null,
    "error": {}
}
```

### Get all authors:
Method: GET

URL: http://localhost:9005/api/v1/authors/all?page=1&perPage=10

This endpoint retrieves a list of authors with pagination support.

#### Query Parameters
- page (integer) - page number
- perPage (integer) - total data per page

#### Response:
```json
{
    "success": true,
    "data": {
        "authors": [
            {
                "id": 1,
                "name": "Rosa Thompson",
                "email": "thompr@gmail.com",
                "created_at": "2006-11-15T19:53:07+07:00",
                "updated_at": "2022-01-28T17:21:48+07:00"
            },
            ...
        ],
        "total_rows": 50,
        "total_page": 5
    },
    "error": {}
}

```
### Delete author:
Method: DELETE

URL: http://localhost:9005/api/v1/authors/delete

This endpoint is used to delete an author.
#### Request Body
- id (integer) - The unique identifier of the author.

```json
{
    "id": 1,
}
```

#### Response:
```json
{
    "success": true,
    "data": null,
    "error": {}
}
```

### Get author by ID:
Method: GET

URL: http://localhost:9005/api/v1/books/get?id=1

The endpoint retrieves author information based on the provided ID.

#### Request Parameter
- id (integer) - The unique identifier of the author.

```json
{
    "id": 1,
}
```

#### Response:
```json
{
    "success": true,
    "data": {
        "id": 2,
        "name": "Andhana",
        "email": "andhanacool@mail.com",
        "created_at": "2024-06-04T10:28:50.690829+07:00",
        "updated_at": "2024-06-04T10:28:50.690829+07:00"
    },
    "error": {}
}
```


## Book categories
### Add new author:
Method: POST

URL: http://localhost:9005/api/v1/book_categories/store

This endpoint is used to store book category information.

#### Request Body
- name (text, required): The name of the book category.
- description (text, required): The description of the book category.
- is_active (boolean, required): Indicates if the book category is active.


```json
{
    "name": "Science Fiction",
    "description": "Hello description",
    "is_active": true
}
```

#### Response:
```json
{
    "success": true,
    "data": {
        "id": 1,
        "name": "Science Fiction",
        "description": "Hello description",
        "is_active": true,
        "created_at": "2024-06-04T11:45:49.095772+07:00",
        "updated_at": "2024-06-04T11:45:49.095772+07:00"
    },
    "error": {}
}

```

### Update book category:
Method: PUT

URL: http://localhost:9005/api/v1/book_categories/update

This endpoint is used to update book category details.
#### Request Body
- id (number): The unique identifier of the book category.
- name (string): The updated name of the book category.
- description (string): The updated description of the book category.
- is_active (boolean): Indicates if the book category is active.


```json
{
    "id": 1,
    "name": "Science Fiction",
    "description": "Hello description",
    "is_active": true
}
```

#### Response:
```json
{
    "success": true,
    "data": null,
    "error": {}
}
```

### Get all book categories:
Method: GET

URL: http://localhost:9005/api/v1/book_categories/all?page=1&perPage=10

This endpoint retrieves a list of book categories with pagination support.

#### Query Parameters
- page (integer, required): The page number to retrieve.
- perPage (integer, required): The number of book categories per page.

#### Response:
```json
{
    "success": true,
    "data": {
        "book_categories": [
            {
                "id": 1,
                "name": "Science Fiction",
                "description": "Hello description",
                "is_active": true,
                "created_at": "2024-06-04T11:45:49.095772+07:00",
                "updated_at": "2024-06-04T11:45:49.095772+07:00"
            }
        ],
        "total_rows": 1,
        "total_page": 1
    },
    "error": {}
}

```
### Delete book category:
Method: DELETE

URL: http://localhost:9005/api/v1/book_categories/delete

This endpoint is used to delete a book category.
#### Request Body
- id (integer) - The unique identifier of the author.

```json
{
    "id": 1,
}
```

#### Response:
```json
{
    "success": true,
    "data": null,
    "error": {}
}
```

### Get book category by ID:
Method: GET

URL: http://localhost:9005/api/v1/book_categories/get?id=2

The endpoint retrieves book category information based on the provided ID.

#### Request Parameter
- id (integer) - The unique identifier of the book category.

```json
{
    "id": 1,
}
```

#### Response:
```json
{
    "success": true,
    "data": {
        "id": 2,
        "name": "Science Fiction",
        "description": "Hello description",
        "is_active": true,
        "created_at": "2024-06-04T11:50:47.178007+07:00",
        "updated_at": "2024-06-04T11:50:47.178007+07:00"
    },
    "error": {}
}
```