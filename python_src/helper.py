# If you need to import stuff

def text_to_hex(text):
    return ''.join([hex(ord(c))[2:] for c in text])

def hex_to_text(hex):
    return ''.join([chr(int(hex[i:i+2], 16)) for i in range(0, len(hex), 2)])

def split_into_blocks(text, block_size=128):
    text_hex = text_to_hex(text)
    # 1 hexadecimal digit = 4 bits
    # So, 128 bits = 32 hexadecimal digits
    print(len(text_hex))
    blocks = [text_hex[i:i+block_size//4] for i in range(0, len(text_hex), block_size//4)]
    for block_idx in range(len(blocks)):
        blocks[block_idx] = blocks[block_idx].ljust(block_size//4, '0')
    print(blocks)
    return blocks

text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

split_into_blocks(text)
