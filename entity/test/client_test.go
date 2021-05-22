package test

import (
	"InvoiceGen/entity"
	"testing"
)

func TestNewClient_Success(t *testing.T) {
	client, ex := entity.NewClient("Ethen", "Mumbai Maharashtra India", "ethen@gmail.com", "9839283939", "asdfgh324")
	if ex == nil {
		if client.Name != "Ethen" {
			t.Error("Name property not set properly")
		}
		if client.Address != "Mumbai Maharashtra India" {
			t.Error("Address property not set properly")
		}
		if client.Email != "ethen@gmail.com" {
			t.Error("Email property not set properly")
		}
		if client.ContactNumber != "9839283939" {
			t.Error("Contact Number property not set properly")
		}
		if client.GSTNumber != "asdfgh324" {
			t.Error("GSTNumber property not set properly")
		}

	} else {
		t.Error("Error occured: ", ex.Error())
	}
}
