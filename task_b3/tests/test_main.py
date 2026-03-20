import pytest
import melnikova_a

BASE_VALUE = 10
INCREMENT_VALUE = 20
EXPECTED_SUM = 30

TEXT_TO_REVERSE = "engineering"
EXPECTED_REVERSED = "gnireenigne"

EMPTY_STRING = ""

class TestRustModuleIntegration:

    def test_addition_logic(self):
        result = melnikova_a.calculate_sum(BASE_VALUE, INCREMENT_VALUE)
        assert result == EXPECTED_SUM, f"Ожидалось {EXPECTED_SUM}, получено {result}"

    def test_string_reversal_logic(self):
        result = melnikova_a.get_reversed_string(TEXT_TO_REVERSE)
        assert result == EXPECTED_REVERSED, "Строка развернута некорректно"

    def test_empty_string_handling(self):
        assert melnikova_a.get_reversed_string(EMPTY_STRING) == EMPTY_STRING

    @pytest.mark.parametrize("a, b, expected", [
        (0, 0, 0),
        (-1, -1, -2),
        (100, -50, 50)
    ])
    def test_parametrized_addition(self, a, b, expected):
        assert melnikova_a.calculate_sum(a, b) == expected

    def test_invalid_types(self):
        with pytest.raises(TypeError):
            melnikova_a.calculate_sum("string_instead_of_int", BASE_VALUE)