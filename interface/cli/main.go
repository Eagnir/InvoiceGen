package main

import (
	"InvoiceGen/entity"
	"InvoiceGen/infrastructure/repository"
	"InvoiceGen/usecase/adminUser"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello CMD")
	fmt.Println("-------------")

	os.Remove("test.db")

	db := repository.DBContext{}
	er := db.OpenContext()
	if er != nil {
		panic(er)
	}
	defer db.CloseContext()
	db.Context.AutoMigrate(entity.AllModels...)
	defaultData := entity.GenerateDefaultData()
	db.Context.Transaction(func(tx *gorm.DB) error {
		for _, model := range defaultData {
			tx.Create(model)
		}
		return nil
	})

	outputFormat := "text"

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listEntity := listCmd.String("entity", "clients", "Name of an entity like 'client' or 'company' to get a list of entity records")
	// list -entity adminusers
	//fooName := fooCmd.String("name", "", "name")

	if len(os.Args) < 2 {
		fmt.Println("expected 'list' subcommands")
		os.Exit(1)
	}
	primaryCommand := os.Args[1]
	lastCommand := os.Args[len(os.Args)-1]
	if strings.ToLower(lastCommand) == "json" {
		outputFormat = "json"
	}
	switch strings.ToLower(primaryCommand) {
	case "list":
		listCmd.Parse(os.Args[2:])
		w := &tabwriter.Writer{}
		// minwidth, tabwidth, padding, padchar, flags
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)
		defer w.Flush()

		switch strings.ToLower(*listEntity) {
		case "adminusers":

			auService := adminUser.NewService(repository.NewDBContext())
			users, err := auService.ListAll()
			if err != nil {
				panic(err)
			}

			switch outputFormat {
			case "text":
				fmt.Fprintf(w, "%s\t%s\t%s\t\n", "ID", "Name", "Email")
				fmt.Fprintf(w, "%s\t%s\t%s\t\n", "----", "----", "----")
				for _, user := range users {
					fmt.Fprintf(w, "%d\t%s\t%s\t\n", user.AdminUserId, user.Name, user.Email)
				}
			case "json":
				json, err := json.Marshal(users)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(json))
			}

		default:
			fmt.Printf("List command for %s not yet implemented\n", *listEntity)
		}
		fmt.Fprintf(w, "\n")
	}

}
