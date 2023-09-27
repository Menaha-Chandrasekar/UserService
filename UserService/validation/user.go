package validation

import (
	"net/mail"
	"strconv"
)

// message User{  //CreateUser
//     string name=1;
//     string email=2;
//     string password=3;
//     int64 contact=4;
//     repeated string role=5;
// }

func ValidName(name string) bool{
	if len(name)<=30{
		return true
	}
	return false
}


func ValidEmail(email string) bool{
	_, err := mail.ParseAddress(email)
	if err!=nil{
		return false
	}
	return true
}

func ValidContact(num int64)bool{
	str := strconv.Itoa(int(num))
	if len(str)==10{
		return true
	}
	return false
}