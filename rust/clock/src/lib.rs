#[deny(clippy::all)]

const MINUTES_IN_DAY: i32 = 24 * 60;
const MINUTES_IN_HOUR: i32 = 60;

#[derive(Eq, Debug, PartialEq)]
pub struct Clock {
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        Self {
            minutes: (((hours * MINUTES_IN_HOUR + minutes) % MINUTES_IN_DAY) + MINUTES_IN_DAY)
                % MINUTES_IN_DAY,
        }
    }

    pub fn add_minutes(self, minutes: i32) -> Self {
        Self::new(0, self.minutes + minutes)
    }
}

impl std::fmt::Display for Clock {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(
            f,
            "{:02}:{:02}",
            self.minutes / MINUTES_IN_HOUR,
            self.minutes % MINUTES_IN_HOUR
        )
    }
}
