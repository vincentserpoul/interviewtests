use std::vec::Vec;
struct Solution {}

impl Solution {
    pub fn count_smaller(&self, nums: Vec<i32>) -> Vec<i32> {
        // contains the count of bigger right elems
        let mut count = Vec::with_capacity(nums.len());
        // contains the ordered nums
        let mut rev = Vec::with_capacity(nums.len());
        nums.iter().rev().for_each(|val| {
            // find where to place in count (bin search, then insert)
            let rev_idx = rev.binary_search(val).unwrap_or_else(|x| x);
            rev.insert(rev_idx, *val);

            // this index describe the number of nums on the right
            count.push(rev_idx as i32);
        });
        count.into_iter().rev().collect()
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn solution() {
        let sol = Solution {};
        assert_eq!(sol.count_smaller(vec![5, 2, 6, 1]), vec![2, 1, 1, 0]);
        assert_eq!(sol.count_smaller(vec![5, -2, 6, 1]), vec![2, 0, 1, 0]);
    }
}
