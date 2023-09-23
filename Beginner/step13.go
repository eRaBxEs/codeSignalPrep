/****

Let's imagine objects located on the canvas at a certain moment in time.
You are given an array of integer pairs centers representing the coordinates of those objects.
 Each object has a collision box - a square area around its center with a side equal to 2.
 Two objects are supposed to collide if their collision boxes have at least one common point.
 Calculate the number of object pairs that collide.

The example of the object located in coordinates [1, 1] can be represented as following:
Example of object with collision box
Where the green square is the collision box.

Note

Object collision boxes intersect if the distance in each coordinate between the object centers does not exceed 2.

If x1, y1 are coordinates of one object and x2, y2 are coordinates of the second object,
then the collision condition for them can be written in the form |x[j] - x[i]| <= 2 and |y[j] - y[i]| <= 2.

Example

For

centers = [
  [1, 1],
  [2, 2],
  [0, 4]
]
the output should be solution(centers) = 2.

Three objects

Explanation:

|x[1] - x[0]| = 1 and |y[1] - y[0]| = 1 both are <= 2, so the first two objects collide.
|x[2] - x[1]| = 2 and |y[2] - y[1]| = 2 both are <= 2, so the last two objects collide.
|x[2] - x[0]| = 1 which is <= 2, but and |y[2] - y[0]| = 3 which is > 2, so the first and the last object don't collide.
For

centers = [
  [1, 1],
  [2, 2],
  [0, 4],
  [1, 1]
]
the output should be solution(centers) = 4.

Four objects

Explanation:
The first three objects collide the same as in the previous example.
The last object is located in the same point as the first one, so:

It also collides with the second one and does not collide with the third one (+1 collision),
It collides with the first one (+1 collision).

***/

func solution(centers [][]int) int {
	// Initialize a variable to keep track of the number of collisions.
	collisions := 0

	// Iterate through each pair of objects.
	for i := 0; i < len(centers); i++ {
		for j := i + 1; j < len(centers); j++ {
			// Calculate the absolute differences in x and y coordinates.
			dx := abs(centers[i][0] - centers[j][0])
			dy := abs(centers[i][1] - centers[j][1])

			// Check if both dx and dy are less than or equal to 2.
			if dx <= 2 && dy <= 2 {
				collisions++
			}
		}
	}

	return collisions
}

// Helper function to calculate the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
