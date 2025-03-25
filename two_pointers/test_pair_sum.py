import unittest
from .pair_sum import pair_sum


class TestPairSum(unittest.TestCase):

    def test_pair_sum(self):
        self.assertEqual(pair_sum([-5, -2, 3, 4, 6], 7), [2, 3])
        self.assertEqual(pair_sum([1, 1, 1], 2), [0, 2])
        self.assertEqual(pair_sum([], 0), [])
        self.assertEqual(pair_sum([1], 1), [])
        self.assertEqual(pair_sum([2, 3], 5), [0, 1])
        self.assertEqual(pair_sum([2, 2, 3], 5), [0, 2])
        self.assertEqual(pair_sum([-1, 2, 3], 2), [0, 2])
        self.assertEqual(pair_sum([-3, -2, -1], -5), [0, 1])


if __name__ == '__main__':
    unittest.main()
