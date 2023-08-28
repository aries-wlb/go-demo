package auth_common

import (
	"patrick.com/abroad/app/domain/user"
	"patrick.com/abroad/app/repository"
)

type AuthCommonHandler struct {
	Impl *repository.RepoImpl
}

type MatchReq struct {
	SchoolName         string                    `json:"school_name"`
	SchoolType         string                    `json:"school_type"`
	Major              string                    `json:"major"`
	GPA                float64                   `json:"gpa"`
	LanguageAchi       string                    `json:"language_achi"`
	AcademicExperience []user.AcademicExperience `json:"academic_experience"`
	IntentRegion       string                    `json:"intent_region"`
	IntentMajor        string                    `json:"intent_major"`
	OtherDetails       string                    `json:"other_details"`
}
