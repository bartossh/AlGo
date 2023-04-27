# AlGo

Playground repository containing collection of algorithms written in Go.

## Algorithms

### Sorting

1. [Bubble](https://github.com/bartossh/AlGo/blob/main/bubblesort/bubblesort.go) - bubble sort is the simplest sorting algorithm with O(n2) time complexity
2. [Insertion](https://github.com/bartossh/AlGo/blob/main/insertionsort/insertionsort.go) - insertion sort is simple sorting algorithm with  O(n) time complexity
3. [Bucket](https://github.com/bartossh/AlGo/blob/main/bucketsort/bucketsort.go) - bucket sort uses insertion sort for sorting bucket that are in increasing order, best time complexity is O(n+k), average O(n)
4. [CocktailShaker](https://github.com/bartossh/AlGo/blob/main/cocktailshakersort/cocktailshakersort.go) - cocktail shaker sort is swings back and forth to sort data with time complexity average of 0(n2) and best 0(n)
5. [Merge](https://github.com/bartossh/AlGo/blob/main/mergesort/mergesort.go) - merge sort is a divide and conquer sort algorithm with time complexity average of O(n log n)
6. [Shell](https://github.com/bartossh/AlGo/blob/main/shellsort/shellsort.go) - shell sort algorithm will sort values that are far apart from each other rather than adjacent, worst time complexity is O(n log^2 n), worst is O(n log n)
7. [Quick](https://github.com/bartossh/AlGo/blob/main/quicksort/quicksort.go) - quick sort algorithm is a divide and conquer sort algorithm with time complexity average of O(n\log n)
8. [Bartosz](https://github.com/bartossh/AlGo/blob/main/bartoszsort/bartoszsort.go) - sorting algorithm that is my playground for prototyping with sort algorithms 

### Path traversal

1. [Dijkstra](https://github.com/bartossh/AlGo/blob/main/dijkstra/dijkstra.go) - calculates shortest path without need for walking all available paths


### Data structure
1. [BTree](https://github.com/bartossh/AlGo/blob/main/btree/btree.go) - B-tree is a special type of self-balancing search tree in which each node can contain more than one key and can have more than two children [Animated example](https://www.cs.usfca.edu/~galles/visualization/BTree.html)
2. [Tries](https://github.com/bartossh/AlGo/blob/main/tries/tries.go) - Tries are a form of string-indexed look-up data structure, which is used to store a dictionary list of words that can be searched on in a manner that allows for efficient generation of completion lists.
3. [MaxSubArray](https://github.com/bartossh/AlGo/blob/main/maximumsubarray/maximumsubarray.go) - The maximum sum subarray problem, also known as the maximum segment sum problem, is the task of finding a contiguous subarray with the largest sum, within a given one-dimensional array A[1...n] of numbers
4. [Fifo Buffer](https://github.com/bartossh/AlGo/blob/main/fifobuf/fifobuf.go) - First in first out buffer is a buffer that queues data in a way that items are retrieved in the order they are added.
5. [Circular Buffer](https://github.com/bartossh/AlGo/blob/main/circularbuf/circularbuf.go) - circular buffer is a buffer allowing to move in circle using next and prev methods. It has some extensions methods such as seek and seek remove.
6. [Geo Math](https://github.com/bartossh/AlGo/blob/main/geomath/geo.math.go) - Geo Math algorithm is able to read geojson and validate if given location is within a geo polygon.

### Math

1. [BabyStepGiantStep](https://github.com/bartossh/AlGo/blob/main/babystepgiantstep/babystepgiantstep.go) - small step big step algorithm solves discrete logarithm problem of a^x = b (mod n) , with respect to gcd(a, n) == 1 with time complexity of O(sqrt(n))
2. [ExtendedEuclidean](https://github.com/bartossh/AlGo/blob/main/extendedeuclidean/extendedeuclidean.go) - in arithmetic and computer programming, the extended Euclidean algorithm is an extension to the Euclidean algorithm, and computes, in addition to the greatest common divisor (gcd) of integers a and b, also the coefficients of Bézout's identity, which are integers x and y such that
3. [Josephus Problem](https://github.com/bartossh/AlGo/blob/main/josephusproblem/josephusproblem.go) - Flavius Josephus, a Jewish-Roman historian from the first century, tells the story like this: A company of 40 soldiers, along with Josephus himself, were trapped in a cave by Roman soldiers during the Siege of Yodfat in 67 A.D. The Jewish soldiers chose to die rather, than surrender, so they devised a system to kill off each other until only one person remained. (That last person would be the only one required to die by their own hand.) All 41 people stood in a circle. The first soldier killed the man to his left, the next surviving soldier killed the man to his left, and so on. Josephus was among the last two men standing, "whether we must say it happened so by chance, or whether by the providence of God," and he convinced the other survivor to surrender rather than die. The solution is: number of people minus  the biggest power of 2 less then number of people times two plus one: f(N) = 2L + 1 where N =2^M + L and 0 <= L < 2^M)
4. [V8XorShift128](https://github.com/bartossh/AlGo/blob/main/v8XorShift128/v8XorShift128.go) - chrome v8 random function xor shift 128 

### Dynamic Programming

#### Memoization

1. [Fibonacci nth element](https://github.com/bartossh/AlGo/blob/main/fibonacci/fibonacci.go) - calculates n'th fibonacci sequence element recursively storing each n'th fibonacci value in the map to only calculate it once
2. [Coin Change](https://github.com/bartossh/AlGo/blob/main/coinchange/coinchange.go) - calculates the fewest number of coins that need to make up that amount
3. [Maximum Sub Array](https://github.com/bartossh/AlGo/blob/main/maximumsubarray/maximumsubarray.go) - finds sub array with the maximum sum and returns it sum
4. [Repetition Search](https://github.com/bartossh/AlGo/blob/main/repetitionsearch/repetitionsearch.go) - searches repetition in slice of integers
5. [Cracking RSA when p and q are close prime numbers](https://github.com/bartossh/AlGo/blob/main/crackrsa/crackrsa.go) - uses [Fermat's factorization method](https://en.wikipedia.org/wiki/Fermat%27s_factorization_method) to find p and q values assuming those are relatively close values of prime numbers. This algorithm may help to determine if you RSA algorithm pics values of p and q badly causing insecurity.

### Patterns

1. [Chained Calls](https://github.com/bartossh/AlGo/blob/main/chainedcalls/chainedcalls.go) - allows for restricted chained method calls. This prevents for calling API methods in a way that can cause unpredicted behavior.
