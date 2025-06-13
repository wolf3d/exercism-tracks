pub fn production_rate_per_hour(speed: u8) -> f64 {    
    success_rate_percent(speed) * (speed as f64 * 221.00_f64 )
}

pub fn working_items_per_minute(speed: u8) -> u32 {    
    (production_rate_per_hour(speed) / 60_f64) as u32    
}

fn success_rate_percent(speed: u8) -> f64 {
    match speed {
        1 | 2 | 3 | 4 => 1.00,
        5 | 6 | 7 | 8 => 0.90,
        9 | 10 => 0.77,
        _ => 0.0
    }
}