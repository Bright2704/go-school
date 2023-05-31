package entity

type Subject struct {
	ID uint `json: "primary_Key"`
}

func NewSubject(id uint) *Subject {
	return &Subject{
		ID: id,
	}
}
