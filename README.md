### Project details
This is project work for the course DD1327 - Fundamentals of Computer Science. It is an implementation of a hash table in Go, last commit May 11, 2020.

# Hash table built on linked lists

### Hash table data structure
A hashtable is a common data structure that allows
efficient lookup, insertion and deletion of elements in the
average case wherein it takes constant time (the cost of resizing the
hash table itself is a amortized cost). 
Note that the worst case time complexity of all these operations are O(N) however.
This implementation resolves collisions by the use of linked lists,
which allows for a balance between speed and space efficiency that can be 
customized to ones preference.

### Installation (Not Available)

Once you have [installed Go][golang-install], run this command
to install the `hashtable` package:

    go get github.com/teojm/... //TODO, version 1.0.0 not yet available
    
### Roadmap
* The API of this library is frozen.
* Eventual licensing and release of version 1.0.0

Teo Jansson Minne – [teojm](https://gits-15.sys.kth.se/teojm)

