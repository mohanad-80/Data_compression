import os
import heapq
import json
from collections import Counter
from bitarray import bitarray

class Node:
    def __init__(self, char=None, freq=0, left=None, right=None):
        self.char = char
        self.freq = freq
        self.left = left
        self.right = right

    def __lt__(self, other):
        return self.freq < other.freq

def build_huffman_tree(freq_table):
    heap = [Node(char, freq) for char, freq in freq_table.items()]
    heapq.heapify(heap)
    
    while len(heap) > 1:
        left = heapq.heappop(heap)
        right = heapq.heappop(heap)
        merged = Node(freq=left.freq + right.freq, left=left, right=right)
        heapq.heappush(heap, merged)
    
    return heap[0]

def generate_codes(node, prefix="", codes=None):
    if codes is None:
        codes = {}
    if node.char is not None:  # Leaf node
        codes[node.char] = bitarray(prefix)  # Convert prefix to a bitarray
    else:
        generate_codes(node.left, prefix + "0", codes)
        generate_codes(node.right, prefix + "1", codes)
    return codes

def compress(input_file, output_file):
    with open(input_file, 'r', encoding='utf-8') as f:
        input_text = f.read()

    freq_table = Counter(input_text)
    tree = build_huffman_tree(freq_table)
    codes = generate_codes(tree)
    
    # Encode the input text as a bitarray
    compressed = bitarray()
    compressed.encode(codes, input_text)
    
    # Write the frequency table and compressed bitstream
    with open(output_file, 'wb') as f:
        # Serialize frequency table as JSON
        freq_table_bytes = json.dumps(freq_table).encode('utf-8')
        
        # Write the length of the frequency table (4 bytes, fixed size)
        f.write(len(freq_table_bytes).to_bytes(4, 'big'))
        
        # Write frequency table
        f.write(freq_table_bytes)
        
        # Write compressed bitstream
        compressed.tofile(f)

def decompress(input_file, output_file):
    with open(input_file, 'rb') as f:
        # Read the length of the frequency table
        freq_table_length = int.from_bytes(f.read(4), 'big')
        
        # Read the frequency table
        freq_table_bytes = f.read(freq_table_length)
        freq_table = json.loads(freq_table_bytes.decode('utf-8'))
        
        # Read the compressed bitstream
        compressed_bitstream = bitarray()
        compressed_bitstream.fromfile(f)
    
    tree = build_huffman_tree(freq_table)
    
    # Decode the bitstream
    decoded_text = []
    node = tree
    for bit in compressed_bitstream:
        node = node.left if bit == 0 else node.right
        if node.char is not None:  # Leaf node
            decoded_text.append(node.char)
            node = tree
        if len(decoded_text) == sum(freq_table.values()):
            break  # Stop when all characters are decoded
    
    with open(output_file, 'w') as f:
        f.write(''.join(decoded_text))

# Example Usage
# if __name__ == "__main__":
#     # input_text = "hello huffman"
#     compress("input.txt", "compressed.bin")
#     decompress("compressed.bin", "decompressed.txt")

##############################################
##############################################
##############################################
##############################################
##############################################

# Create directories if they don't exist
compressed_dir = "compressed_texts"
decompressed_dir = "decompressed_texts"
os.makedirs(compressed_dir, exist_ok=True)
os.makedirs(decompressed_dir, exist_ok=True)

# Loop through files in the "texts" directory
text_dir = "texts"
for filename in os.listdir(text_dir):
    if filename.endswith(".txt"):
        input_file = os.path.join(text_dir, filename)
        compressed_file = os.path.join(compressed_dir, filename[:-4] + ".bin")
        decompressed_file = os.path.join(decompressed_dir, filename)

        print(f"Compressing {filename}...")
        compress(input_file, compressed_file)
        print(f"Compression of {filename} complete.")
        print(f"Decompressing {filename}...")
        decompress(compressed_file, decompressed_file)
        print(f"Decompression of {filename} complete.")