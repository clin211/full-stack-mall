package consignee

type AddConsigneeReq struct {
	UserId    int    `json:"user_id,omitempty"`
	IsDefault int    `json:"is_default" v:"required#是否默认地址必传"`
	Name      string `json:"name" v:"required#收货人姓名必传"`
	Phone     string `json:"phone" v:"required#收货人手机号必传"`
	Province  string `json:"province" v:"required#省份必传"`
	City      string `json:"city" v:"required#城市必传"`
	Town      string `json:"town" v:"required#县区必传"`
	Street    string `json:"street" v:"required#乡镇街道必传"`
	Detail    string `json:"detail" v:"required#详细地址必传"`
}

type UpdateConsigneeReq struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id,omitempty"`
	IsDefault int    `json:"is_default" v:"required#是否默认地址必传"`
	Name      string `json:"name" v:"required#收货人姓名必传"`
	Phone     string `json:"phone" v:"required#收货人手机号必传"`
	Province  string `json:"province" v:"required#省份必传"`
	City      string `json:"city" v:"required#城市必传"`
	Town      string `json:"town" v:"required#县区必传"`
	Street    string `json:"street" v:"required#乡镇街道必传"`
	Detail    string `json:"detail" v:"required#详细地址必传"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Page  int    `json:"page" v:"required#请合理输入页数"`
	Limit int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ListConsigneeRes struct {
	Count int                 `json:"count"`
	List  []*ListConsigneeSql `json:"list"`
}

type ListConsigneeSql struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id,omitempty"`
	IsDefault int    `json:"is_default" v:"required#是否默认地址必传"`
	Name      string `json:"name" v:"required#收货人姓名必传"`
	Phone     string `json:"phone" v:"required#收货人手机号必传"`
	Province  string `json:"province" v:"required#省份必传"`
	City      string `json:"city" v:"required#城市必传"`
	Town      string `json:"town" v:"required#县区必传"`
	Street    string `json:"street" v:"required#乡镇街道必传"`
	Detail    string `json:"detail" v:"required#详细地址必传"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DetailReq struct {
	Id int `json:"id"`
}
