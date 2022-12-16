package database

import (
	"log"
	"time"

	"gorm.io/gorm"
	"library.demo/rent/models"
)

func Seed() {
	seedUsers()
	seedBooks()
	seedBookItems()
}

func seedUsers() {
	clients := []models.Client{
		{Name: "Ahmed", Email: "email1@gg.co", Phone: "0123456789", Address: "address"},
		{Name: "Hassan", Email: "email2@gg.co", Phone: "0123456789", Address: "address"},
		{Name: "Younes", Email: "email3@gg.co", Phone: "0123456789", Address: "address"},
		{Name: "Zead", Email: "email4@gg.co", Phone: "0123456789", Address: "address"},
		{Name: "read", Email: "email5@gg.co", Phone: "0123456789", Address: "address"},
	}

	result := Instance.Create(&clients)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	for _, client := range clients {
		if client.ID != 0 {
			log.Println("client : ", client.ID, " inserted")
		}
	}

	log.Println(result.RowsAffected, " client inserted")
}

func seedBooks() {
	books := []models.Book{
		{Title: "book 1 title", Year: 1995, Writer: "Ahmed", Publisher: "Pub1", Location: "sh1", ImageURI: "https://c1.wallpaperflare.com/preview/304/436/847/book-embossing-leather-empty.jpg"},
		{Title: "book 2 title", Year: 2004, Writer: "Sat", Publisher: "Pub2", Location: "sh1", ImageURI: "https://c1.wallpaperflare.com/preview/304/436/847/book-embossing-leather-empty.jpg"},
		{Title: "book 3 title", Year: 1999, Writer: "Mohamed", Publisher: "Pub3", Location: "sh1", ImageURI: "https://c1.wallpaperflare.com/preview/304/436/847/book-embossing-leather-empty.jpg"},
		{Title: "book 4 title", Year: 2001, Writer: "Ahmed", Publisher: "Pub3", Location: "sh1", ImageURI: "https://c1.wallpaperflare.com/preview/304/436/847/book-embossing-leather-empty.jpg"},
		{Title: "book 5 title", Year: 2015, Writer: "Hammed", Publisher: "Pub2", Location: "sh1", ImageURI: "https://c1.wallpaperflare.com/preview/304/436/847/book-embossing-leather-empty.jpg"},
		{Title: "book 6 title", Year: 2020, Writer: "Rabia", Publisher: "Pub1", Location: "sh1", ImageURI: "https://c1.wallpaperflare.com/preview/304/436/847/book-embossing-leather-empty.jpg"},
		{Title: "book 7 title", Year: 1990, Writer: "Ahmed", Publisher: "Pub1", Location: "sh1", ImageURI: "https://c1.wallpaperflare.com/preview/304/436/847/book-embossing-leather-empty.jpg"},
	}

	result := Instance.Create(&books)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	for _, book := range books {
		if book.ID != 0 {
			log.Println("book : ", book.ID, " inserted")
		}
	}

	log.Println(result.RowsAffected, " book inserted")

}

func seedBookItems() {

	bookItems := []models.BookItem{
		{RentedAt: gorm.DeletedAt{Valid: true, Time: time.Now()}, BookRefer: 1, RenterRefer: 1},
		{RentedAt: gorm.DeletedAt{Valid: true, Time: time.Now()}, BookRefer: 2, RenterRefer: 1},
		{RentedAt: gorm.DeletedAt{Valid: true, Time: time.Now()}, BookRefer: 3, RenterRefer: 1},
		{RentedAt: gorm.DeletedAt{Valid: true, Time: time.Now()}, BookRefer: 4, RenterRefer: 1},
		{RentedAt: gorm.DeletedAt{Valid: true, Time: time.Now()}, BookRefer: 1, RenterRefer: 2},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 1},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 2},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 5},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 3},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 4},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 2},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 1},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 5},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 3},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 1},
		{RentedAt: gorm.DeletedAt{Valid: false}, BookRefer: 2},
	}

	result := Instance.Create(&bookItems)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	for _, bookItem := range bookItems {
		if bookItem.ID != 0 {
			log.Println("book Item : ", bookItem.ID, " inserted")
		}
	}

	log.Println(result.RowsAffected, " book Item inserted")

}
