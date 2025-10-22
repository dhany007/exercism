use std::fmt;

#[derive(Debug, PartialEq, Eq)]
pub struct Clock {
    hours: i32,
    minutes: i32
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let all_minute = (hours * 60) + minutes;

        let day_in_minute = 24 * 60;
        let mut minute = all_minute % day_in_minute;

        if minute < 0 {
            minute += day_in_minute
        }
        
        Clock {
            hours: minute/60,
            minutes: minute%60
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours, self.minutes+minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}