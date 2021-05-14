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

	company, ex := NewCompany(
		"AdventureCode Ltd",
		"IT Plaza Dadabhai road, Malad West, Mumbai 400069, Maharashtra, India.",
		"info@domain.com",
		"+91 9833562829",
		"27BFKPS1782Q2XX",
	)
	if ex == nil {
		adminUser, ex := NewAdminUser(
			"Ethen Hunt",
			"admin@domain.com",
			"helloworld",
		)
		if ex == nil {
			adminUser.SetCompany(company)
			res = append(res, adminUser)
		}
		//res = append(res, company)
	}

	client, ex := NewClient(
		"Abc Inc.",
		"Office #20, Silver Star IT Park, Andheri West, Mumbai 400058, Maharashtra, India.",
		"abc@gmail.com",
		"+91 9877374937",
		"",
	)
	if ex == nil {
		res = append(res, client)
	}

	igst, ex := NewTax("Integrated GST", "IGST", 18)
	if ex == nil {
		igst.SetNewTaxGroup("Integrated GST", "IGST")
		res = append(res, igst)
	}

	gst, ex := NewTaxGroup("Goods & Service Tax", "GST")
	if ex == nil {
		sgst, ex := NewTax("State GST", "SGST", 9)
		if ex == nil {
			sgst.SetTaxGroup(gst)
			res = append(res, sgst)
		}
		cgst, ex := NewTax("Central GST", "CGST", 9)
		if ex == nil {
			cgst.SetTaxGroup(gst)
			res = append(res, cgst)
		}
	}

	return res
}
