import socket
import unittest
import time
import subprocess
import os

class TestTcpServerIntegration(unittest.TestCase):
    HOST = "localhost"
    PORT = 8080
    BUFFER_SIZE = 1024
    EXECUTABLE_PATH = os.path.join(os.path.dirname(__file__), "server.exe")

    @classmethod
    def setUpClass(cls):
        cls.server_process = subprocess.Popen([cls.SERVER_EXECUTABLE_PATH if hasattr(cls, 'SERVER_EXECUTABLE_PATH') else cls.EXECUTABLE_PATH])
        time.sleep(1)

    @classmethod
    def tearDownClass(cls):
        cls.server_process.terminate()

    def send_request(self, message):
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
            client_socket.connect((self.HOST, self.PORT))
            formatted_message = f"{message}\n"
            client_socket.sendall(formatted_message.encode())
            return client_socket.recv(self.BUFFER_SIZE).decode().strip()

    def test_basic_conversion(self):
        self.assertEqual(self.send_request("hello"), "HELLO")

    def test_numbers_and_symbols(self):
        self.assertEqual(self.send_request("123!@#"), "123!@#")

if __name__ == "__main__":
    unittest.main()