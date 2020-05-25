struct HitCounter {
    hits: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl HitCounter {
    /** Initialize your data structure here. */
    fn new() -> Self {
        HitCounter { hits: Vec::new() }
    }

    /** Record a hit.
    @param timestamp - The current timestamp (in seconds granularity). */
    fn hit(&mut self, timestamp: i32) {
        self.hits.push(timestamp);
    }

    /** Return the number of hits in the past 5 minutes.
    @param timestamp - The current timestamp (in seconds granularity). */
    fn get_hits(&mut self, timestamp: i32) -> i32 {
        let min_timestamp = timestamp - 300 + 1;
        let start_idx = match self.hits.binary_search(&min_timestamp) {
            Ok(mut idx) => {
                while idx > 0 && self.hits[idx - 1] == self.hits[idx] {
                    idx -= 1;
                }
                idx
            }
            Err(idx) => idx,
        };
        self.hits = self.hits[start_idx..].to_vec();
        self.hits.len() as i32
    }
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * let obj = HitCounter::new();
 * obj.hit(timestamp);
 * let ret_2: i32 = obj.get_hits(timestamp);
 */

#[cfg(test)]
mod tests {
    use super::HitCounter;

    #[test]
    fn hit_counter_example() {
        let mut hc = HitCounter::new();
        assert_eq!(hc.get_hits(0), 0, "get_hits(0)");
        hc.hit(1);
        hc.hit(2);
        hc.hit(3);
        assert_eq!(hc.get_hits(4), 3, "get_hits(4)");
        hc.hit(300);
        assert_eq!(hc.get_hits(300), 4, "get_hits(300)");
        assert_eq!(hc.get_hits(301), 3, "get_hits(301)");
    }

    #[test]
    fn hit_counter_repeat() {
        let mut hc = HitCounter::new();
        assert_eq!(hc.get_hits(0), 0, "get_hits(0)");
        hc.hit(1);
        hc.hit(1);
        hc.hit(1);
        assert_eq!(hc.get_hits(4), 3, "get_hits(4)");
        hc.hit(300);
        assert_eq!(hc.get_hits(300), 4, "get_hits(300)");
        assert_eq!(hc.get_hits(301), 1, "get_hits(301)");
    }

    #[test]
    fn hit_counter_sample1() {
        let mut hc = HitCounter::new();
        hc.hit(100);
        hc.hit(282);
        hc.hit(411);
        hc.hit(609);
        hc.hit(620);
        hc.hit(744);
        assert_eq!(hc.get_hits(879), 3, "get_hits(879)");
        hc.hit(956);
        assert_eq!(hc.get_hits(976), 2, "get_hits(976)");
        hc.hit(998);
        assert_eq!(hc.get_hits(1055), 2, "get_hits(976)");
    }
}
