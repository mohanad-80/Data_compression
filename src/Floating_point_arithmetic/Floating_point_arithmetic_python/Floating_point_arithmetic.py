import os
import json
from collections import Counter
import struct

class Range:
    def __init__(self, low, high):
        self.low = low
        self.high = high

def build_freq_table(data):
    return Counter(data)

def build_ranges_table(freq_table, data_size):
    ranges_table = {}
    cumulative_freq = 0

    for symbol, frequency in freq_table.items():
        low = cumulative_freq / data_size
        high = (cumulative_freq + frequency) / data_size
        ranges_table[symbol] = Range(low, high)
        cumulative_freq += frequency

    return ranges_table

def compress(input_file, output_file):
    # Read input data
    with open(input_file, 'r', encoding='utf-8') as f:
        input_text = f.read()

    # Build frequency table and ranges table
    freq_table = build_freq_table(input_text)
    ranges_table = build_ranges_table(freq_table, len(input_text))

    # Perform arithmetic compression
    lower = 0.0
    upper = 1.0
    for char in input_text:
        symbol_range = ranges_table[char]
        range_width = upper - lower
        upper = lower + range_width * symbol_range.high
        lower = lower + range_width * symbol_range.low

    final_value = (upper + lower) / 2
    print(final_value)

    # Write compressed data and metadata to a binary file
    with open(output_file, 'wb') as f:
        # Write frequency table as JSON
        freq_table_bytes = json.dumps(freq_table).encode('utf-8')
        f.write(len(freq_table_bytes).to_bytes(4, 'big'))  # 4 bytes for length
        f.write(freq_table_bytes)

        # Write compressed value in binary
        compressed_value_bytes = struct.pack('d', final_value)  # 'd' is for double precision (float64)
        f.write(compressed_value_bytes)

def decompress(input_file, output_file):
    # Read compressed data and metadata
    with open(input_file, 'rb') as f:
        freq_table_length = int.from_bytes(f.read(4), 'big')
        freq_table_bytes = f.read(freq_table_length)
        freq_table = json.loads(freq_table_bytes.decode('utf-8'))

        # Read the compressed value in binary
        compressed_value_bytes = f.read(8)  # 8 bytes for double precision
        compressed_value = struct.unpack('d', compressed_value_bytes)[0]

    # Rebuild ranges table
    total_symbols = sum(freq_table.values())
    ranges_table = build_ranges_table(freq_table, total_symbols)

    # Perform arithmetic decompression
    decoded_text = []
    value = compressed_value
    for _ in range(total_symbols):
        for symbol, symbol_range in ranges_table.items():
            if symbol_range.low <= value < symbol_range.high:
                decoded_text.append(symbol)
                range_width = symbol_range.high - symbol_range.low
                value = (value - symbol_range.low) / range_width
                break

    # Write decompressed data to output file
    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(''.join(decoded_text))

def main_menu():
    while True:
        print("\nFloating Point Arithmetic Compression")
        print("1. Compress a file")
        print("2. Decompress a file")
        print("3. Exit")
        
        choice = input("Enter your choice: ")
        
        if choice == '1':
            input_file = input("Enter the name of the file to compress: ")
            output_file = input("Enter the name of the output compressed file: ")
            if os.path.exists(input_file):
                compress(input_file, output_file)
                print(f"File '{input_file}' compressed successfully to '{output_file}'.")
                print('\\o/')
            else:
                print("Input file does not exist.")
        elif choice == '2':
            input_file = input("Enter the name of the file to decompress: ")
            output_file = input("Enter the name of the output decompressed file: ")
            if os.path.exists(input_file):
                decompress(input_file, output_file)
                print(f"File '{input_file}' decompressed successfully to '{output_file}'.")
                print('\\o/')
            else:
                print("Input file does not exist.")
        elif choice == '3':
            print("Exiting the program. Goodbye!")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main_menu()

# Example usage
# compress("input.txt", "compressed.bin")
# decompress("compressed.bin", "decompressed.txt")

# Directory-based operations
# compressed_dir = "compressed_texts"
# decompressed_dir = "decompressed_texts"
# os.makedirs(compressed_dir, exist_ok=True)
# os.makedirs(decompressed_dir, exist_ok=True)

# text_dir = "texts"
# for filename in os.listdir(text_dir):
#     if filename.endswith(".txt"):
#         input_file = os.path.join(text_dir, filename)
#         compressed_file = os.path.join(compressed_dir, filename[:-4] + ".bin")
#         decompressed_file = os.path.join(decompressed_dir, filename)

#         print(f"Compressing {filename}...")
#         compress(input_file, compressed_file)
#         print(f"Compression of {filename} complete.")
#         print(f"Decompressing {filename}...")
#         decompress(compressed_file, decompressed_file)
#         print(f"Decompression of {filename} complete.")
