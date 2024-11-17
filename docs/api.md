# `book-cli` REST API

There are two components in REST API, which are:

* `/books`

* `/collections`

---

## `GET /`

---

## `GET /v1/books`

* Description: get all books stored in the system
* Access: open
* Operation: sync
* Return: list of book objects

### Response

Example:

```json
[
  {
    "id": 1,
    "title": "Angels & Demons",
    "author": "Dan Brown",
    "published_date": "2000-05-01",
    "edition": "1st",
    "genre": ["Thriller", "Mystery"],
    "description": "A good thriller."
  },
  {
    "id": 2,
    "title": "The Da Vinci Code",
    "author": "Dan Brown",
    "published_date": "2003-03-18",
    "edition": "1st",
    "genre": ["Thriller", "Mystery"],
    "description": null
  },
  {
    "id": 3,
    "title": "The Lost Symbol",
    "author": "Dan Brown",
    "published_date": "2009-09-15",
    "edition": "1st",
    "genre": ["Thriller", "Mystery"],
    "description": "A mysterious mystery."
  }
]
```

### Fields

* `id`: a unique identifier of the book
* `title`: the title of the book
* `author`: the author of the book
* `published_date`: the date of publication of the book in the format (YYYY-MM-DD)
* `edition`: the edition of the book (optional)
* `genre`: the genre(s) the book belongs to
* `description`: a description of the book (optional)

---

## `GET /v1/books/{id}`

* Description: get a specific book by ID
* Access: open
* Operation: sync
* Return: a book object

### Response

Example:

```json
{
  "id": 1,
  "title": "Angels & Demons",
  "author": "Dan Brown",
  "published_date": "2000-05-01",
  "edition": "1st",
  "genre": ["Thriller", "Mystery"]
}
```

### Fields

See fields from `GET /v1/books`.

---

## `POST /v1/books`

* Description: add a new book
* Access: open
* Operation: sync
* Return: dict with the ID of the book added

### Response

Example:

```json
{
    "id": 4
}
```

### Request

Example:

```json
{
  "title": "Inferno",
  "author": "Dan Brown",
  "published_date": "2013-05-14",
  "edition": "1st",
  "genre": ["Thriller", "Mystery"] 
}
```

### Fields

See fields from `GET /v1/books`.

---

## `PUT /v1/books/{id}`

* Description: update an existing book by ID
* Access: open
* Operation: sync
* Return: 200 OK or an error

```json
Request:
{
  "edition": "2nd",
  "genre": ["Thriller", "Mystery", "Suspense"] 
}
```

### Fields

See fields from `GET /v1/books`. All fields are optional.

---

## `DELETE /v1/books/{id}`

* Description: delete a book from the system
* Access: open
* Operation: sync
* Return: 200 OK or an error

---

## `GET /v1/collections`

* Description: get all collections
* Access: open
* Operation: sync
* Return: list of collection objects

### Response

Example:

```json
[
  {
    "id": 1,
    "name": "Dan Brown Thrillers",
    "books": [1, 2, 3]
  }
]
```

### Fields

* `id`: a unique identifier for the collection
* `name`: name of the collection
* `books`: IDs of books that belong to the collection

---

## `GET /v1/collections/{id}`

* Description: get a specific collection by ID
* Access: open
* Operation: sync
* Return: a collection object

### Response

Example:

```json
{
  "id": 1,
  "name": "Dan Brown Thrillers",
  "books": [1, 2]
}
```

### Fields

See fields from `GET /v1/collections`.

---

## `POST /v1/collections`

* Description: create a new collection
* Access: open
* Operation: sync
* Return: dict with the ID of the created collection

### Request

Example:

```json
{
  "name": "Robert Langdon Series"
}
```

### Fields

See fields from `GET /v1/collections`.

---

## `PUT /v1/collections/{id}`

* Description: update an existing collection
* Access: open
* Operation: sync
* Return: 200 OK or an error

### Request

```json
{
  "name": "A better name."
}
```

### Fields

See fields from `GET /v1/collections`.

---

## `DELETE /v1/collections/{id}`

* Description: delete a collection
* Access: open
* Operation: sync
* Return 200 OK or an error

## `POST /v1/collections/{id}/books`

* Description: add a book to a collection
* Access: open
* Operation: sync
* Return: 200 OK or an error

### Request

```json
{
    "book_id": 3
}
```

## `DELETE /v1/collections/{id}/books`

* Description: remove a book from a collection
* Access: open
* Operation: sync
* Return: 200 OK or an error

### Request

```json
{
    "book_id": 2
}
```
