
// https://www.hackerrank.com/challenges/cats-and-a-mouse/problem

function catAndMouse(x, y, z) {
  const xDiff = Math.abs(x - z)
  const yDiff = Math.abs(y - z)
  return xDiff === yDiff ? "Mouse C" : xDiff < yDiff ? "Cat A" : "Cat B"
}