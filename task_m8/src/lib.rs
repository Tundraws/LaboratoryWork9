use pyo3::prelude::*;

#[pyfunction]
fn calculate_sum_of_squares(numbers: Vec<i32>) -> PyResult<i32> {
    let total: i32 = numbers.iter().map(|&x| x * x).sum();
    Ok(total)
}

#[pymodule]
fn task_m8(m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(calculate_sum_of_squares, m)?)?;
    Ok(())
}