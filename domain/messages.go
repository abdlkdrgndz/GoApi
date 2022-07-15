package domain

type Messages struct {
	Created  string "New customer created"
	Updated  string "Customer data updated."
	Deleted  string "Data was deleted."
	NotFound string "No registered user found at the specified e-mail address."
	NoCostumerFound string "No customer found."
}
