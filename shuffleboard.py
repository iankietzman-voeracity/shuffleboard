from itertools import combinations
import random
import time

tic = time.perf_counter_ns()

player_ids = [1,2,3,4,5,6,7,8,9]
players_per_team = 2
lanes = 2

def flatten_2D_matrix(matrix):
    """
    Accepts a two dimensional matrix of integers and returns it flattened to
    one dimension
    """
    return [item for row in matrix for item in row]

def filter_duplicated_player_team_assignments(matrix):
    """
    Accepts a two dimensional matrix of integers and returns True if 
    no duplicated player id integers are found in the flattened matrix 

    For use in a list comprehension
    """
    flat = flatten_2D_matrix(matrix)
    flat_set = set(flat)
    if len(flat) == len(flat_set):
        return True

def filter_duplicated_team_combos(combo, game):
    """
    Accepts a two dimensional matrix of two integers lists of player
    ids representing team assignments and a two item list of player 
    ids integers representing a recently assigned team assignment 
    and returns True if that team assignment is not found in the matrix

    For use in a list comprehension
    """
    for team in game:
        if team in combo:
            return False
    return True

assignments = []
i = 0

def build_rounds(byes, remaining_combos):
    """
    Recursively builds out team+round assignments, filtering out
    any eliminated combinations as it goes
    """
    global i
    combo = []
    rounds = len(byes)
    rounds -= 1
    if (rounds <=  -1):
        return assignments
    for combination in remaining_combos:
        i += 1
        flat = flatten_2D_matrix(combination)
        if not (flat.count(byes[0])):
            flat_set = set(flat)
            if len(flat) == len(flat_set):
                assignments.append(combination)
                combo = combination
                break
    remaining_combos = [remaining_combo for remaining_combo in remaining_combos if filter_duplicated_team_combos(remaining_combo, combo)]
    build_rounds(byes[1:], remaining_combos)

byes_per_round = 1
byes = player_ids.copy()
byes.reverse()
assignments = []

all_team_combos = list(combinations(combinations(player_ids, players_per_team), lanes * 2))
# print(len(combos)) #58905

remaining_combos = [combos for combos in all_team_combos if filter_duplicated_player_team_assignments(combos)]
# print(len(newlist)) #945

build_rounds(byes, remaining_combos)

for index, round_assignments in enumerate(assignments):
    round_assignments = list(round_assignments)
    print('--------------------------------')
    print(f'Round {index + 1}:')
    print('Teams: ', round_assignments)
    random.shuffle(round_assignments)
    print('Lanes: ', round_assignments)
    print('"Checksum":', sum(flatten_2D_matrix(round_assignments)) + byes[index] )

print('\nElapsed: ' + str((time.perf_counter_ns() - tic) / 1000000) + 'ns')
print(i)