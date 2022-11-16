use std::{io, process::Stdio, fmt::Error};

fn main() {
    println!("Hello, world!");
    println!("Please input your guess: ");
    let mut guess: String = String::new();

    io::stdin().read_line(&mut guess)
        .expect("Failed to read line!");
    
    println!("Your guessed: {}",guess);
}
