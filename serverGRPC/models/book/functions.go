package book

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Ali-Farhadnia/serverGRPC/connections"
	uuid "github.com/satori/go.uuid"
)

//cinvert book to the json string
func (book Book) String() (string, error) {
	b, err := json.Marshal(book)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//CreateBooksTable create a table of books in book db and if table alredy exist it just return nil
func CreateBooksTable() error {
	db, err := connections.GetBookDb()
	if err != nil {
		return err
	}
	e := fmt.Sprintf("pq: relation %q already exists", "books")

	sqlStatement := `
	CREATE TABLE books (
		id VARCHAR ( 50 ) NOT NULL UNIQUE,
		name VARCHAR ( 50 ) NOT NULL,
		author VARCHAR ( 50 ) NOT NULL,
		pagecount INT NOT NULL,
		inventory INT8 NOT NULL
	);
	`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		switch err.Error() {
		case e:
			log.Println("books table created")
			return nil
		case "":
			log.Println("failed to create book tablr")
			return err
		default:
			log.Println("books table created")
			return nil
		}
	}
	return nil

}

/*Insert function takes a list of books and adds them to the books table.
If a book already exists, it only increases the book inventory.
book id is a uniqe id that each time a book insert to books table sets*/
func (b Book) InsertToDb() error {
	db, err := connections.GetBookDb()
	if err != nil {
		return err
	}

	var id string
	id, err = b.FindBookID()

	if err == sql.ErrNoRows {
		id := uuid.NewV4().String()

		sqlStatement := `
			INSERT INTO books (id, name, author, pagecount,inventory)
			VALUES ($1, $2, $3, $4,$5)
			RETURNING id`
		err := db.QueryRow(sqlStatement, id, b.Name, b.Author, b.Pagecount, b.Inventory).Scan(&id)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	if id != "" {

		book2, err := FindBookById(id)
		if err != nil {
			return err
		}
		b.Inventory += book2.Inventory
		err = b.UpdateBook(book2)
		if err != nil {
			return err
		}
	}

	return nil

}

/*Find function Checks if the book is on the books table or not,
and returns the book's id.and if book not exist return empty string as id and an error*/
func (b Book) FindBookID() (string, error) {
	db, err := connections.GetBookDb()
	if err != nil {
		return "", err
	}
	sqlStatement := `SELECT id FROM books WHERE name=$1 AND author =$2 AND 
	pagecount =$3;`
	var id string
	row := db.QueryRow(sqlStatement, b.Name, b.Author, b.Pagecount)
	err = row.Scan(&id)

	switch err {
	case sql.ErrNoRows:
		return "", err
	case nil:
		return id, nil
	default:
		return "", err
	}
}

/*FindBookById function takes the book id and returns a book with that id in the books table.
If there is no book with the desired id, it returns the "there is no book with this id" message as an error.*/
func FindBookById(id string) (*Book, error) {
	db, err := connections.GetBookDb()
	if err != nil {
		return nil, err
	}
	sqlStatement := `SELECT * FROM books WHERE id=$1;`
	var book Book
	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&book.ID, &book.Name, &book.Author,
		&book.Pagecount, &book.Inventory)
	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("there is no book with this id")
	case nil:
		return &book, nil
	default:
		return nil, err
	}
}

/*The function takes one book and the ID of another book as input,
and if there is a book with that ID, it replaces the input book with the book in the books table.*/
func (b Book) UpdateBook(book *Book) error {
	db, err := connections.GetBookDb()
	if err != nil {
		return err
	}
	sqlStatement := `
	UPDATE books
	SET name = $2, author = $3,pagecount=$4,inventory=$5
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, b.ID, book.Name, book.Author, book.Pagecount, book.Inventory)
	if err != nil {
		return err
	}
	i, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return errors.New("somthing whent wrong")
	}
	return nil

}
func DeleteBook(id string) error {
	db, err := connections.GetBookDb()
	if err != nil {
		return err
	}
	sqlStatement := `
	DELETE FROM books
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err

	}
	i, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return errors.New("somthing whent wrong")
	}
	return nil

}
