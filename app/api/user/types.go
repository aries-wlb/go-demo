package user

import (
	"time"

	"patrick.com/abroad/app/domain/user"
)

type UserResp struct {
	Account              string                    `json:"account"`
	AccountName          string                    `json:"account_name"`
	PhoneNumber          string                    `json:"phone_number"`
	RoleId               int                       `json:"role_id"`
	UserId               int                       `json:"user_id"`
	GPA                  float64                   `json:"gpa"`
	SchoolName           string                    `json:"school_name"`
	SchoolType           string                    `json:"school_type"`
	Major                string                    `json:"major"`
	LanguageAchi         string                    `json:"language_achi"`
	AcademicExperience   []user.AcademicExperience `json:"academic_experience"`
	IntentRegion         string                    `json:"intent_region"`
	IntentMajor          string                    `json:"intent_major"`
	OtherDetails         string                    `json:"other_details"`
	UserName             string                    `json:"user_name"`
	AvatarUrl            string                    `json:"avatar_url"`
	Gender               string                    `json:"gender"`
	Location             string                    `json:"location"`
	WechatID             string                    `json:"wechat_id"`
	Email                string                    `json:"email"`
	PersonalIntroduction string                    `json:"personal_introduction"`
	CreatedAt            time.Time                 `json:"created_at"`
	LastLogin            time.Time                 `json:"last_login"`
}

type UserUpdateReq struct {
	AccountName          string                    `json:"account_name"`
	UserName             string                    `json:"user_name"`
	Email                string                    `json:"email"`
	PhoneNumber          string                    `json:"phone_number"`
	WechatID             string                    `json:"wechat_id"`
	Location             string                    `json:"location"`
	SchoolName           string                    `json:"school_name"`
	SchoolType           string                    `json:"school_type"`
	Major                string                    `json:"major"`
	GPA                  float64                   `json:"gpa"`
	LanguageAchi         string                    `json:"language_achi"`
	AcademicExperience   []user.AcademicExperience `json:"academic_experience"`
	IntentRegion         string                    `json:"intent_region"`
	IntentMajor          string                    `json:"intent_major"`
	OtherDetails         string                    `json:"other_details"`
	PersonalIntroduction string                    `json:"personal_introduction"`
	Gender               string                    `json:"gender"`
}
