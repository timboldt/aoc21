use std::error::Error;
use std::fs;
use std::path::Path;

mod day01;

fn main() -> Result<(), Box<dyn Error>> {
    // Parse arguments.
    let path_name = std::env::args().nth(1).ok_or("no path given")?;
    let day = std::env::args()
        .nth(2)
        .ok_or("no day given")?
        .parse::<i32>()?;

    let path = Path::new(&path_name)
        .join(format!("day{:02}", day))
        .join("input.txt");
    let input = fs::read_to_string(path)?;
    println!(
        "Day {}:\n\tPart 1: {}\n\tPart 2: {}\n",
        day,
        match day {
            1 => day01::part1(&input),
            _ => "not implemented yet".to_string(),
        },
        match day {
            1 => day01::part2(&input),
            _ => "not implemented yet".to_string(),
        }
    );
    Ok(())
}
