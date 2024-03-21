# Implementation of Feistel Network
from cipher import *
from helper import *

class FeistelNetwork:
    def __init__(self, lhs, rhs, key, loop=0):
        self.lhs = lhs # Left half of the block, string of 16 hexadecimal digits
        self.rhs = rhs # Right half of the block, string of 16 hexadecimal digits
        self.key = key # Original key inputted, key generation use loop count
        self.loop = loop # Loop count for key generation and for Feistel Network, 0 if not started yet
    
    def __str__(self):
        return f"FeistelNetwork(lhs={self.lhs}, rhs={self.rhs}, key={self.key}, loop={self.loop})"

    def encrypt(self):
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
        rhs_permuted = ShiftRows(rhs_substituted)
        xor_result = xor_operation(rhs_permuted, subkey)
        return xor_result