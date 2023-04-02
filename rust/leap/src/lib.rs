pub fn is_leap_year(year: u64) -> bool {
    if year % 400 == 0 {
        return true;
    }

    if year % 100 == 0 {
        return false;
    }

    if year % 4 == 0 {
        return true;
    }

    return false;
}

pub fn is_leap_year_using_match1(year: u64) -> bool {
    match (year % 400, year % 100, year % 4) {
        (0, _, _) => true,
        (_, 0, _) => false,
        (_, _, 0) => true,
        _ => false,
    }
}

pub fn is_leap_year_using_match2(year: u64) -> bool {
    match year {
        x if x % 400 == 0 => true,
        x if x % 100 == 0 => false,
        x if x % 4 == 0 => true,
        _ => false,
    }
}
