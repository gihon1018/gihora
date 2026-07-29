package model

type Mod struct {
	Id     string `json:"id"`
	Remark string `json:"remark"`
}

func NewMod(id, remark string) Mod {
	return Mod{
		Id:     id,
		Remark: remark,
	}
}
