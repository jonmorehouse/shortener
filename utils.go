package main

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var iterator [6]uint8 

func factorial(input uint8) uint64 {
  //var value uint64
  // factorial(3) = 3 * 2 * 1
  // this should be a tail recursive solution
  if input <= 1 {
    return 1
  }

  return uint64(input) * factorial(input - 1)
}

// permutations, order matters ... ie [0, 1, 2] != [0, 2, 1]
// n! / (n-k)!
func possiblePermutations(permutationSize uint8, setSize uint8) uint64 {
  permutations := factorial(setSize) / factorial(setSize - permutationSize)
  return permutations
}

func iterate() {
  for i := range iterator {
    // current value can be iterated
    if iterator[i] < uint8(len(characters) - 1) {
      iterator[i] = iterator[i] + 1
      return
    }
  
    // have maxed out the iterator. Need to start over
    if i == len(iterator) - 1 {
      // ERROR
      return
    }

    // increase the next piece of the array
    iterator[i] = 0
    iterator[i+1] = iterator[i+1] + 1
  }
}

func buildString() {
  // map character to index
  // bit shift index and then 

}




