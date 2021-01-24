package model

// Community estructura de una comunidad
type Community struct {
	// Name nombre de una comunidad. Ej: EDteam
	Name string `json:"name" bson:"name"`
}

// Communities slice de comunidades
type Communities []Community

// Person estructura de una persona
type Person struct {
	ID string `json:"_id" bson:"_id,omitempty"`
	// Name nombre de la persona Ej: Alexys
	Name string `json:"name" bson:"name"`
	// Age edad de la persona Ej: 40
	Age uint8 `json:"age" bson:"age"`
	// Communities comunidades a las que pertenece una persona
	Communities Communities `json:"communities" bson:"communities"`
}

// Persons slice de personas
type Persons []Person