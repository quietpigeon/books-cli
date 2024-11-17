# User guide

The basic usage of the commands:

```bash
book <COMMAND>
```

## add

To add a book, you can use a one-line command:

```bash
book add --title <TITLE> --author <AUTHOR> --published-date <DATE> --edition <EDITION> --genre <GENRE>
```

The field for `--published-date` should be in the format `YYYY-MM-DD`.

For books with more than one genre, you can add multiple genres using multiple `--genre` flags.

For example,

```bash
$ book add --title "Digital Fortress" --author "Dan Brown" --published-date 1998-01-01 --edition 1 --genre fiction --genre mystery
Would you like to add a short description? [y/n]
```

The description is optional and can be added after adding the book.

After a book is added, the system will assign an ID to the book. You can check the ID of a book using the `list` command.

## list

To list all of the books in the system:

```bash
book list --all
```

This command will list out the necessary fields, i.e. `id`, `title`, `author`, `collection`

To list out all of the fields, simply add the `--verbose` flag to the command:

```bash
book list --all --verbose
```

To list the books with specific author(s) or genre(s):

```bash
book list --author <AUTHOR>...
```

```bash
book list --genre <GENRE>...
```

You can also combine the two:

```bash
book list --author <AUTHOR> --genre <GENRE>
```

For example,

```bash
book list --author "Dan Brown" --author "Roald Dahl" --genre fiction
```

## remove

To remove a book or books from the system:

```bash
book remove <ID>...
```

For example,

```bash
book remove 0 3 4
```

## update

To update the information of a specific book, use the following command and include the necessary flags accordingly:

```bash
book update <ID> --title <TITLE> --author <AUTHOR> ...
```

For example, if you would like to update the title of the book with an ID of 0:

```bash
book update 0 --title "Digital Fortresses"
```

## collection

To create a new collection:

```bash
book collection create <COLLECTION_NAME>
```

To add a book or multiple books to a collection:

```bash
book collection <COLLECTION_ID> add <BOOK_ID>...
```

For example:

```bash
$ book collection create "Favourites"
The collection "Favourites" has been created with ID 0.
$ book collection 0 add 2 3
```
