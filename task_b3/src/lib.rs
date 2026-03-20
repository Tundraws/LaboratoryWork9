use pyo3::prelude::*;

#[pyfunction]
#[pyo3(name = "calculate_sum")]
fn calculate_sum_fn(first_value: i32, second_value: i32) -> i32 {
    first_value + second_value
}

#[pyfunction]
#[pyo3(name = "get_reversed_string")]
fn get_reversed_string_fn(input_data: String) -> String {
    input_data.chars().rev().collect()
}

#[pymodule]
fn melnikova_a(m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(calculate_sum_fn, m)?)?;
    m.add_function(wrap_pyfunction!(get_reversed_string_fn, m)?)?;
    Ok(())
}