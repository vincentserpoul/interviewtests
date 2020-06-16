use std::cmp;
use std::vec::Vec;
struct Solution {}

impl Solution {
    pub fn max_sum_two_no_overlap(&self, a: Vec<i32>, l: i32, m: i32) -> i32 {
        cmp::max(
            max_sum_two_no_overlap_lr(&a, l as usize, m as usize),
            max_sum_two_no_overlap_lr(&a, m as usize, l as usize),
        )
    }
}

fn max_sum_two_no_overlap_lr(a: &Vec<i32>, size_l: usize, size_r: usize) -> i32 {
    (0..=a.len() - size_l)
        .map(|l_idx| {
            let sum_l: i32 = a[l_idx..l_idx + size_l].iter().sum();

            // for each of the right of the array, calculate the max sum
            let max_r = (l_idx + size_l..=a.len() - size_r)
                .map(|r_idx| a[r_idx..r_idx + size_r].iter().sum())
                .max()
                .unwrap_or(0);

            sum_l + max_r
        })
        .max()
        .unwrap_or(0)
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn solution() {
        let sol = Solution {};
        assert_eq!(
            sol.max_sum_two_no_overlap(vec![0, 6, 5, 2, 2, 5, 1, 9, 4], 1, 2),
            20
        );
        assert_eq!(
            sol.max_sum_two_no_overlap(vec![3, 8, 1, 3, 2, 1, 8, 9, 0], 3, 2),
            29
        );
        assert_eq!(
            sol.max_sum_two_no_overlap(vec![2, 1, 5, 6, 0, 9, 5, 0, 3, 8], 4, 3),
            31
        );
    }
}
