package carrera

type Carrera struct {
	ID          string `json:"_id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Cursos      []struct {
		ID     string `json:"_id"`
		Nombre string `json:"nombre"`
	} `json:"cursos"`
}