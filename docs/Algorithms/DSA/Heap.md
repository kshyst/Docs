# Heap

## Max Heap

Every parent node is bigger than its child nodes

## Min Heap

Every parent node is smaller than its child nodes

## Heap Sort

works in O(log n)

### Max Heap in python

```python
# in this heap implementation, index 0 is not used to simplify parent/child calculations.
# the index 0 is set to sys.maxsize to ensure it is always larger than any other element.
# then we define root = 1 as the starting point of the heap.


import sys

class MaxHeap:
    def __init__(self, cap):
        self.cap = cap
        self.n = 0
        self.a = [0] * (cap + 1)
        self.a[0] = sys.maxsize
        self.root = 1

    def parent(self, i):
        return i // 2

    def left(self, i):
        return 2 * i

    def right(self, i):
        return 2 * i + 1

    def isLeaf(self, i):
        return i > (self.n // 2) and i <= self.n

    def swap(self, i, j):
        self.a[i], self.a[j] = self.a[j], self.a[i]

    def maxHeapify(self, i):
        if not self.isLeaf(i):
            largest = i
            if self.left(i) <= self.n and self.a[i] < self.a[self.left(i)]:
                largest = self.left(i)
            if self.right(i) <= self.n and self.a[largest] < self.a[self.right(i)]:
                largest = self.right(i)
            if largest != i:
                self.swap(i, largest)
                self.maxHeapify(largest)

    def insert(self, val):
        if self.n >= self.cap:
            return
        self.n += 1
        self.a[self.n] = val
        i = self.n
        while self.a[i] > self.a[self.parent(i)]:
            self.swap(i, self.parent(i))
            i = self.parent(i)

    def extractMax(self):
        if self.n == 0:
            return None
        max_val = self.a[self.root]
        self.a[self.root] = self.a[self.n]
        self.n -= 1
        self.maxHeapify(self.root)
        return max_val

    def printHeap(self):
        for i in range(1, (self.n // 2) + 1):
            print(f"PARENT: {self.a[i]}", end=" ")
            if self.left(i) <= self.n:
                print(f"LEFT: {self.a[self.left(i)]}", end=" ")
            if self.right(i) <= self.n:
                print(f"RIGHT: {self.a[self.right(i)]}", end=" ")
            print()

# Example
if __name__ == "__main__":
    print("The maxHeap is:")
    h = MaxHeap(15)
    vals = [5, 3, 17, 10, 84, 19, 6, 22, 9]
    for val in vals:
        h.insert(val)

    h.printHeap()
    print("The Max val is", h.extractMax())
```