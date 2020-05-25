struct MedianFinder {
    data: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MedianFinder {
    /** initialize your data structure here. */
    fn new() -> Self {
        MedianFinder { data: Vec::new() }
    }

    fn add_num(&mut self, num: i32) {
        let idx = self.data.binary_search(&num).unwrap_or_else(|x| x);
        self.data.insert(idx, num);
    }

    fn find_median(&self) -> f64 {
        if self.data.len() == 0 {
            return 0.0;
        }
        if self.data.len() == 1 {
            return self.data[0] as f64;
        }
        if self.data.len() % 2 != 0 {
            return self.data[(self.data.len() - 1) / 2] as f64;
        }
        ((self.data[(self.data.len() / 2) - 1] + self.data[(self.data.len() / 2)]) as f64) / 2.0
    }
}

#[cfg(test)]
mod tests {
    use super::MedianFinder;

    #[test]
    fn median_finder() {
        // [[],[1],[2],[],[3],[]]
        let mut mf = MedianFinder::new();
        assert_eq!(mf.find_median(), 0.0);
        mf.add_num(1);
        mf.add_num(2);
        assert_eq!(mf.find_median(), 1.5);
        mf.add_num(3);
        assert_eq!(mf.find_median(), 2.0);
    }
}
