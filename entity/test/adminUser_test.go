package test

import (
	"InvoiceGen/entity"
	"testing"
)

func TestNewAdminUser_Success(t *testing.T) {
	adminUser, ex := entity.NewAdminUser("Ethen", "ethen@gmail.com", "abc1234")
	if ex == nil {
		if adminUser.Name != "Ethen" {
			t.Error("Name property not set properly")
		}
		if adminUser.Email != "ethen@gmail.com" {
			t.Error("Email property not set properly")
		}
		if adminUser.Password != "abc1234" {
			t.Error("Password property not set properly")
		}
	} else {
		t.Error("Error occured: ", ex.Error())
	}
}
