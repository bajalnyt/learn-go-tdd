### learn-go-tdd

Learning Test Driven Development in Go by following : https://quii.gitbook.io/learn-go-with-tests/  
Each folder presents a new technique/concept.


It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. Every test has a cost.  
Go's built-in testing toolkit features a coverage tool. Whilst striving for 100% coverage should not be your end goal, the coverage tool can help identify areas of your code not covered by tests. If you have been strict with TDD, it's quite likely you'll have close to 100% coverage anyway.


* https://go.dev/blog/examples
* [Here](https://play.golang.org/p/bTrRmYfNYCp) is an example of slicing an array and how changing the slice affects the original array; but a "copy" of the slice will not affect the original array. [Another](https://play.golang.org/p/Poth8JS28sc) example of why it's a good idea to make a copy of a slice after slicing a very large slice.
