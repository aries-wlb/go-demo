package match

var usCategories = []USCategory{
	{
		Name: "top10",
		Conditions: USCondition{
			BaseCondition: BaseCondition{
				TOEFL: 110,
				IELTS: 8.0,
			},
			SAT: 1520,
			ACT: 33,
			AP:  3,
			GPA: 85.0,
		},
	},
	{
		Name: "top10 ~ top30",
		Conditions: USCondition{
			BaseCondition: BaseCondition{
				TOEFL: 105,
				IELTS: 7.5,
			},
			SAT: 1470,
			ACT: 32,
			AP:  3,
			GPA: 85.0,
		},
	},
	{
		Name: "top30 ~ top50",
		Conditions: USCondition{
			BaseCondition: BaseCondition{
				TOEFL: 95,
				IELTS: 7,
			},
			SAT: 1410,
			ACT: 30,
			AP:  3,
			GPA: 80.0,
		},
	},
	{
		Name: "top50 ~ top100",
		Conditions: USCondition{
			BaseCondition: BaseCondition{
				TOEFL: 80,
				IELTS: 6.5,
			},
			SAT: 0,
			ACT: 0,
			AP:  0,
			GPA: 80.0,
		},
	},
}

func matchUSCondition(c USCondition, c_ USCondition) bool {
	return (c.TOEFL >= c_.TOEFL || c.IELTS >= c_.IELTS) &&
		((c.SAT >= c_.SAT) || c.ACT >= c_.ACT) &&
		(c.AP >= 3) && (c.GPA >= c_.GPA)
}

func USMatch(condition USCondition) string {
	for _, category := range usCategories {
		matched := matchUSCondition(condition, category.Conditions)
		if matched {
			return category.Name
		}
	}
	return ""
}
