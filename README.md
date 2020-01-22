# rubik

(team of 2)

The goal of the project is to create a rubik's cube solver.

Given a shuffle, we need to find a sequence of move to solve the cube.
The goal was to optimize the number of move required to solve the cube.

We use the internation notation (F R U B L D for Front/Right/Up/Back/Left/Down)

We only have acces to 18 moves, `F` is a clockwise turn of the Front face, `F'` is
a counterclockwise tune of the Front face and `F2` correspond to `F' F'` or `F F`.
The modificators `'` and `2` can be use for any face.

You can choose from different algorithm to solve the cube:
- Default one : `Thistlewaite` algorithm. This one is the most efficient.
It require about 30 moves to solve any cube in less than 3 seconds.
- `CFOP` algorithm. It's a speedcubing algorithm. Really fast but require a lot of moves.
- `Bidirectionnal BFS`. Always return the optimal sequence of move to solve a cube. If it needs
more than 14 moves it is really long. We are using it in the `Thistlewaite` algorithm 
to go from G0 to G1, from G1 to G2 and from G2 to G4.
- `A*` algorithm. We have two not admissible heuristics.
The first one correspond to the number of move the `CFOP` algorithm will take to solve the cube.
The second one correspond to a BFS of depth 2 on the given state, then we evaluate the number of move
required by the `CFOP` algorithm to solve each state, and we take the minimum.

techno : Go

Programming method : Group theory, graph theory.

