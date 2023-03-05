package model

import "time"

type Customer struct {
	Id          int
	CustomerId  string // 客户编号
	Name        string // 客户名称
	OrderId     string //派单号
	PhoneNumber string //联系电话
	Sex         uint8  //性别: 1男 2女

	HospitalId   string //医院编号
	HospitalName string //医院名称
	OrderStatus  string //派单状态

	ServiceObject    string //服务项目
	PromotionChannel string //推广渠道

	NextContactTime        *time.Time //下次联系时间
	EstimatedOperationTime *time.Time //预计手术时间
	MeetingTime            *time.Time //见面时间
	SubmissionTime         *time.Time //成交时间

	CreatedAt *time.Time //创建时间
}
