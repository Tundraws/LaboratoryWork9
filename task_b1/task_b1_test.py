import unittest
import requests

class TestFibonacciService(unittest.TestCase):
    URL = "http://localhost:8080/compute"

    def test_fib_logic(self):
        response = requests.post(self.URL, json={"value": 10})
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["result"], 55)

    def test_invalid_method(self):
        response = requests.get(self.URL)
        self.assertEqual(response.status_code, 405)

if __name__ == "__main__":
    unittest.main()