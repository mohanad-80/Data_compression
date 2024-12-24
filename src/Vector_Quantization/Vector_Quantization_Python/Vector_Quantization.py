import cv2
import numpy as np
import json


def euclidean_distance(vec1, vec2):
    """
    Calculate the Euclidean distance between two vectors.
    :param vec1: First vector.
    :param vec2: Second vector.
    :return: Euclidean distance.
    """
    distance = 0.0

    # Loop through each pair of elements in the vectors
    for i in range(len(vec1)):
        difference = vec1[i] - vec2[i]  # Compute the difference
        distance += difference**2  # Add the square of the difference

    # Return the square root of the accumulated sum
    return distance**0.5


def calculate_mean(blocks):
    """
    Calculate the mean vector for a list of blocks.
    :param blocks: List of 1D blocks (flattened lists).
    :return: A 1D list representing the mean vector.
    """
    block_size = len(blocks[0])  # Number of elements in each block
    mean = [0] * block_size  # vector storing the mean with the same block size
    for block in blocks:
        for i in range(block_size):
            mean[i] += block[i]
    return [x / len(blocks) for x in mean]


def lbg_algorithm(blocks, codebook_size, delta=1e-5):
    """
    Perform the Linde-Buzo-Gray (LBG) algorithm to create a codebook.
    :param blocks: List of 1D blocks (flattened lists).
    :param codebook_size: Desired size of the codebook.
    :param delta: Small perturbation value for splitting vectors.
    :return: Final codebook.
    """
    # Initialize with the global centroid
    codebook = [calculate_mean(blocks)]

    while len(codebook) < codebook_size:
        # Split codebook
        new_codebook = []
        for vector in codebook:
            new_codebook.append([v - delta for v in vector])  # Perturb down
            new_codebook.append([v + delta for v in vector])  # Perturb up

        codebook = new_codebook

        # Assign blocks to the nearest codebook vector and recompute centroids
        while True:
            assignments = [[] for _ in range(len(codebook))]
            for block in blocks:
                distances = [
                    euclidean_distance(block, codeword) for codeword in codebook
                ]
                closest_idx = distances.index(min(distances))
                assignments[closest_idx].append(block)

            new_codebook = []
            for group in assignments:
                if group:  # Avoid empty clusters
                    new_codebook.append(calculate_mean(group))
                else:
                    new_codebook.append([0] * len(blocks[0]))  # Handle empty clusters

            # Check for convergence
            converged = True
            for i in range(len(codebook)):
                for a, b in zip(codebook[i], new_codebook[i]):
                    if abs(a - b) >= delta:
                        converged = False
                        break
                if not converged:
                    break

            if converged:
                break
            codebook = new_codebook

    return codebook


def extract_block(image, i, j, block_size):
    """
    Extract a block of size block_size x block_size from the image, with zero padding for out-of-bound indices.
    :param image: 2D list representing the grayscale pixel values.
    :param i: Starting row of the block.
    :param j: Starting column of the block.
    :param block_size: The size of the block (e.g., 2 for 2x2 blocks).
    :return: A 1D list representing the flattened block.
    """
    height = len(image)
    width = len(image[0])
    block = []

    # Iterate through the rows of the block
    for x in range(i, i + block_size):
        # Iterate through the columns of the block
        for y in range(j, j + block_size):
            if x < height and y < width:
                # If within bounds, add the pixel value
                block.append(image[x][y])
            else:
                # If out of bounds, add zero for padding
                block.append(0)

    return block


def divide_into_blocks(image, block_size):
    """
    Divide a 2D image array into non-overlapping blocks of size block_size x block_size.
    :param image: 2D list representing the grayscale pixel values.
    :param block_size: Size of the block (e.g., 4 for 4x4 blocks).
    :return: A list of 1D blocks (each block is a flattened list of pixel values).
    """
    height = len(image)
    width = len(image[0])
    blocks = []

    # Iterate through the image in steps of block_size
    for i in range(0, height, block_size):
        for j in range(0, width, block_size):
            # Use the extract_block function to get the block
            block = extract_block(image, i, j, block_size)
            blocks.append(block)

    return blocks


def read_image(image_path):
    """
    Reads a grayscale image and converts it to a 2D list of pixel values.
    :param image_path: Path to the input image file.
    :return: A 2D list representing the grayscale pixel values.
    """
    # Read the image in grayscale mode
    image = cv2.imread(image_path, cv2.IMREAD_GRAYSCALE)

    if image is None:
        raise FileNotFoundError(f"Image file not found: {image_path}")

    # Convert the OpenCV image (NumPy array) into a list of lists
    pixel_values = [[int(pixel) for pixel in row] for row in image]

    return pixel_values


def write_image(image, output_path):
    """
    Save a 2D list as a grayscale image using OpenCV.
    :param image: 2D list representing the grayscale pixel values.
    :param output_path: Path to save the grayscale image.
    """
    # Convert the 2D list to a NumPy array (required by OpenCV)
    image_array = np.array(image, dtype=np.uint8)

    # Save the image using OpenCV
    success = cv2.imwrite(output_path, image_array)

    if not success:
        raise IOError(f"Failed to write image to {output_path}")


def compress_image(image_path, block_size, codebook_size, compressed_file):
    """
    Compress an image using Vector Quantization.
    :param image_path: Path to the input grayscale image.
    :param block_size: Size of the blocks (e.g., 2 for 2x2 blocks).
    :param codebook_size: Desired size of the codebook.
    :param compressed_file: Path to save the compressed data.
    """
    # Read the image and divide it into blocks
    image = read_image(image_path)
    blocks = divide_into_blocks(image, block_size)

    # Generate the codebook using LBG algorithm
    codebook = lbg_algorithm(blocks, codebook_size)

    # Assign each block to the closest codebook vector
    compressed_matrix = []
    for block in blocks:
        distances = [euclidean_distance(block, codeword) for codeword in codebook]
        closest_idx = distances.index(min(distances))
        compressed_matrix.append(closest_idx)

    # Prepare compressed data
    compressed_data = {
        "compressed_matrix": compressed_matrix,  # Indices of codebook vectors
        "codebook": codebook,  # List of codebook vectors
        "image_shape": [len(image), len(image[0])],  # Original image dimensions
        "block_size": block_size,  # Block size used for compression
    }

    # Serialize the data to a JSON string and write it as bytes
    with open(compressed_file, "wb") as f:
        f.write(json.dumps(compressed_data).encode("utf-8"))

    print(f"Image compressed and saved to {compressed_file}")


def decompress_image(compressed_file, decompressed_image_path):
    """
    Decompress a compressed image file to reconstruct the original image.
    :param compressed_file: Path to the compressed data file.
    :param decompressed_image_path: Path to save the decompressed image.
    """
    # Read the binary data and deserialize the JSON
    with open(compressed_file, "rb") as f:
        compressed_data = json.loads(f.read().decode("utf-8"))

    compressed_matrix = compressed_data[
        "compressed_matrix"
    ]  # Indices of codebook vectors
    codebook = compressed_data["codebook"]  # List of codebook vectors
    image_shape = compressed_data["image_shape"]  # Original image dimensions
    block_size = compressed_data["block_size"]  # Block size used for compression

    height, width = image_shape
    decompressed_image = [[0] * width for _ in range(height)]  # Initialize empty image

    # Reconstruct the image block by block
    idx = 0
    for i in range(0, height, block_size):
        for j in range(0, width, block_size):
            if idx < len(compressed_matrix):
                vector = codebook[compressed_matrix[idx]]  # Get the codebook vector
                for bi in range(block_size):
                    for bj in range(block_size):
                        if i + bi < height and j + bj < width:
                            # Map vector back to the 2D image
                            decompressed_image[i + bi][j + bj] = int(
                                vector[bi * block_size + bj]
                            )
                idx += 1

    # Save the reconstructed image
    write_image(decompressed_image, decompressed_image_path)
    print(f"Image decompressed and saved to {decompressed_image_path}")


# Example usage
if __name__ == "__main__":
    block_size = 2
    codebook_size = 16

    images = ["fruit.bmp", "House.bmp", "photographer.bmp", "image.jpg"]
    image_path = images[2]
    compressed_file = (
        image_path[:-4] + "-b" + str(block_size) + "-c" + str(codebook_size) + ".bin"
    )
    decompressed_image_path = (
        image_path[:-4] + "-b" + str(block_size) + "-c" + str(codebook_size) + ".bmp"
    )

    print("Starting compression...")
    compress_image(image_path, block_size, codebook_size, compressed_file)
    decompress_image(compressed_file, decompressed_image_path)
