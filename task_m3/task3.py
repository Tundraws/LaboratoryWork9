import subprocess
import json
import unittest
import os

def run_go_calculator(numbers):
    executable = os.path.join("go_task", "calculator.exe")
    
    proc = subprocess.Popen(
        [executable], 
        stdin=subprocess.PIPE, 
        stdout=subprocess.PIPE, 
        stderr=subprocess.PIPE
    )
    
    input_data = json.dumps({"numbers": numbers})
    stdout, stderr = proc.communicate(input=input_data.encode())
    
    if proc.returncode != 0:
        raise RuntimeError(f"Go error: {stderr.decode()}")
        
    return json.loads(stdout.decode())["sum"]

class TestGoIntegration(unittest.TestCase):

    def test_sum_squares_success(self):
        self.assertEqual(run_go_calculator([1, 2, 3]), 14)

    def test_empty_list(self):
        self.assertEqual(run_go_calculator([]), 0)

    def test_large_numbers(self):
        self.assertEqual(run_go_calculator([1000, 2000]), 5000000)

    def test_invalid_data_type(self):
        with self.assertRaises(Exception):
            run_go_calculator(["apple", "orange"])

if __name__ == "__main__":
    unittest.main()