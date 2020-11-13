package scan

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Run(args []string) error {
	connectionString, err := buildConnString(args[0], args[1])
	if err != nil {
		fmt.Errorf("Failed to generate connection string, these inputs are invalid IP: %s and Port %s", args[0], args[1])
	}

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Errorf("Failed to generate connection object, received error %s", err)
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

func buildConnString(IP, port string) (string, error) {
	if net.ParseIP(IP) == nil {
		return "", errors.New("Bad IP address format")
	}
	intPort, err := strconv.Atoi(port)
	if err != nil {
		return "", errors.New("Bad Port format")
	}

	return fmt.Sprintf("tcp(%s:%d)/", IP, intPort), nil
}
