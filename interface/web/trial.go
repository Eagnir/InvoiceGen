package main

import (
	"InvoiceGen/entity"
	"InvoiceGen/infrastructure/repository"
	"InvoiceGen/usecase/adminUser"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func trial() {
	fmt.Println("Hello Web")
	fmt.Println("-------------")

	/* invoice, ex := entity.NewInvoice(entity.InvoiceCreated)
	if ex == nil {
		invoice.AddTagByName("CLS Project")
		invoice.AddTagByName("Home Coming")
		invoice.AddTagByName("Urgent")
		invoice.AddTagByName("Overdue")

		//var x := append(invoice.Tags[:2])
		//var y := append(invoice.Tags[2:])

		//fmt.Println("Before 2: ", x, y)

		for _, value := range invoice.Tags {
			fmt.Println(value.Name)
		}

		fmt.Println("-------------")

		invoice.RemoveTagByName("Urgent")
		tg, ex := entity.NewTag("Home Coming")
		if ex == nil {
			invoice.RemoveTag(tg)
		}

		for _, value := range invoice.Tags {
			fmt.Println(value.Name)
		}
	} */

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

	auService := adminUser.NewService(db)

	obj, err := entity.NewAdminUser("Trial Account", "Trial@gmail.com", "123")
	if err != nil {
		panic(err)
	}
	auService.SaveObject(obj)

	_, err = auService.SaveObjectFromNew(entity.NewAdminUser("Ethen", "ethen@gmail.com", "123"))
	if err != nil {
		panic(err)
	}
	_, err = auService.SaveObjectFromNew(entity.NewAdminUser("Jane", "janedoe@gmail.com", "373"))
	if err != nil {
		panic(err)
	}
	_, err = auService.SaveObjectFromNew(entity.NewAdminUser("John", "jdoe@gmail.com", "333"))
	if err != nil {
		panic(err)
	}

	/* _, er := auService.DeleteByObject(&entity.AdminUser{Password: "123"})
	if er != nil {
		panic(er)
	}
	*/
	niravUser, err := auService.GetEntityById(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello " + niravUser.Email)

	fmt.Println("----------------")

	users, err := auService.Search(entity.AdminUser{
		Password: "123",
	})
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user.Name)
	}

	/*
		var wg sync.WaitGroup

		wg.Add(2)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			db := &repository.DBContext{}
			db.OpenContext()
			defer db.CloseContext()
			time.Sleep(time.Second * 20)
			companies := &[]entity.Company{}
			db.Context.First(companies)
			for _, company := range *companies {
				fmt.Println(company.Email)
			}
		}(&wg)
		//wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			db := &repository.DBContext{}
			db.OpenContext()
			defer db.CloseContext()
			time.Sleep(time.Second * 10)
			companies := &[]entity.Company{}
			db.Context.First(companies)
			for _, company := range *companies {
				fmt.Println(company.Name)
			}
		}(&wg)

		wg.Wait() */

	/* invoice, ex := entity.NewInvoice(entity.InvoiceCreated)
	if ex == nil {
		invoice.ClientId = 1
		invoice.CompanyId = 1
		invoice.AddTagByName("Urgent")
		invoice.AddTagByName("Overdue")
		db.Create(invoice)
	}

	invoice2, ex2 := entity.NewInvoice(entity.InvoiceCreated)
	if ex2 == nil {
		invoice2.ClientId = 1
		invoice2.CompanyId = 1
		invoice2.AddTagByName("Urgent")
		invoice2.AddTagByName("CLS Project")
		//invoice2.AutoFillInvoiceNumber(2)
		dc := db.Create(invoice2)
		if dc.Error != nil {
			fmt.Println(dc.Error.Error())
		}
	} */

	//res := db.Model(&entity.TaxGroup{TaxGroupId: 2}).Association("Taxes").Count()
	//fmt.Println(res)

	/* tg := []entity.TaxGroup{}
	db.Debug().Preload(clause.Associations).Find(&tg)
	fmt.Println(len(tg)) */

	/* var taxGroup []entity.TaxGroup = []entity.TaxGroup{}
	db.Preload("Taxes").Find(&taxGroup)
	fmt.Println("List of Tax Groups")
	for _, tg := range taxGroup {
		fmt.Println(tg.ShortName)
		for _, tx := range *tg.Taxes {
			fmt.Printf("\t Tax: %s \t Perct: %f\n", tx.ShortName, tx.Percentage)
		}
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("List of Taxes")
	var tax []entity.Tax = []entity.Tax{}
	db.Find(&tax)
	for _, tg := range tax {
		fmt.Println(tg.ShortName, tg.TaxGroup)
	} */

}
