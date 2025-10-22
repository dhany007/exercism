// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

#[derive(Debug)]
pub struct Duration {
    seconds: u64
}

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Duration {
            seconds: s
        }
    }
}

pub trait Planet {
    fn years_during(d: &Duration) -> f64;
}

const EARTH_YEAR_IN_DAYS: f64 = 365.256;
const EARTH_YEAR_IN_SECONDS: f64 = EARTH_YEAR_IN_DAYS * 24.0 * 60.0 * 60.0;

pub struct Mercury;
pub struct Venus;
pub struct Earth;
pub struct Mars;
pub struct Jupiter;
pub struct Saturn;
pub struct Uranus;
pub struct Neptune;

const MERCURY_ORBITAL_PERIOD: f64 = 0.2408467;
impl Planet for Mercury {
    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / (EARTH_YEAR_IN_SECONDS * MERCURY_ORBITAL_PERIOD)
    }
}

const VENUS_ORBITAL_PERIOD: f64 = 0.61519726;
impl Planet for Venus {
    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / (EARTH_YEAR_IN_SECONDS * VENUS_ORBITAL_PERIOD)
    }
}

const EARTH_ORBITAL_PERIOD: f64 = 1.0;
impl Planet for Earth {
    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / (EARTH_YEAR_IN_SECONDS * EARTH_ORBITAL_PERIOD)
    }
}

const MARS_ORBITAL_PERIOD: f64 = 1.8808158;
impl Planet for Mars {
    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / (EARTH_YEAR_IN_SECONDS * MARS_ORBITAL_PERIOD)
    }
}

const JUPTER_ORBITAL_PERIOD: f64 = 11.862615;
impl Planet for Jupiter {
    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / (EARTH_YEAR_IN_SECONDS * JUPTER_ORBITAL_PERIOD)
    }
}

const SATURN_ORBITAL_PERIOD: f64 = 29.447498;
impl Planet for Saturn {
    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / (EARTH_YEAR_IN_SECONDS * SATURN_ORBITAL_PERIOD)
    }
}

const URANUS_ORBITAL_PERIOD: f64 = 84.016846;
impl Planet for Uranus {
    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / (EARTH_YEAR_IN_SECONDS * URANUS_ORBITAL_PERIOD)
    }
}

const NEPTUNE_ORBITAL_PERIOD: f64 = 164.79132;
impl Planet for Neptune {
    fn years_during(d: &Duration) -> f64 {
        d.seconds as f64 / (EARTH_YEAR_IN_SECONDS * NEPTUNE_ORBITAL_PERIOD)
    }
}
