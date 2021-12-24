package models

import (
	"fmt"

	"github.com/Jeanpigi/go-api/src/database"
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID uint64 `json:"id,omitempty"`
	Name string `json:"name"`
	Writer string `json:"writer"`
}

type Books struct {
	Books []Book `json:"Books"`
}

//InsertBook is function which add a book from postgreSQL
func InsertBook(c *fiber.Ctx) error {
	db := database.GetConnection()	
	// New Book struct
	b := new(Book)
	// Parse body into struct
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	// Insert Book into database
	_, err := db.Query("INSERT INTO book (name, writer) VALUES($1, $2)", &b.Name, &b.Writer)
	if err != nil {
		return err
	}
	// Print result
	fmt.Println("El libro se guardo correctamente")
	defer db.Close()
	return c.Status(201).JSON(b)
	
}

//GetAllBooks is function which Select all books from postgreSQL
func GetAllBooks(c *fiber.Ctx) error {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM book ORDER BY id ASC LIMIT 100")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer db.Close()
	books := Books{}
	for rows.Next() {
		b := Book{}

		if err := rows.Scan(&b.ID, &b.Name, &b.Writer); err != nil {
			return err // Exit if we get an error
		}

		books.Books = append(books.Books, b)
	}
	return c.JSON(books)
	
}


//UpdateBook is function which Update book from postgreSQL
func UpdateBook(c *fiber.Ctx) error {
	db := database.GetConnection()
	// New Book struct
	b := new(Book)
	// Parse body into struct
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	res, err := db.Query("UPDATE book SET name=$1, writer=$2 WHERE id=$3", &b.Name, &b.Writer, &b.ID)
	if err != nil {
		return err
	}
	// Print result
	fmt.Println(res)
	defer db.Close()
	return c.JSON(b)
}

//DeleteBook is function which Delete book from postgreSQL
func DeleteBook(c *fiber.Ctx) error {
	db := database.GetConnection()
	// New Book struct
	b := new(Book)
	// Parse body into struct
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	res, err := db.Query("DELETE FROM book WHERE id=$1", b.ID)
	if err != nil {
		return err
	}
	// Print result
	fmt.Println(res)
	defer db.Close()
	// Return Employee in JSON format
	return c.JSON("Deleted it")
	
}
