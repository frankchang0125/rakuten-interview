from bitwise_xor_product import solution


def test_solution():
    assert solution(5, 0) == 0
    assert solution(0, 6) == 7
    assert solution(5, 8) == 12
    assert solution(10, 4) == 5
    assert solution(987654321, 101001) == 987654321
