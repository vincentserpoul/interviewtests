use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::collections::HashMap;

struct Logger {
    message_last_timestamp: HashMap<String, i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Logger {
    /** Initialize your data structure here. */
    fn new() -> Self {
        Logger {
            message_last_timestamp: HashMap::new(),
        }
    }

    /** Returns true if the message should be printed in the given timestamp, otherwise returns false.
    If this method returns false, the message will not be printed.
    The timestamp is in seconds granularity. */
    fn should_print_message(&mut self, timestamp: i32, message: String) -> bool {
        match self.message_last_timestamp.entry(message) {
            Occupied(mut entry) => {
                if timestamp - *entry.get() >= 10 {
                    entry.insert(timestamp);
                    return true;
                }
                return false;
            }
            Vacant(entry) => {
                entry.insert(timestamp);
                return true;
            }
        }
    }
}

/**
 * Your Logger object will be instantiated and called as such:
 * let obj = Logger::new();
 * let ret_1: bool = obj.should_print_message(timestamp, message);
 */

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn logger() {
        let mut l = Logger::new();
        assert_eq!(l.should_print_message(1, "test".to_string()), true);
        assert_eq!(l.should_print_message(1, "test".to_string()), false);
    }
}
