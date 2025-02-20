package main

import (
	"fmt"
	"github.com/elwin/ftp"
	"log"
	"os"
	"time"
)

func main() {

	c, err := ftp.Dial("localhost:2121",
		ftp.DialWithDebugOutput(os.Stdout),
		ftp.DialWithTimeout(60 * time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("admin", "123456")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn
	// c.Stor("test1.txt", strings.NewReader("My message"))
	// c.Stor("test2.txt", strings.NewReader("Bye World"))


	entries, _ := c.List("/")
	for _, entry := range entries {
		fmt.Println(entry.Name)
	}

	// c.Delete("test1.txt")
	// c.Delete("test2.txt")


	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
