package main

import (
	"fmt"
	"reflect"
	"slices"
	"time"
)

type TeamAssignment []uint8

type RoundAssignment []TeamAssignment

type FullAssignment []RoundAssignment

func reverse(numbers []uint8) []uint8 {
	// Accepts a slice of uint8 and reverses it
	new := make([]uint8, len(numbers))
	copy(new, numbers)
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		new[i], new[j] = numbers[j], numbers[i]
	}
	return new
}

func combinations_2D(pool []uint8, number uint8) [][]uint8 {
	// Placeholder function for initial call to combinations_recursive_2D
	return combinations_recursive_2D(pool, number, nil, nil)
}

func combinations_recursive_2D(pool []uint8, number uint8, c []uint8, cc [][]uint8) [][]uint8 {
	// Accepts a slice of uint8 and a number of selections for each
	//combination and returns all unique combinations of the integers
	if len(pool) == 0 || number <= 0 {
		return cc
	}
	number--
	for i := range pool {
		r := make([]uint8, 2)
		copy(r, c)
		r[len(r)-int(number)-1] = pool[i]
		if number == 0 {
			cc = append(cc, r)
		}
		cc = combinations_recursive_2D(pool[i+1:], number, r, cc)
	}
	return cc
}

func combinations_3D(pool [][]uint8, number uint8) [][][]uint8 {
	// Placeholder function for initial call to combinations_recursive_3D
	return combinations_recursive_3D(pool, number, nil, nil)
}

func combinations_recursive_3D(pool [][]uint8, number uint8, c [][]uint8, cc [][][]uint8) [][][]uint8 {
	// Accepts a slice of []uint8 and a number of selections for each
	// combination and returns all unique combinations of the underlying
	// []uint8 slices
	if len(pool) == 0 || number <= 0 {
		return cc
	}
	number--
	for i := range pool {
		r := make([][]uint8, 4)
		copy(r, c)
		r[len(r)-int(number)-1] = pool[i]
		if number == 0 {
			cc = append(cc, r)
		}
		cc = combinations_recursive_3D(pool[i+1:], number, r, cc)
	}
	return cc
}

func GetByIndex[T any](slice []T, i int) T {
	return slice[i]
}

func flatten(lists [][]uint8) []uint8 {
	// Accepts a slice of []uint8 slices and returns
	// a flattened uint8 slice
	var res []uint8
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func filter_duplicated_player_team_assignments(team_combos [][][]uint8) [][][]uint8 {
	// Accepts a 3D matrix of team+round assignments and removes any entries
	// where the same player id is repeated multiple times for the same
	// assignment rount
	filtered_combos := [][][]uint8{}
	for _, combo := range team_combos {
		flat_combo := flatten(combo)
		slices.Sort(flat_combo)
		flat_length := len(flat_combo)
		deduped_combo := slices.Compact((flat_combo))
		final_length := len(deduped_combo)
		if flat_length == final_length {
			filtered_combos = append(filtered_combos, combo)
		}
	}
	return filtered_combos
}

func filter_duplicated_team_combos(remaining_combos [][][]uint8, last_combo [][]uint8) [][][]uint8 {
	// Accepts a 3D matrix of team+round assignments and the last selected
	// team+round assignment and removes any remaining combinations that
	// contain an already-used team combination
	filtered_combos := [][][]uint8{}
	for _, combo := range remaining_combos {
		found := [4]int{}
		for index, team := range last_combo {
			i := slices.IndexFunc(combo, func(c []uint8) bool {
				return reflect.DeepEqual(team, c)
			})
			found[index] = i
		}
		if reflect.DeepEqual(found, [4]int{-1, -1, -1, -1}) {
			filtered_combos = append(filtered_combos, combo)
		}
	}

	return filtered_combos
}

func build_rounds(byes []uint8, remaining_combos [][][]uint8, assignments [][][]uint8, i int) [][][]uint8 {
	// Recursively builds out team+round assignments, filtering out
	// any elimated combinations as it goes
	var found_combo [][]uint8
	rounds := len(byes)
	rounds -= 1
	if rounds <= -1 {
		return assignments
	}
	for _, combo := range remaining_combos {
		i += 1
		combo_copy := combo
		flat_combo := flatten(combo_copy)

		if slices.Index(flat_combo, byes[0]) == -1 {
			assignments = append(assignments, combo)
			found_combo = combo
			break
		}
	}
	remaining_combos = filter_duplicated_team_combos(remaining_combos, found_combo)
	assignments = build_rounds(byes[1:], remaining_combos, assignments, i)
	return assignments
}

func main() {
	start := time.Now()
	player_ids := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var players_per_team uint8 = 2
	var lanes uint8 = 2
	var byes_per_round uint8 = 1
	byes := reverse(player_ids[:])
	var assignments [][][]uint8

	all_team_combos := combinations_2D(player_ids, players_per_team)

	all_round_combos := combinations_3D(all_team_combos, lanes*2)

	remaining_combos := filter_duplicated_player_team_assignments(all_round_combos)

	fmt.Println("Player Ids: ", player_ids)
	fmt.Println("Players Per Team: ", players_per_team)
	fmt.Println("Lanes: ", lanes)
	fmt.Println("Rounds: ", len(byes))
	fmt.Println("Byes per round", byes_per_round)
	fmt.Println("Byes", byes)

	assignments = build_rounds(byes, remaining_combos, assignments, len(byes))

	for i, assignment := range assignments {
		var sum uint8 = 0
		for _, num := range flatten(assignment) {
			sum += num
		}

		fmt.Println("------------------------")
		fmt.Println("Round:", i+1)
		fmt.Println("Teams:", assignment)
		fmt.Println("Checksum:", sum+byes[i])
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %s%s", elapsed, "\n")

}
