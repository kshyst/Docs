# Majority

Given an array arr. Find the majority element in the array. If no majority exists, return -1. A majority element in an array is an element that appears strictly more than arr.size() / 2 times in the array.

```text
Input : arr[] = {1, 1, 2, 1, 3, 5, 1}
Output : 1
Explanation: Note that 1 appear 4 times which is more than  7 / 2 times 


Input : arr[] = {3, 3, 4, 2, 4, 4, 2, 4}
Output :  -1 
Explanation: There is no element whose frequency is greater than the half of the size of the array size.


Input : arr[] = {3}
Output : 3
Explanation: Appears more than n/2 times
```

## Finding Majority

### Sorting O(nlogn)

The intuition behind this approach is that if an element occurs more than n/2 times in the array (where n is the size of the array), it will always occupy the middle position when the array is sorted. Therefore, we can sort the array and return the element at index n/2.

### Hashmap O(n)

The intuition behind using a hash map is to count the occurrences of each element in the array and then identify the element that occurs more than n/2 times. By storing the counts in a hash map, we can efficiently keep track of the occurrences of each element.

### Moore voting algorithm

Initialize two variables: count and candidate. Set count to 0 and candidate to an arbitrary value.
Iterate through the array nums:

- a. If count is 0, assign the current element as the new candidate and increment count by 1.
- b. If the current element is the same as the candidate, increment count by 1.
- c. If the current element is different from the candidate, decrement count by 1.

After the iteration, the candidate variable will hold the majority element.