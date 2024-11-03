package models

type Person struct {
	ID              string `json:"id"`
	FirstName        string`json:"first_name"`
	MiddleName       string `json:"middle_name"`
	Surname          string `json:"surname"`
	Age              int    `json:"age"`
	Gender           string `json:"gender"`
	LastSeenLocation string `json:"lastseen_location"`
	LastSeenDate     string `json:"lastseen_date"`
	ContactPerson    string `json:"contact_person"`
	ContactPhone     string `json:"contact_phone"`
	ContactEmail     string `json:"contact_email"`
	ImageURL         string  `json:"image_url"`
} 


type ImageUploadResponse struct {
    ImageURL string `json:"image_url"`
    Success  bool   `json:"success"`
    Error    string `json:"error,omitempty"`
}