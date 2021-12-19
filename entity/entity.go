package entity

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var AllModels = []interface{}{
	&Currency{},
	&AdminUser{},
	&Client{},
	&Company{},
	&Invoice{},
	&InvoiceItem{},
	&Tax{},
	&TaxGroup{},
}

type UUID = uuid.UUID

func NewUUID() UUID {
	return UUID(uuid.New())
}

func StringToUUID(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	return UUID(id), err
}

type DefaultStruct struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CopyProperties(source interface{}, destination interface{}) error {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(source)
	if err != nil {
		return err
	}
	err = gob.NewDecoder(&buf).Decode(destination)
	if err != nil {
		return err
	}
	return err
}

//https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func GenerateRandomStringOfSize(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func GenerateDefaultData() []interface{} {
	res := []interface{}{}

	curINR, ex := NewCurrency("INR", 1)
	if ex == nil {
		res = append(res, curINR)
	}
	curUSD, ex := NewCurrency("USD", 0.014)
	if ex == nil {
		res = append(res, curUSD)
	}
	curCAD, ex := NewCurrency("CAD", 0.017)
	if ex == nil {
		res = append(res, curCAD)
	}

	company, ex := NewCompany(
		"AdventureCode Ltd",
		"IT Plaza Dadabhai road, Malad West, Mumbai 400069, Maharashtra, India.",
		"info@domain.com",
		"+91 9833562829",
		"27BFKPS1782Q2XX",
		curINR,
	)
	if ex == nil {
		res = append(res, company)
	}

	adminUser, ex := NewAdminUser(
		"Ethen Hunt",
		"admin@domain.com",
		"helloworld",
		company,
	)
	if ex == nil {
		res = append(res, adminUser)
	}

	client, ex := NewClient(
		"Abc Inc.",
		"Office #20, Silver Star IT Park, Andheri West, Mumbai 400058, Maharashtra, India.",
		"abc@gmail.com",
		"+91 9877374937",
		"India",
		"27BFKPS1283Q2AA",
		curINR,
		company,
	)
	if ex == nil {
		res = append(res, client)
	}

	client2, ex := NewClient(
		"John & Sons",
		"Office #101, New York, New York 10012, United States of America.",
		"jsons@gmail.com",
		"+01 555 192-2028",
		"USA",
		"",
		curUSD,
		company,
	)
	if ex == nil {
		res = append(res, client2)
	}

	client3, ex := NewClient(
		"XYZ Incorporated",
		"Office #500, Ottawa 10012, Canada.",
		"xyz@domain.com",
		"+01 551 192-2028",
		"CAD",
		"",
		curCAD,
		company,
	)
	if ex == nil {
		res = append(res, client3)
	}

	igstGroup, ex := NewTaxGroup("Integrated GST", "IGST")
	if ex == nil {
		igst, ex := NewTax("Integrated GST", "IGST", 18.0)
		if ex == nil {
			igstGroup.AddTax(igst)
			//res = append(res, igst)
		}
		res = append(res, igstGroup)
	}

	gstGroup, ex := NewTaxGroup("Goods & Service Tax", "GST")
	if ex == nil {

		sgst, ex := NewTax("State GST", "SGST", 9.0)
		if ex == nil {
			gstGroup.AddTax(sgst)
			//res = append(res, sgst)
		}
		cgst, ex := NewTax("Central GST", "CGST", 9.0)
		if ex == nil {
			gstGroup.AddTax(cgst)
			//res = append(res, cgst)
		}

		res = append(res, gstGroup)
	}

	invoice, ex := NewInvoice(InvoicePending, adminUser)
	if ex == nil {
		invItems := []*InvoiceItem{}
		invoiceItem, ex := NewInvoiceItem("Development services", 1, 30250.12)
		if ex == nil {
			invItems = append(invItems, invoiceItem)
		}
		invoiceItem, ex = NewInvoiceItem("IT consultancy services", 2, 12500)
		if ex == nil {
			invItems = append(invItems, invoiceItem)
		}

		invoice.AutoFillInvoiceNumber(0)
		invoice.SetDetails(company, client, curINR, gstGroup)
		invoice.SetInvoiceItems(invItems)

		res = append(res, invoice)
	}

	invoice2, ex := NewInvoice(InvoicePaid, adminUser)
	if ex == nil {
		invItems := []*InvoiceItem{}
		invoiceItem, ex := NewInvoiceItem("Responsive website design", 1, 42500)
		if ex == nil {
			invoiceItem.SetClassificationCode("998314")
			invoiceItem.SetNote("Website: www.domain.com")
			invItems = append(invItems, invoiceItem)
		}
		invoiceItem, ex = NewInvoiceItem("Website performance review and optimization", 1, 22000)
		if ex == nil {
			invoiceItem.SetClassificationCode("998314")
			invItems = append(invItems, invoiceItem)
		}

		invoice2.AutoFillInvoiceNumber(1)
		invoice2.SetDetails(company, client, curUSD, nil)
		invoice2.SetInvoiceItems(invItems)

		res = append(res, invoice2)
	}

	invoice3, ex := NewInvoice(InvoiceCancelled, adminUser)
	if ex == nil {
		invItems := []*InvoiceItem{}
		invoiceItem, ex := NewInvoiceItem("Hosting services", 3, 2500)
		if ex == nil {
			invoiceItem.SetClassificationCode("998315")
			invoiceItem.SetNote("Websites: www.domain.com, abc.com, xyz.co.net")
			invItems = append(invItems, invoiceItem)
		}

		invoice3.AutoFillInvoiceNumber(2)
		invoice3.SetDetails(company, client, curINR, igstGroup)
		invoice3.SetInvoiceItems(invItems)

		res = append(res, invoice3)
	}

	invoice4, ex := NewInvoice(InvoicePaid, adminUser)
	if ex == nil {
		invItems := []*InvoiceItem{}
		invoiceItem, ex := NewInvoiceItem("Hosting services", 3, 2500)
		if ex == nil {
			invoiceItem.SetClassificationCode("998315")
			invoiceItem.SetNote("Websites: www.domain.com, abc.com, xyz.co.net")
			invItems = append(invItems, invoiceItem)
		}

		invoice4.SetInvoiceNumber("1005211")
		invoice4.SetDetails(company, client2, curUSD, igstGroup)
		invoice4.SetInvoiceItems(invItems)

		res = append(res, invoice4)
	}

	return res
}
