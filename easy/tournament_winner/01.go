package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	winner := ""
	score := make(map[string]int)

	for i, teams := range competitions {
		tw := "";

		if results[i] == HOME_TEAM_WON {
			tw = teams[0];
		} else {
			tw = teams[1];
		}

		cval, ok := score[tw];
		if !ok { cval = 0; }
		score[tw] = cval + 3;

		wval, ok := score[winner];
		if !ok { wval = 0; }

		if score[tw] > wval {
			winner = tw;
		}
	}

	return winner;
}
