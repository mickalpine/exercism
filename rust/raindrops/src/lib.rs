#![allow(unused)]

use std::vec;
pub fn raindrops(n: u32) -> String {
    let mut drops: Vec<String> = vec![];
    let factors = [3, 5, 7];

    if factors
        .into_iter()
        .all(|factor| n % factor != 0)
    {
        return n.to_string();
    }

    for i in factors {
        if n % i == 0 {
            match i {
                3 => drops.push("Pling".to_string()),
                5 => drops.push("Plang".to_string()),
                7 => drops.push("Plong".to_string()),
                _ => drops.push(n.to_string()),
            }
        }
    }

    drops
        .into_iter()
        .collect()
}
