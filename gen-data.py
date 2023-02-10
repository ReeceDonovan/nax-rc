import os
import random
import sys
import time


def populate_revlog(file_name, revision_range):
    revisions = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J"]
    num_revisions = random.randint(revision_range[0], revision_range[1])
    current_content = ""
    with open(file_name, "w") as f:
        timestamp = time.time_ns() // 1000
        for i in range(num_revisions):
            revision = random.choice(revisions)
            current_content += revision
            timestamp += random.randint(1, 1000)
            f.write(f"Revision {i + 1}: {revision} ({timestamp})\n")
            f.write(f"Content: {current_content}\n")


def generate_revlogs(num_files, revision_range):
    if os.path.exists("test_data"):
        for filename in os.listdir("test_data"):
            os.remove(os.path.join("test_data", filename))
    os.makedirs("test_data", exist_ok=True)
    for i in range(num_files):
        file_name = f"test_data/file{i + 1}.txt"
        populate_revlog(file_name, revision_range)


if len(sys.argv) != 3:
    print("Usage: python script.py NUM_FILES [MIN_REVISIONS, MAX_REVISIONS]")
    sys.exit(1)

num_files = int(sys.argv[1])
revision_range = list(map(int, sys.argv[2].strip("[]").split(",")))

generate_revlogs(num_files, revision_range)
