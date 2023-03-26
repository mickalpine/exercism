pub fn is_armstrong_number(num: u32) -> bool {
    let num_digits = num
        .to_string()
        .len() as u32;

    if num_digits >= 10 {
        return false;
    }

    let num_sum = num
        .to_string()
        .chars()
        .map(|digit| {
            digit
                .to_digit(10)
                .unwrap()
                .pow(num_digits)
        })
        .sum();
    num == num_sum
}
