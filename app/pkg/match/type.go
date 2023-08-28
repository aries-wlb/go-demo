package match

type USCategory struct {
	Name       string
	Conditions USCondition
}

type UKCategory struct {
	Name       string
	Conditions UKCondition
}

type BaseCondition struct {
	TOEFL int
	IELTS float64
}

type USCondition struct {
	BaseCondition
	SAT int
	ACT int
	AP  int
	GPA float64
}

type UKCondition struct {
	BaseCondition
	AP     int
	IB     int
	ALevel int
}

type ALevelScore struct {
	AScore  int
	BScore  int
	APScore int
}

type AUCategory struct {
	Name       string
	Conditions AUCondition
}

type AUCondition struct {
	BaseCondition
	YBorSeven bool
}

type HKCategory struct {
	Name       string
	Conditions HKCondition
}

type HKCondition struct {
	IELTS  float64
	ALevel int
}
