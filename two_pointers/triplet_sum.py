from typing import List


def triplet_sum(nums: List[int]) -> List[List[int]]:
    # list must be sorted first
    nums.sort()

    solution = [[]]
    for i in range(len(nums)):
        if i > 0 and i == i - 1:
            continue  # skip dupes
        sums = pair_sum_modified(nums, i + 1, -nums[i])
        for sum in sums:
            solution.append(sum)

    return solution


def pair_sum_modified(nums: List[int], start: int, target: int) -> List[List[int]]:
    sums = []
    left, right = start, len(nums) - 1
    while left < right:
        sum = nums[left] + nums[right]
        if sum == target:
            sums.append([left, right])
            continue
        if sum < target:
            left += 1
        if sum > target:
            right -= 1

    return sums
