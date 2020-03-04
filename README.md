# rubik

The goal of the project is to solve a rubik cube with a minimum of movements.

The international notation: F R U B L D (Front/Right/Up/Back/Left/Down) is used.

`F` is a clockwise turn of the Front face.

`F'` is a counterclockwise tune of the Front face.

`F2` correspond to `F' F'` or `F F`.

The modificators `'` and `2` can be use for any face.

You can choose from different algorithm to solve the cube:
- `Thistlewaite`: 30 moves on average to solve any cube in less than 3 seconds.
- `CFOP`: It's a speedcubing algorithm. Really fast but require a lot of moves.
- `Bidirectionnal BFS`: Always return the optimal sequence of move to solve a cube, but can be very long for more than 14 moves.
- `A*`: With `CFOP` as heuristique.

Programming skill : Group theory, Graph theory.
