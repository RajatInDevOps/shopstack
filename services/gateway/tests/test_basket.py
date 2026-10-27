import pytest

from basket import Basket


@pytest.fixture
def basket():
    return Basket()


def test_a_new_basket_is_empty(basket):
    assert basket.count() == 0
    assert basket.total() == 0
    assert basket.items() == []


def test_adding_the_same_sku_twice_merges_the_lines(basket):
    basket.add("SKU-1", 1, 10.0)
    basket.add("SKU-1", 2, 10.0)

    assert len(basket.items()) == 1
    assert basket.count() == 3


def test_the_total_multiplies_quantity_by_price(basket):
    basket.add("SKU-1", 2, 9.99)
    basket.add("SKU-2", 1, 0.02)

    assert basket.total() == 20.0


def test_removing_a_line_takes_it_out(basket):
    basket.add("SKU-1", 1, 5.0)

    assert basket.remove("SKU-1") is True
    assert basket.remove("SKU-1") is False
    assert basket.count() == 0


def test_clearing_empties_everything(basket):
    basket.add("SKU-1", 3, 1.0)
    basket.clear()

    assert basket.count() == 0
    assert basket.total() == 0
