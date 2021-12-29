use std::num::ParseIntError;

fn parse_input(input: &str) -> Result<Vec<i32>, ParseIntError> {
    input.lines().map(|l| l.parse()).collect()
}

pub fn part1(input: &str) -> String {
    let depths = match parse_input(input) {
        Ok(x) => x,
        Err(e) => return e.to_string(),
    };
    let mut count = 0;
    let mut prev = 0;
    for (i, &depth) in depths.iter().enumerate() {
        if i > 0 && depth > prev {
            count += 1;
        }
        prev = depth;
    }
    count.to_string()
}

pub fn part2(input: &str) -> String {
    let depths = match parse_input(input) {
        Ok(x) => x,
        Err(e) => return e.to_string(),
    };
    let mut count = 0;
    let mut prev = 0;
    for (i, _) in depths[0..depths.len() - 2].iter().enumerate() {
        let val = depths[i] + depths[i + 1] + depths[i + 2];
        if i > 0 && val > prev {
            count += 1;
        }
        prev = val;
    }
    count.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    static TEST_INPUT: &str = r#"199
200
208
210
200
207
240
269
260
263"#;

    #[test]
    fn part1_example() {
        assert_eq!(part1(TEST_INPUT), "7");
    }

    #[test]
    fn part2_example() {
        assert_eq!(part2(TEST_INPUT), "5");
    }
}
