import csv
import os
import re

log_files = {}
for root, dirs, files in os.walk("benchmarks"):
    for file in files:
        if file.endswith(".log"):
            log_files[os.path.join(root, file)] = os.path.join(root, file.replace(".log", ".csv"))


benchmark_pattern = re.compile(
    r"Benchmark([\w_]+)/numRevisions=(\d+),dataSize=(\d+)-(\d+)\s+(\d+)\s+(\d+) ns/op\s+(\d+) B/op\s+(\d+) allocs/op"
)

csv_headers = [
    "benchmark_name",
    "num_revisions",
    "data_size",
    "num_cpus",
    "iterations",
    "ns_per_op",
    "bytes_per_op",
    "allocs_per_op",
]

for log_file, csv_file in log_files.items():
    with open(log_file, "r") as f:
        content = f.read()

    benchmarks = benchmark_pattern.findall(content)

    with open(csv_file, "w", newline="") as f:
        writer = csv.writer(f)
        writer.writerow(csv_headers)
        for benchmark in benchmarks:
            writer.writerow(benchmark)
