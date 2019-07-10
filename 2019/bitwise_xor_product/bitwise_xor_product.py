def solution(M, N):
    if N == 0:
        return 0
    elif M == 0:
        return bitxor(N)

    bitxor_m = bitxor(M - 1)
    bitxor_n = bitxor(N)

    return bitxor_m ^ bitxor_n


# Reference: http://bit.ly/335GI25
def bitxor(num):
    remainder = num % 4

    if remainder == 0:
        return num
    elif remainder == 1:
        return 1
    elif remainder == 2:
        return num + 1
    else:
        return 0
