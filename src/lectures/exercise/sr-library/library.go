//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string //book title
type Name string  //libaray member name

type LendAudit struct {
	checkout time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}
type BookEntry struct {
	total  int // total owned by the library
	lended int // total books lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {

	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			fmt.Println("[Not returned yet]")
		} else {
			returnTime = audit.checkout.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkout.String(), "through", returnTime)
	}
}
func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}
func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/lended:", book.lended)
	}
	fmt.Println()
}
func checkOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]

	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("no More books availabe to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkout: time.Now()}

	return true
}
func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Memeber did not checkOut this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {

	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}
	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The Little Go book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["let's learn go "] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["go Bootcamp "] = BookEntry{
		total:  1,
		lended: 0,
	}
	library.members["armin"] = Member{"armin", make(map[Title]LendAudit)}
	library.members["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.members["jim"]   = Member{"jim", make(map[Title]LendAudit)}
	library.members["mike"]  = Member{"mike", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	memebr:=library.members["armin"]
	checkOut:=checkOutBook(&library,"Go Bootcamp",&memebr)
	fmt.Println("\nCheck out a book:")
	if checkOut{
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
	returned:=returnBook(&library,"Go Bootcamp",&memebr)
	fmt.Println("\nCheck in a books:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
