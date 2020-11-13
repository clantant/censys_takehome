package scan

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Run(args []string) error {
	db, err := sql.Open("mysql", buildConnString(args[0], args[1]))
	if err != nil {
		errors.Wrapf(err, "Failed to generate connection, these inputs are invalid IP: %s and Port %s", args[0], args[1])
	}

	defer db.Close()

	available := db.Ping()
	if !strings.Contains(available.Error(), "1045") {
		fmt.Printf("No connection available at IP: %s with Port %s\n", args[0], args[1])
		return nil
	}

	fmt.Printf("Connected to MySQL server at IP: %s with Port: %s\n", args[0], args[1])

	var version string

	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil || version == "" {
		fmt.Printf("Failed to receive version info from MySQL")
		return nil
	}

	fmt.Printf("Received version: %s from database", version)

	return nil
}

func buildConnString(IP, port string) string {
	return fmt.Sprintf("tcp(%s:%s)/", IP, port)
}
