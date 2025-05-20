package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Note struct {
	ID    int
	Title string
	Body  string
}

func main() {
	db, err := sql.Open("sqlite", "notes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Створення таблиці, якщо ще не існує
	createTable := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT NOT NULL
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	// CRUD-операції
	addNote(db, "Перша нотатка", "Це текст першої нотатки.")
	addNote(db, "Друга нотатка", "Ще один текст.")

	fmt.Println("\n📋 Всі нотатки:")
	listNotes(db)

	updateNote(db, 1, "Оновлений заголовок", "Оновлений текст")
	deleteNote(db, 2)

	fmt.Println("\n📋 Після змін:")
	listNotes(db)
}

// Додавання нотатки
func addNote(db *sql.DB, title, body string) {
	stmt, err := db.Prepare("INSERT INTO notes(title, body) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ Нотатку додано.")
}

// Вивід усіх нотаток
func listNotes(db *sql.DB) {
	rows, err := db.Query("SELECT id, title, body FROM notes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Body); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[%d] %s: %s\n", n.ID, n.Title, n.Body)
	}
}

// Оновлення нотатки
func updateNote(db *sql.DB, id int, newTitle, newBody string) {
	stmt, err := db.Prepare("UPDATE notes SET title = ?, body = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(newTitle, newBody, id)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	fmt.Printf("✏️ Оновлено %d запис(ів).\n", affected)
}

// Видалення нотатки
func deleteNote(db *sql.DB, id int) {
	stmt, err := db.Prepare("DELETE FROM notes WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	fmt.Printf("🗑️ Видалено %d запис(ів).\n", affected)
}
