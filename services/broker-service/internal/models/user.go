package models

import "time"


// Gender type to hold gender constants
type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
	Other  Gender = "Other"
)


type Tokens struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserType  string `json:"user_type"`
}
// Address struct to hold address details
type Address struct {
	Street     string `json:"street" bson:"street"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	PostalCode string `json:"postal_code" bson:"postal_code"`
	Country    string `json:"country" bson:"country"`
}

// Patient struct to hold patient details
type Patient struct {
	ID                int       `json:"id" bson:"id"`
	FirstName         string    `json:"first_name" bson:"first_name"`
	LastName          string    `json:"last_name" bson:"last_name"`
	DateOfBirth       time.Time `json:"date_of_birth" bson:"date_of_birth"`
	Gender            Gender    `json:"gender" bson:"gender"`
	ContactNumber     string    `json:"contact_number" bson:"contact_number"`
	Email             string    `json:"email" bson:"email"`
	Address           Address   `json:"address" bson:"address"`
	MedicalHistory    string    `json:"medical_history" bson:"medical_history"`
	Allergies         []string  `json:"allergies" bson:"allergies"`
	CurrentMedications []string `json:"current_medications" bson:"current_medications"`
	UserType  string `json:"user_type"`
}

// Doctor struct to hold doctor details
type Doctor struct {
	ID                int              `json:"id" bson:"id"`
	FirstName         string           `json:"first_name" bson:"first_name"`
	LastName          string           `json:"last_name" bson:"last_name"`
	Specialization    string           `json:"specialization" bson:"specialization"`
	YearsOfExperience int              `json:"years_of_experience" bson:"years_of_experience"`
	ContactNumber     string           `json:"contact_number" bson:"contact_number"`
	Email             string           `json:"email" bson:"email"`
	Address           Address          `json:"address" bson:"address"`
	LicenseNumber     string           `json:"license_number" bson:"license_number"`
	ClinicHours       ClinicalHours `json:"clinic_hours" bson:"clinic_hours"` // e.g., {"Monday": "9am-5pm", "Tuesday": "9am-5pm", ...}
	UserType  string `json:"user_type"`
}

type Admin struct {
	Name string
	Email string
	Password string
	UserType  string `json:"user_type"`
}

 type TimeSlot struct {
	Start string
	End string
 }

 type WeekDay string

 const (
	Monday WeekDay = "monday"
	Tuesday WeekDay = "tuesday"
	Wednesday WeekDay = "wednesday"
	Thursday WeekDay = "thursday"
	Friday WeekDay = "friday"
	Saturdaty WeekDay = "saturday"
	Sunday WeekDay = "sunday"
 )

 type ClinicalHours map[WeekDay][]TimeSlot