import sys

if __name__ == "__main__":
    print("python test harness running")
    lines = sys.stdin.readlines()

    for line in lines:
        sys.stdout.write("false\n")
