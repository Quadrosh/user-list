package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	_ "github.com/lib/pq"
	rp "github.com/quadrosh/user-list/repository/postgres"
)

// CommandLine is an interface for manage the bloackchain
type CommandLine struct {
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("initdb -dbname=DB_NAME    - initialisation of database")

}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

// Run executes commands
func (cli *CommandLine) Run() {
	cli.validateArgs()

	initdbCmd := flag.NewFlagSet("initdb", flag.ExitOnError)
	dbname := initdbCmd.String("dbname", "", "Database name")
	dbpass := initdbCmd.String("dbpass", "", "Database password")
	dbuser := initdbCmd.String("dbuser", "", "Database user")
	dbport := initdbCmd.String("dbport", "", "Database port")
	dbhost := initdbCmd.String("dbhost", "", "Database host")

	switch os.Args[1] {
	case "initdb":
		err := initdbCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if initdbCmd.Parsed() {
		if *dbname == "" || *dbuser == "" || *dbpass == "" {
			initdbCmd.Usage()
			runtime.Goexit()
		}
		cli.initdb(*dbhost, *dbport, *dbname, *dbuser, *dbpass)
	}

}

func (cli *CommandLine) initdb(dbHost, dbPort, dbname, dbuser, dbpass string) {
	repo, err := rp.ConnectPostgres(dbHost, dbPort, dbuser, dbname, dbpass)
	if err != nil {
		log.Fatal(err)
	}
	err = repo.CreateUserTableIfNotExists()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done!")
}
