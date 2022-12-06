# Grokking algorithms exercises

## Intro to algorithms

Key terms:
- Big-O notation
- linear time
- log time
- constant time

Common Big-O run times

When determining how fast an algorithm is it isn't enough to know how long an algorithm took to run. You need to figure out how the run time increases as the size increases, which is what Big-O notation is used for. Big-O notation tells you how fast an algorithm is by letting you compare the number of operations an algorithm will take by expressing them with a mathematical expression.

- `O(1)`, also known as *constant time*  
- `O(log n)`, also known as *log time*. Example: [Binary search](https://go.dev/play/p/_8l_9M-JHbo).  
- `O(n)`, also known as linear time. Example: Simple search.  
- `O(log n)` is faster than O(n), but it gets a lot faster as the list of items you’re searching grows  
- `O(n * log n)`. Example: A fast sorting algorithm, like quicksort  
- `O(n2)`. Example: A slow sorting algorithm, like *[selection sort](https://go.dev/play/p/qIKqlckz-gZ)*  
- `O(n!)`, also known as factorial time. Example: A really slow algorithm, like the traveling salesperson  

## Select sort

- [Binary search](https://go.dev/play/p/FCC3NEn5SQ-)
- [Selection sort](https://go.dev/play/p/tH96VhF7ync)

## Recursion

Key terms:
- base case
- recursive case
- call stack

Algorithms:
- [Factioral Recursion](https://play.golang.com/p/Ta9vPb-SC0N)
- [Count items in a slice (recursion)](https://go.dev/play/p/BcpvpAq1MKP)
- [Find maximum number (recursion)](https://go.dev/play/p/KLy0MyZ9K5W)
- [Binary search (recursion)](https://go.dev/play/p/7qM0WKge6O_x)
- [Sum (recursion)](https://go.dev/play/p/Nnvc4BXF9de)

## Quicksort

Key terms:
- Divide & Conquer
- Pivot
- Partitioning

Algorithms:
- [Quicksort divide and conquer](https://go.dev/play/p/Cx7Mq9svxTa)


## Breadth-first Search

Key terms:
- queue
- enqueue
- dequeue
- FIFO data structure
- LIFO data structure (stack)
- graph (data structure)
- direct graph
- indirect graph

Useful links
- https://codesource.io/how-to-implement-the-queue-in-golang

Algorithms:
- [Graph](https://go.dev/play/p/NMssEtyBWbX)

https://go.dev/play/p/bCz8OzwRAFp

