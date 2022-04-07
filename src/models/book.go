package models

import (
	"log"

	"github.com/Jeanpigi/go-api/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model 
	ID uuid.UUID `gorm:"type:uuid"`
	Author string `json:"author"`
	Title string `json:"title"`
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
	res, err := db.Query("INSERT INTO book (author, title) VALUES($1, $2)", b.Author, b.Title)
	if err != nil {
		return err
	}
	// Print result
	log.Println(res)
	defer db.Close()
	return c.Status(201).JSON(b)
	
}

//GetAllBooks is function which Select all books from postgreSQL
func GetAllBooks(c *fiber.Ctx) error {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM book ORDER BY id")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer db.Close()
	books := Books{}
	for rows.Next() {
		b := Book{}

		if err := rows.Scan(&b.ID, &b.Author, &b.Title); err != nil {
			return err // Exit if we get an error
		}

		books.Books = append(books.Books, b)
	}
	return c.JSON(books)
	
}


//UpdateBook is function which Update book from postgreSQL
func UpdateBook(c *fiber.Ctx) error {
	type updateBook struct {
		Author string `json:"author"`
		Title string `json:"title"`
	}

	db := database.GetConnection()
	// New Book struct
	var b *Book
	
	// Parse body into struct
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var updateNoteData updateBook

	_, err := db.Query("UPDATE book SET author=$1, title=$2 WHERE id=$3", b.Author, b.Title, b.ID)
	if err != nil {
		return err
	}
	b.Author = updateNoteData.Author
	b.Title = updateNoteData.Title

	defer db.Close()
	return c.JSON(fiber.Map{"status": "success", "message": "Book updated", "data": &b})
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
	log.Println(res)
	defer db.Close()
	// Return Employee in JSON format
	return c.JSON("Deleted it")
	
}
