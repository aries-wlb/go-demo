package match

var ukCategories = []UKCategory{
	{
		Name: "G5",
		Conditions: UKCondition{
			BaseCondition: BaseCondition{
				TOEFL: 100,
				IELTS: 7.0,
			},
			IB:     40,
			AP:     3,
			ALevel: aLevelScore.AScore * 3,
		},
	},
	{
		Name: "Top100",
		Conditions: UKCondition{
			BaseCondition: BaseCondition{
				TOEFL: 80,
				IELTS: 6.5,
			},
			IB:     33,
			AP:     0,
			ALevel: aLevelScore.AScore * 3,
		},
	},
}

func matchUKCondition(c UKCondition, c_ UKCondition) bool {
	return (c.TOEFL >= c_.TOEFL || c.IELTS >= c_.IELTS) &&
		((c.IB >= c_.IB) || c.ALevel >= c_.ALevel) &&
		(c.AP >= c_.AP)
}

func UKMatch(condition UKCondition) string {
	for _, category := range ukCategories {
		matched := matchUKCondition(condition, category.Conditions)
		if matched {
			return category.Name
		}
	}
	return ""
}
