import unittest
import melnikova_a

class TestMelnikovaModule(unittest.TestCase):
    def test_sum_positive(self):
        self.assertEqual(melnikova_a.sum_as_string(10, 20), "30")

    def test_sum_zero(self):
        self.assertEqual(melnikova_a.sum_as_string(0, 0), "0")

    def test_large_numbers(self):
        self.assertEqual(melnikova_a.sum_as_string(1000000, 1000000), "2000000")

    def test_type_error(self):
        with self.assertRaises(TypeError):
            melnikova_a.sum_as_string("10", 20)

    def test_module_metadata(self):
        self.assertEqual(melnikova_a.__name__, "melnikova_a")

if __name__ == "__main__":
    unittest.main()