package models

type ClubTranslation struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ClubTranslations struct {
	Data []ClubTranslation `json:"data"`
}
