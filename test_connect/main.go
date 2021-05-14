package main

import (
	"fmt"
	"database/sql"
	"log"
        _ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// connect to db
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=postgres password=StrongAdminP@ssw0rd")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to db!")

	// test db connection
	err = conn.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot ping db: %v\n", err))
        }
        log.Println("Pinged the db!")

	// select rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// insert
	query := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(query, "Julie", "Ch")
	if err != nil {
                  log.Fatal(err)
         }
	 log.Println("Inserted a row")

	// select again
	err = getAllRows(conn)
	if err != nil {
                  log.Fatal(err)
          }

	// update
	stmt := `update users set last_name = $1 where first_name = $2`
	_, err = conn.Exec(stmt, "Chinal", "Julie")
	if err != nil {
                  log.Fatal(err)
          }

	 log.Println("updated a row")

	// select and verify
	err = getAllRows(conn)
        if err != nil {
                  log.Fatal(err)
          }

	// select by id
	stmt = `select id, first_name, last_name from users where id = $1`
	var firstName, lastName string
	var id int

	row := conn.QueryRow(stmt, 2)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
                  log.Fatal(err)
         }
	 log.Println("QueryRow returns", id, firstName, lastName)

         log.Println("updated a row")



	// delect row
	query = `delete from users where first_name = $1`
	_, err = conn.Exec(query, "Julie")
	if err != nil {
                  log.Fatal(err)
          }

         log.Println("deleted row/s")


	// verify delete
	err = getAllRows(conn)
        if err != nil {
                log.Fatal(err)
        }


}

func getAllRows(conn *sql.DB) error {
//	rows, err := conn.Query("select id, first_name, last_name from users")
	rows, err := conn.Query("select * from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}

	fmt.Println("Record is:", id, firstName, lastName)
      }

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("----------------------------")

	return nil
}
