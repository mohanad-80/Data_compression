import os
import csv

def calculate_compression_ratio(original_size, compressed_size):
    return original_size / compressed_size

def analyze_compression(text_dir, compressed_dir):
    compressed_dir = "compressed_texts"
    text_dir = "texts"
    stats = []

    for filename in os.listdir(text_dir):
        if filename.endswith(".txt"):
            original_file = os.path.join(text_dir, filename)
            compressed_file = os.path.join(compressed_dir, filename[:-4] + ".bin")

            original_size = os.path.getsize(original_file)
            compressed_size = os.path.getsize(compressed_file)

            compression_ratio = calculate_compression_ratio(original_size, compressed_size)
            stats.append({"File Name": filename, "Original Size": original_size, "Compressed Size": compressed_size, "Compression Ratio": compression_ratio})

    print("Average Compression Ratio:", sum(d["Compression Ratio"] for d in stats) / len(stats))

    with open("compression_stats.csv", "w", newline="") as csvfile:
        writer = csv.DictWriter(csvfile, fieldnames=["File Name", "Original Size", "Compressed Size", "Compression Ratio"])
        writer.writeheader()
        writer.writerows(stats)

if __name__ == "__main__":
    text_dir = "texts"
    compressed_dir = "compressed_texts"
    analyze_compression(text_dir, compressed_dir)
