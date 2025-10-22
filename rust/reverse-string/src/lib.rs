pub fn reverse(input: &str) -> String {
    // 1
    // let mut result = String::new();

    // for c in input.chars().rev() {
    //     result.push(c)
    // }

    // result

    // 2
    input.chars().rev().collect()
}
