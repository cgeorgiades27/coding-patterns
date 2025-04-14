# Pair Sum Unsorted

## Description

This package provides a function to find a pair of numbers in an unsorted array that sum up to a given target value. The function returns the indices of the two numbers that satisfy the condition.

## Algorithm

The `pairsumunsorted` function in `pairsumunsorted.go` finds the pair of numbers using the following steps:

1.  Create a hash map `complimentMap` to store each number from the input array `nums` and its index.
2.  Iterate through the `nums` array.
3.  For each number, calculate its complement by subtracting it from the `target`.
4.  Check if the complement exists in the `complimentMap`.
    *   If the complement exists, return an array containing the index of the complement (from the map) and the current index.
5.  If the complement is not found, store the current number and its index in the `complimentMap`.
6.  If no such pair is found after iterating through the entire array, return an empty array.

## Files

*   `pairsumunsorted.go`: Contains the implementation of the `pairsumunsorted` function.
*   `pairsumunsorted_test.go`: Contains unit tests for the `pairsumunsorted` function.