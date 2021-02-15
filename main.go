package main

type TypeEntreprise struct {
	Name      string
	ID        int
	TaxRate   float64
	Mandatory []string
}

var AllTypes []TypeEntreprise

type Entreprise struct {
	Siret        string
	Denomination string
	Adress       string
	TypeId       int
	Revenue      float64
}

func initTypes() {
	AllTypes = []TypeEntreprise{
		TypeEntreprise{Name: "SAS", ID: 1, TaxRate: 0.33, Mandatory: []string{"Denomination", "Siret", "Adress"}},
		TypeEntreprise{Name: "Auto Entreprise", ID: 2, TaxRate: 0.25, Mandatory: []string{"Denomination", "Siret"}},
	}
}

func main() {

	initRoutes()
}
