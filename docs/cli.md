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
book list
```

This command will list out all books stored in the database, along with all of the fields.

## remove

To remove a book from the system:

```bash
book remove <ID>
```

For example,

```bash
book remove 0 
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

