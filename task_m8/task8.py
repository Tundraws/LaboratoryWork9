import task_m8
import unittest

class TestRustModule(unittest.TestCase):
    def test_basic_logic(self):
        self.assertEqual(task_m8.calculate_sum_of_squares([1, 2, 3]), 14)

    def test_empty_list(self):
        self.assertEqual(task_m8.calculate_sum_of_squares([]), 0)

    def test_negative_numbers(self):
        self.assertEqual(task_m8.calculate_sum_of_squares([-1, -2]), 5)

if __name__ == "__main__":
    unittest.main()