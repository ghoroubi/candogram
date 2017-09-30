package main

import (
	"fmt"
	"os"

	"github.com/ghoroubi/mt"
	"github.com/ghoroubi/mtx"

	"github.com/gin-gonic/gin"
)

func Init() {
	var err error
	//MainMTP init and Connect
	m, err = mtprotox.NewMTProto(".telegram_go")
	if err != nil {
		fmt.Printf("Create failed: %s\n", err)
		os.Exit(2)
	}

	//Connecting to telegram Nearest DataCenter
	err = m.Connect()
	if err != nil {
		fmt.Printf("Connect failed: %s\n", err)
		os.Exit(2)
	}
	//EhMTP init and Connect
	nm, err = mtproto.NewMTProto(".telegram_go", "", 0)
	if err != nil {
		fmt.Printf("Create failed: %s\n", err)
		os.Exit(2)
	}
	err = nm.Connect()
	if err != nil {
		fmt.Printf("Connect failed: %s\n", err)
		os.Exit(2)
	}

}
func NewContactController(c *gin.Context) {
	var id int32
	Init()
	phone := c.Param("phone")
	/*switch CheckContactExisting(phone) {
	case -1,false:
		id=NewContact(phone)
	default:
		CheckContactExisting(phone)
	}*/
	if id, _ = CheckContactExisting(phone,0); id == -1 {
		id = NewContact(phone)
	}
	fmt.Println("Phone Number Entered:", phone)
	c.JSON(200, gin.H{"New Contact ID": id})
}
func NewContact(phone string) int32 {
	//var err error
	var userId int32
	firstName, lastName := phone, "new Contact"
	contact := mtproto.Contact{Firstname: firstName, Lastname: lastName, Phone: phone}
	nm.Contacts_ImportContacts([]mtproto.TL{contact.GetInputContact()})
	result := nm.Contacts_GetContacts()
	for _, v := range result {
		if phone == v.Phone {
			userId = v.ID
			fmt.Println(v.ID)
		}
	}
	return userId
}
func SendMessageController(c *gin.Context) {
	var err error
	var userId int32
	//Init()
	phone := c.Query("phone")
	message := c.Query("message")
	fmt.Println("Phone:", phone, " Message:", message, "Imported UserID:", userId)
	if userId, _ = CheckContactExisting(phone, 1); userId == -1 {
		userId = NewContact(phone)
	}
	fmt.Println("Phone:", phone, " Message:", message, "Imported UserID:", userId)
	err = m.SendMsg(userId, message)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		c.BindJSON(gin.H{"status": "Send"})

	}
}
func CheckContactExisting(phone string, protocol int) (int32, bool) {
	//protocol : ehMTPROTO =1   MTPROTO =0
	var userID int32 = -1
	if protocol == 1 {
		contacts := nm.Contacts_GetContacts()
		for _, v := range contacts {
			if v.Phone == phone {
				userID = v.ID
				return userID, true
			}
		}
		return userID, false
	}else{
		userID=-1
m.GetUserId(phone)
	}
	return userID,false
}