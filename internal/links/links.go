package links

import (
	"log"

	database "github.com/usadamasa/hackernews/internal/pkg/db/mysql"
	"github.com/usadamasa/hackernews/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title, Address, UserID) VALUES (?,?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())

	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID") // changed
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	var links []Link
	for rows.Next() {
		var link Link
		var user users.User
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &user.ID, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &user
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return links
}
