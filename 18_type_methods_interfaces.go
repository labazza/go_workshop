
package main

import "fmt"

func main() {

  // make a new type MyInt that is an integer

  type MyInt int

  // attach a method to MyInt to square a number

  func (n MyInt) sqr() MyInt {
	  return n * n
  }

  // make a new MyInt-type variable
  // called "number" and set it to 5

  var number MyInt = 5

  // and now the sqr() method can be used

  var square = number.sqr()

  // the value of square is now 25

  /*
  Suppose you've created types called cat, dog and bird, and each has a method called age() that returns the age of the animal.
  If you want to add the ages of all animals in one operation, you can define an interface like this:
  */

  type animal interface {
	  age() int
  }
}
