# Implementation of Feistel Network
from cipher import *
from helper import *

class FeistelNetwork:
    def __init__(self, text, key, loop=0):
        self.arr_text = split_into_blocks(text)
        self.lhs = "" # Left half of the block, string of 16 hexadecimal digits
        self.rhs = "" # Right half of the block, string of 16 hexadecimal digits
        self.key = key # Original key inputted, key generation use loop count
        self.loop = loop # Loop count for key generation and for Feistel Network, 0 if not started yet
    
    def __str__(self):
        return f"FeistelNetwork(lhs={self.lhs}, rhs={self.rhs}, key={self.key}, loop={self.loop})"
    
    def combine_arr_text(self):
        # Combine the blocks into a single string
        return ''.join(self.arr_text)

    def encrypt(self):
        # # Loop for 16 times
        # for i in range(16):
        #     # Generate subkey (round key) from the round loop count
        #     subkey = SubKeyGenerator(self.key, i)
        #     # Perform Feistel Network
        #     new_lhs = self.rhs
        #     new_rhs = xor_operation(self.lhs, self.FeistelFunction(self.rhs, subkey))
        #     # Update the left and right half of the block
        #     self.lhs = new_lhs
        #     self.rhs = new_rhs
        # Loop for all blocks in arr_text
        for block_idx in range(len(self.arr_text)):
            # Split the block into left and right half
            block = self.arr_text[block_idx]
            self.lhs = hexstring_to_intarray(block[:len(block)//2])
            self.rhs = hexstring_to_intarray(block[len(block)//2:])
            # Loop for 16 times
            for i in range(16):
                # Generate subkey (round key) from the round loop count
                subkey = SubKeyGenerator(self.key, i)
                # Perform Feistel Network
                new_lhs = self.rhs
                new_rhs = xor_operation(self.lhs, self.FeistelFunction(self.rhs, subkey))
                # Update the left and right half of the block
                self.lhs = new_lhs
                self.rhs = new_rhs
            # Combine the left and right half of the block
            self.arr_text[block_idx] = self.lhs + self.rhs
        return self
    
    def decrypt(self):
        # TODO: NotImplemented
        pass

    def FeistelFunction(self, rhs, subkey):
        # Feistel function is a function that takes a right half of the block and a subkey as input
        # The function is defined as follows:
        # 1. Perform substitution using S-box
        # 2. Perform permutation using shift-rows
        # 3. Perform XOR with the subkey
        # 4. Return the result
        rhs_substituted = SubBytesSubstitutionArr(rhs)
        print(rhs_substituted)
        rhs_permuted = ShiftRows(rhs_substituted)
        xor_result = xor_operation(rhs_permuted, subkey)
        return xor_result