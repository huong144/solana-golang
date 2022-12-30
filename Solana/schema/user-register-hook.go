package schema

import "gorm.io/gorm"

type UserRegisterWebhook struct {
	gorm.Model
	UserId         int    `gorm:"column:user_id;index:idx_user_register,unique"`
	RegistrationId int    `gorm:"column:registration_id;index:idx_user_register,unique"`
	Address        string `gorm:"column:address;index:idx_user_register,unique"`
	WebhookAddress string `gorm:"column:webhook_address;index:idx_user_register,unique"`
}
