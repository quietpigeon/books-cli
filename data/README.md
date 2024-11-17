# Data structure

This directory contains the JSON data storage for `book-cli`.

## `books.json`

This json file stores information about books. Each book is represented by a JSON object with the following fields:

* `id`: (integer) unique identifier for the book
* `title`: (string) title of the book
* `author`: (string) author of the book
* `published_date`: (string) publication date in YYYY-MM-DD format
* `edition`: (string) edition of the book (optional)
* `genre`: (array of strings) list of genres
* `description`: (string) brief description (optional)

Example:

```json
{
    "id": 1,
    "title": "Angels & Demons",
    "author": "Dan Brown",
    "published_date": "2000-05-01",
    "edition": "1st",
    "genre": ["Thriller", "Mystery"],
    "description": "A good thriller."
}
```

## `collections.json`

This json file stores information about book collections. Each collection is represented by a JSON object with the following fields:

* `id`: (integer) unique identifier for the collection
* `name`: (string) name of the collection
* `books`: (array of integers) list of book IDs in the collection

Example:

```json
{
  "id": 1,
  "name": "Dan Brown Thrillers",
  "books": [1, 2]
}
```
