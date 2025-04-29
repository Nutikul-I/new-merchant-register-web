package model

type AuditLog struct {
	MerchantId      string `json:"MerchantId"`
	UserId          string `json:"UserId"`
	ActionName      string `json:"ActionName"`
	EntityType      string `json:"EntityType"`
	EntityId        string `json:"EntityId"`
	PropName        string `json:"PropName"`
	OldValue        string `json:"OldValue"`
	NewValue        string `json:"NewValue"`
	Remark          string `json:"Remark"`
	MenuName        string `json:"MenuName"`
	ClientIpAddress string `json:"ClientIpAddress"`
	CreatedAt       string `json:"CreatedAt"`
}

var CREATE_SQL_AUDITLOG = `INSERT INTO AuditLog (MerchantId, UserId, ActionName, EntityType, EntityId, PropName, OldValue, NewValue, Remark, MenuName, ClientIpAddress, CreatedAt, Channel) 
VALUES(@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, 'ADMIN_APP');`
