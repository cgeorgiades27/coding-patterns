# Largest Container

## Description

This directory contains a Go implementation for finding the largest container (area) that can be formed by vertical lines of varying heights. The algorithm uses a two-pointer approach to efficiently determine the maximum area.

## Algorithm

The `largestcontainer` function in `largestcontainer.go` calculates the largest container area using the following steps:

1.  Initialize two pointers, `left` and `right`, to the start and end of the input array `nums`, respectively.
2.  Initialize `maxArea` to 0.
3.  Iterate while `left < right`:
    *   Calculate the area formed by the lines at `left` and `right`. The area is the minimum of the heights at `left` and `right` multiplied by the distance between them (`right - left`).
    *   Update `maxArea` if the calculated area is greater than the current `maxArea`.
    *   Move the pointer with the smaller height inward. If the heights are equal, move both pointers.
4.  Return `maxArea`.

## Files

*   `largestcontainer.go`: Contains the implementation of the `largestcontainer` function.
*   `largestcontainer_test.go`: Contains unit tests for the `largestcontainer` function.

## Testing

To run the unit tests, use the following command:

```bash
go test ./twopointers/largestcontainer
```

## Complexity

*   **Time Complexity:** O(n), where n is the number of elements in the input array.
*   **Space Complexity:** O(1), as the algorithm uses a constant amount of extra space.
```