import requests
import sys

def get_fibonacci(n):
    url = "http://localhost:8080/compute"
    try:
        response = requests.post(url, json={"value": n}, timeout=10)
        response.raise_for_status()
        return response.json()
    except requests.RequestException as e:
        return {"status": "error", "message": str(e)}

if __name__ == "__main__":
    n = 35
    print(f"Requesting Fibonacci({n})...")
    result = get_fibonacci(n)
    print(f"Result: {result}")