with open("enc", "r") as file:
    enc = file.read()

# flag = enc

# fchar = chr((ord(flag[0]) >> 8))

# ''.join([chr((ord(flag[i]) >> 8) + ord(flag[i+1])) for i in range(0, len(flag), 2)])

# flag = ''
# for i in range(0, len(enc)):
    # flag += chr((ord(enc[i]) >> 8)) # - ord(flag[i+1]))

# print(flag)

# print(chr(ord(enc[0]) >> 8))
# print(ord(enc[1]) >> 8)
# print(ord(enc[1]))

flag = "pico"

# print(''.join([chr((ord(flag[i]) << 8) + ord(flag[i + 1])) for i in range(0, len(flag), 2)]))

# print(ord(enc[0]) >> 8)

flag = ''
for ch in enc:
    bin_str = format(ord(ch), 'b')
    while (len(bin_str) < 16):
        bin_str = '0' + bin_str
    first_half = bin_str[:8]
    second_half = bin_str[8:]
    flag += chr(int(first_half,2))
    flag += chr(int(second_half, 2))

print(flag)
