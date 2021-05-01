package test

import (
	"InvoiceGen/entity"
	"testing"
)

func TestNewClient_Success(t *testing.T) {
	client, ex := entity.NewClient("Nirav", "Mumbai Maharashtra India", "nirav@ventureoneit.com", "9833500846", "asdfgh324")
	if ex == nil {
		if client.Name != "Nirav" {
			t.Error("Name property not set properly")
		}
		if client.Address != "Mumbai Maharashtra India" {
			t.Error("Address property not set properly")
		}
		if client.Email != "nirav@ventureoneit.com" {
			t.Error("Email property not set properly")
		}
		if client.ContactNumber != "9833500846" {
			t.Error("Contact Number property not set properly")
		}
		if client.GSTNumber != "asdfgh324" {
			t.Error("GSTNumber property not set properly")
		}

	} else {
		t.Error("Error occured: ", ex.Error())
	}
}
