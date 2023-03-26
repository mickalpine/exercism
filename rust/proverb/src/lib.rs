pub fn build_proverb(list: &[&str]) -> String {
    match list.first() {
        None => String::new(),
        Some(&word) => list
            .windows(2)
            .map(proverb_line)
            .chain(std::iter::once(final_line(word)))
            .collect(),
    }
}

pub fn proverb_line(words: &[&str]) -> String {
    format!("For want of a {} the {} was lost.\n", words[0], words[1])
}

pub fn final_line(word: &str) -> String {
    format!("And all for the want of a {}.", word)
}
