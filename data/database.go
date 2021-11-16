package data

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// declare program struct
type Program struct {
	Id       int
	Name     string
	Inscope  string
	Outscope string
}

// hold database
var db *sql.DB

// checkErr - handle errors
func checkErr(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDb - establish and return db connection
func ConnectDb(dbName string) *sql.DB {

	db, err := sql.Open("sqlite3", dbName)
	checkErr(err)
	return db
}

// AddProgram - add a new bounty program to database
func AddProgram(db *sql.DB, program string, inscope, outscope []string) {

	// convert scope slices to strings
	convertIn := strings.Join(inscope, " ")
	convertOut := strings.Join(outscope, " ")

	// new program config
	newProgram := Program{
		Name:     program,
		Inscope:  convertIn,
		Outscope: convertOut,
	}

	// create programs table if not exists
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS programs (id INTEGER PRIMARY KEY, name TEXT, inscope TEXT, outscope TEXT)")
	checkErr(err)
	stmt.Exec()

	// insert new bounty program into database
	stmt, _ = db.Prepare("INSERT INTO programs (id, name, inscope, outscope) VALUES (?, ?, ?, ?)")
	stmt.Exec(nil, newProgram.Name, newProgram.Inscope, newProgram.Outscope)
	fmt.Printf("\nAdded program: %v\n", newProgram.Name)
}

// splitSlice - convert single string slice to multistring
func splitSlice(s string) []string {
	str1 := strings.Split(s, " ")
	str2 := strings.Join(str1, " ")
	str3 := strings.Split(str2, " ")
	return str3
}

// QueryProgram - query a database program and retrieve scope
func QueryProgram(db *sql.DB, program string) {

	rows, err := db.Query("SELECT id, name, inscope, outscope FROM programs WHERE name like '%" + program + "%'")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		ourProgram := Program{}
		err = rows.Scan(&ourProgram.Id, &ourProgram.Name, &ourProgram.Inscope, &ourProgram.Outscope)
		checkErr(err)
		fmt.Println("[Inscope]")
		str := splitSlice(ourProgram.Inscope)
		for _, l := range str {
			fmt.Println(l)
		}
		fmt.Println("\n[Outscope]")
		str1 := splitSlice(ourProgram.Outscope)
		for _, l := range str1 {
			fmt.Println(l)
		}
	}
	err = rows.Err()
	checkErr(err)
}

// ListPrograms - lists all programs in database
func ListPrograms(db *sql.DB) {

	rows, err := db.Query("SELECT name FROM programs")
	checkErr(err)
	var progs Program
	for rows.Next() {
		err = rows.Scan(&progs.Name)
		checkErr(err)
		fmt.Println(progs.Name)
	}
}

// PipeCommand - print only wildcard domains to stdout
func PipeCommand(db *sql.DB, program string, allProgs bool) {
	re := regexp.MustCompile(`\*\.[a-zA-Z]+\.[a-zA-z]+`)

	if allProgs {
		rows, err := db.Query("SELECT inscope FROM programs")
		checkErr(err)
		for rows.Next() {
			var prog Program
			err = rows.Scan(&prog.Inscope)
			checkErr(err)
			str := splitSlice(prog.Inscope)
			for _, l := range str {
				if l == re.FindString(l) {
					s1 := strings.Replace(l, "*", "", 1)
					s2 := strings.Replace(s1, ".", "", 1)
					fmt.Println(s2)
				}
			}
		}
	} else {
		rows, err := db.Query("SELECT inscope FROM programs WHERE name like '%" + program + "%'")

		checkErr(err)
		for rows.Next() {
			var prog Program
			err = rows.Scan(&prog.Inscope)
			checkErr(err)
			str := splitSlice(prog.Inscope)
			for _, l := range str {
				if l == re.FindString(l) {
					s1 := strings.Replace(l, "*", "", 1)
					s2 := strings.Replace(s1, ".", "", 1)
					fmt.Println(s2)
				}
			}
		}
	}
}

// RemoveProgram - remove program from database
func RemoveProgram(db *sql.DB, program string) {

	stmt, err := db.Prepare("DELETE FROM programs WHERE name = ?")
	checkErr(err)

	_, err = stmt.Exec(program)
	checkErr(err)

	fmt.Printf("Removed %v from database\n", program)
}
