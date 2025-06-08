package main

type Name struct {
	FirstName string // Exported field
	LastName  string // Exported field
}

type Email struct {
	EmailAddress string // Exported field
	EmailType    string // Exported field
}

type Person struct {
	Name   Name // Embedding Name struct
	Emails []Email // Embedding a slice of Email structs
}

func main() {
	// Example usage of the Person struct
	person := Person{
		Name: Name{
			FirstName: "Atul",
			LastName:  "Nath",
		},
		Emails: []Email{
			{EmailAddress: "atulnath@gmail.com", EmailType: "Personal"},
			{EmailAddress: "atulchandra@office.com", EmailType: "Work"},
		},	
	}
	// Print the person's details
	println("First Name:", person.Name.FirstName)
	println("Last Name:", person.Name.LastName)
	for _, email := range person.Emails {
		println("Email Address:", email.EmailAddress, "Type:", email.EmailType)
	}
}
