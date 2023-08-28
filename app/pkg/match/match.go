package match

type MatchParam struct {
	Type      string
	condition interface{}
}

func Match(param MatchParam) {
	switch param.Type {
	case "US":
		USMatch(param.condition.(USCondition))
	case "UK":
		UKMatch(param.condition.(UKCondition))
	case "HK":
		HKMatch(param.condition.(HKCondition))
	case "AU":
		AUMatch(param.condition.(AUCondition))
	}

}
