BASE62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


def base62_encode(data: str) -> str:
    num = int.from_bytes(data.encode(), "big")
    if num == 0:
        return BASE62[0]
    res = []
    while num > 0:
        res.append(BASE62[num % 62])
        num //= 62
    return "".join(reversed(res))


def base62_decode(encoded: str) -> str:
    num = 0
    for c in encoded:
        num = num * 62 + BASE62.index(c)
    byte_length = (num.bit_length() + 7) // 8
    return num.to_bytes(byte_length, "big").decode()
