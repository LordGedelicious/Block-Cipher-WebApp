# If you need to import stuff

def text_to_hex(text):
    return text.encode('UTF-8').hex()

def hex_to_text(hex):
    return bytes.fromhex(hex).decode('UTF-8')

def hex_to_bytes(hex):
    return bytes([int(hex[i:i+2], 16) for i in range(0, len(hex), 2)])

def bytes_to_hex(bytes):
    return ''.join([hex(b)[2:] for b in bytes])

def split_into_blocks(text, block_size=128):
    text_hex = text_to_hex(text)
    blocks = [text_hex[i:i+block_size//4] for i in range(0, len(text_hex), block_size//4)]
    for block_idx in range(len(blocks)):
        blocks[block_idx] = blocks[block_idx].ljust(block_size//4, '0')
    return blocks

def string_to_arr(string):
    # Split a hexadecimal string to an array of 2-digit hexadecimal strings
    return [hex(string[i:i+2]) for i in range(0, len(string), 2)]

def arr_to_string(arr):
    # Join an array of 2-digit hexadecimal strings to a single hexadecimal string
    return ''.join(arr)

def spliceKey(key):
    # If key is less than 128 bit, pad it with 0s
    # If key is more than 128 bit, take the first 128 bit
    key_hex = text_to_hex(key)
    if len(key_hex) < 32:
        key_hex = key_hex.ljust(32, '0')
    elif len(key_hex) > 32:
        key_hex = key_hex[:32]
    return key_hex

def hexstring_to_intarray(hexstring):
    # Convert a hexadecimal string to an array of hexadecimal values
    return [(int(hexstring[i:i+2], 16)) for i in range(0, len(hexstring), 2)]

def intarray_to_hexstring(intarray):
    # Convert an array of hexadecimal values to a hexadecimal string
    return ''.join(intarray)

hexstring = "48656c6c6f2c20576f726c6421"
print(hexstring_to_intarray(hexstring))
