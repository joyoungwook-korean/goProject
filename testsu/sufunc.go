package testsu

func Game(su, game_su int) (string, bool) {
	if su > game_su {
		return "min", false
	} else if su < game_su {
		return "better than", false
	} else {
		return "con", true
	}
}
