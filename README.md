# board-games-helper
Do you want to impress your friends with superhuman skills? This repository is for you! Use it wisely.
### Usage on mobile
Copy-paste the code you want to use in [repl.it](https://repl.it/languages/go) and click *Run*.
# Scotland Yard Helper
### Where's Mister X?
Predict the position of Mister X from a known point and an arbitrary number of tickets among taxi, bus, metro, black.

Use 'point ticket1 ticket2 ...' where *point* is the last number where Mister X was seen; *ticket* is a char among taxi (t), bus (b), metro (m), black (?) that Mister X used from the last number.

You will receive in output all the possible positions of Mister X.

Example:
Input: '1 ? m'
Output: 'from 1 with black ticket to [8, 9, 58, 46] and metro ticket to'
