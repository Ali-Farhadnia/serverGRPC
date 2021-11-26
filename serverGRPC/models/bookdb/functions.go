package bookdb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Ali-Farhadnia/serverGRPC/connections"
	"github.com/Ali-Farhadnia/serverGRPC/models/book"
	uuid "github.com/satori/go.uuid"
)

//SetBookdb - set postgree connection to Book database to  work with
func (db *BookDB) SetBookdb() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		db.Config.Host, db.Config.Port, db.Config.User, db.Config.Password, db.Config.DbName, db.Config.Sslmode)
	res, err := connections.GetDBClientAccessPoint(connStr)
	if err != nil {
		return err
	}
	db.Db = res
	return nil
}

//CreateBooksTable create a table of books in book db and if table alredy exist it just return nil
func (db *BookDB) CreateBooksTable() error {
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
	_, err := db.Db.Exec(sqlStatement)
	if err != nil {
		switch err.Error() {
		case e:
			return nil
		case "":
			return err
		default:
			return nil
		}
	}
	return nil

}

/*Insert function takes a list of books and adds them to the books table.
If a book already exists, it only increases the book inventory.
book id is a uniqe id that each time a book insert to books table sets*/
func (db *BookDB) InsertBooks(books []book.Book) (string, error) {
	success := make([]string, 0)
	failed := make([]string, 0)

	for _, book := range books {
		var id string
		id, err := db.FindBookID(book)
		if err != nil {
			log.Println(err)
		}
		if err == sql.ErrNoRows {
			id := uuid.NewV4().String()

			sqlStatement := `
			INSERT INTO books (id, name, author, pagecount,inventory)
			VALUES ($1, $2, $3, $4,$5)
			RETURNING id`
			err := db.Db.QueryRow(sqlStatement, id, book.Name, book.Author, book.Pagecount, book.Inventory).Scan(&id)
			if err != nil {
				s, _ := book.ToString()
				failed = append(failed, s)
				continue
			}
			success = append(success, id)
			continue
		}
		if id != "" {

			book2, err := db.FindBookById(id)
			if err != nil {
				log.Println(err)
				s, _ := book.ToString()
				failed = append(failed, s)
				continue

			}
			book.Inventory += book2.Inventory
			_, err = db.UpdateBook(book, id)
			if err != nil {
				log.Println(err)
				s, _ := book.ToString()
				failed = append(failed, s)
				continue
			}
			success = append(success, id)
		}

	}
	output := fmt.Sprintf("success:%d:\n%s\n\n\nfailed:%d:\n%s\n\n", len(success), strings.Join(success, "\n"), len(failed), strings.Join(failed, "\n"))
	return output, nil

}

/*The function takes one book and the ID of another book as input,
and if there is a book with that ID, it replaces the input book with the book in the books table.*/
func (db *BookDB) UpdateBook(book book.Book, id string) (string, error) {
	sqlStatement := `
	UPDATE books
	SET name = $2, author = $3,pagecount=$4,inventory=$5
	WHERE id = $1;`
	res, err := db.Db.Exec(sqlStatement, id, book.Name, book.Author, book.Pagecount, book.Inventory)
	if err != nil {
		return "", err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return "there is no book with this id", err
	}
	return "", nil

}
func (db *BookDB) DeleteBook(id string) (string, error) {
	sqlStatement := `
	DELETE FROM books
	WHERE id = $1;`
	res, err := db.Db.Exec(sqlStatement, id)
	if err != nil {
		return "", err

	}
	i, _ := res.RowsAffected()
	if i == 0 {
		return "there is no book with this id", err
	}
	return "", nil

}

/*FindBookById function takes the book id and returns a book with that id in the books table.
If there is no book with the desired id, it returns the "there is no book with this id" message as an error.*/
func (db *BookDB) FindBookById(id string) (*book.Book, error) {
	sqlStatement := `SELECT * FROM books WHERE id=$1;`
	var book book.Book
	row := db.Db.QueryRow(sqlStatement, id)
	err := row.Scan(&book.ID, &book.Name, &book.Author,
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

/*Find function Checks if the book is on the books table or not,
and returns the book's id.and if book not exist return empty string as id and an error*/
func (db *BookDB) FindBookID(book book.Book) (string, error) {
	sqlStatement := `SELECT id FROM books WHERE name=$1 AND author =$2 AND 
	pagecount =$3;`
	var id string
	row := db.Db.QueryRow(sqlStatement, book.Name, book.Author, book.Pagecount)
	err := row.Scan(&id)

	switch err {
	case sql.ErrNoRows:
		return "", err
	case nil:
		return id, nil
	default:
		return "", err
	}
}
