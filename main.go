package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username       string
	HashedPassword string
	Email          string
	Address        string
	PhoneNumber    string
	Role           string
}

func main() {
	// Datenbankverbindung öffnen
	db, err := sql.Open("sqlite3", "/home/lukas/mydatabase.de.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Benutzerdaten vom Benutzer eingeben
	var user User
	fmt.Println("Neuen Benutzer anlegen:")
	fmt.Print("Username: ")
	fmt.Scan(&user.Username)
	fmt.Print("Password: ")
	fmt.Scan(&user.HashedPassword)
	fmt.Print("Email: ")
	fmt.Scan(&user.Email)
	fmt.Print("Address: ")
	fmt.Scan(&user.Address)
	fmt.Print("Phone Number: ")
	fmt.Scan(&user.PhoneNumber)
	fmt.Print("Role (User/Admin): ")
	fmt.Scan(&user.Role)

	// Benutzer in die Datenbank einfügen
	err = insertUser(db, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User erfolgreich hinzugefügt!")

	users, err := getUser(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("all users:")
	for _, user := range users {
		fmt.Println("Username: %s, Email: %s\n", user.Username, user.Email)
	}
}

func insertUser(db *sql.DB, user User) error {
	sql := `
        INSERT INTO users (username, hashed_password, email, address, phone_number, role)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	_, err := db.Exec(sql, user.Username, user.HashedPassword, user.Email, user.Address, user.PhoneNumber, user.Role)
	if err != nil {
		return err
	}
	return nil
}
func getUser(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT *FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Username, &user.HashedPassword, &user.Email, &user.Address, &user.PhoneNumber, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}
