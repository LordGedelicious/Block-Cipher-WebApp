# Create Block Cipher Main Implementation Based on AES
# Steps:
# 1. Receive plaintext, key, and IV from user
# 1a. Convert plaintext and key to hex (1 hexadecimal digit = 4 bits)
# 2. Split plaintext into 128-bit blocks
# 3. Only take the first 128 bits of the key. If the key is less than 128 bits, pad it with 0s
# 4. Create Feistel Network
# 4a. Split the plaintext into 8 16-bit blocks (because the S-box is 16x16 bits)
# 4b. Generate round keys from the key (use the key schedule)
# 4c. For each round, perform the following operations:
# 4c1. XOR the round key with the right half of the block
# 4c2. Pass the right half of the block through the S-box
# 4c3. Pass the right half of the block through the P-box
# 4c4. XOR the left half of the block with the right half of the block
# 4c5. Swap the left and right halves of the block
# 5. Combine the blocks into new ciphertext
# 6. Repeat steps 4-5 for each block for 16 times
# 7. Return the ciphertext
from helper import *
from cipher import *
from feistel import FeistelNetwork

def main():
    # mode = input("Choose mode (encrypt/decrypt): ")
    # operation_mode = input("Choose operation mode (ECB/CBC/CFB/OFB/Counter Mode): ")
    key = input("Enter key: ")
    key = spliceKey(key)
    subkeys = SubKeyGenerator(key, 16)
    print(f"Subkeys:")
    for i in range(16):
        print(f"Round {i+1}: {SubKeyGenerator(key, i)}")
    # plaintext = input("Enter plaintext: ")
    # if mode == "encrypt":
    #     block_cipher = FeistelNetwork(plaintext, key)
    #     ciphertext = block_cipher.encrypt()
    #     print(f"Ciphertext: {ciphertext}")
    # elif mode == "decrypt":
    #     ciphertext = input("Enter ciphertext: ")
    #     block_cipher = FeistelNetwork(ciphertext, key)
    #     plaintext = block_cipher.decrypt()
    #     print(f"Plaintext: {plaintext}")
    # else:
    #     print("Invalid mode")
    #     return

if __name__ == "__main__":
    main()