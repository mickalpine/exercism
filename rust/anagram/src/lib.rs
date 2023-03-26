#![allow(unused)]
use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let word = word.to_lowercase();
    let sorted_word = sort_word(&word);

    possible_anagrams
        .iter()
        .filter(|anagram| {
            let anagram = anagram.to_lowercase();
            anagram.len() == word.len() && anagram != word && sort_word(&anagram) == sorted_word
        })
        .copied()
        .collect()
}

fn sort_word(word: &str) -> Vec<char> {
    let mut word: Vec<char> = word
        .chars()
        .collect();

    word.sort_unstable();

    word
}
