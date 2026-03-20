use pyo3::prelude::*;

#[pyfunction]
fn sum_as_string(a: usize, b: usize) -> PyResult<String> {
    Ok((a + b).to_string())
}

#[pymodule]
fn melnikova_a(m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(sum_as_string, m)?)?;
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_rust_sum() {
        assert_eq!(sum_as_string(2, 2).unwrap(), "4");
        assert_eq!(sum_as_string(0, 0).unwrap(), "0");
    }
}