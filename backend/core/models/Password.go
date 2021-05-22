package models

type Password struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

func (p *Password) IsValid() bool {
	if p.Password == p.NewPassword {
		return false
	}
	return true
}
