package scan

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func Run(args []string) error {
	db, err := sql.Open("mysql", buildConnString(args[0], args[1]))
	if err != nil {
		errors.Wrapf(err, "Failed to open connection, no MySQL at IP: %s with Port %s", args[0], args[1])
	}

	defer db.Close()

	var version string

	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		errors.Wrap(err, "Failed to get MySQL version")
	}

	fmt.Printf("Connected to MySQL server at IP: %s with Port: %s\nReceived version: %s from database", args[0], args[1], version)

	return nil
}

func buildConnString(IP, port string) string {
	return fmt.Sprintf("tcp(%s:%s)", IP, port)
}
