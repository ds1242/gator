package config

type ConfigStruct struct {
	DBUrl       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}
