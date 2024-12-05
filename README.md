# Shuffleboard
## A real life algrorithm 

As commissioner of Winter League, the task has fallen to me to create a round robin bracket for our shuffleboard night. The main goal is to have every player play with every other player one time. There are some known constraints to help guide us. Nine players will be in attendance, which enables a fairly easy decision to reserve two lanes, as two lanes of two teams of two people equals eight players per round. This, of course, forces there to be one player per round with a buy. Luckily, nine players in attendance works very well in enabling us to use the minimum number of rounds. In general, anything where players % 4 == 0 or players % 4 == 1 works smoothly and lets us fill every available spot in every round while letting us have zero or one player byes per round. Extending this algorithm to the other two possibilities with regards to % 4 would add some complication, lower lane usage efficienty, and result in multiple buys per player and open lanes.

The (extremely, as I will lay out) naive approach to this problem is to generate all possible combinations of unique team player combinations, which in our specific case is 36 combinations long. From this we can then generate all possible round combinations of four combinations of team combinations. In our case this gives us 58905 combinations. From there we can generate all combinations of number of rounds from our round combinations and compare them to our known constraints to find a winner. But for our numbers this turns out to be...

`C(n,r) = n! / r!(n - r)! = 58905! / 9!(58905 - 9)! = 2.35e37`

Seems just slightly impractical to try to brute force it.

First we add a filter step to remove from the 58905 any of the round combinations that contain a duplicated player id, as a player can't be on two teams at once. After that, we are left with 945 combinations:

`C(n,r) = n! / r!(n - r)! = 945! / 9!(945 - 9)! = 1.59e21`

A marked improvement, but still impractical. 

Next we write a recursive function to build out our rounds. `build_rounds` is a greedy function that finds the first new unique combination of team assignments, breaks from the loop examining the remaining combinations upon finding one, then filters out any round combinations that contain any of our most recently selected team combinations on each call. This approach enbables us to iteratively whittle down the pool of round X team combinations as we go along. Gathering some empirical data indicates our algorithm only has to peer into our iteratively shrinking data pool forty times to complete the process. Not too shabby compared to trying to examine all possible 2.35e37 combinations.

And thus we have essentially improved our algorithm from the naive approach time complexity of `O(n!)` to the finalized version of `O(log n)`. 