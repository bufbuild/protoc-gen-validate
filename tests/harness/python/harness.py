import sys

from harness_pb2 import TestCase, TestResult

if __name__ == "__main__":
    lines = sys.stdin.readlines()

    result = TestResult()
    result.AllowFailure = True
    sys.stdout.write(result.SerializeToString())
