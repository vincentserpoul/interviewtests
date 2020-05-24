use std::vec::Vec;

struct Solution {}

impl Solution {
    pub fn largest_rectangle_area(&self, heights: Vec<i32>) -> i32 {
        let mut s: Vec<(usize, &i32)> = Vec::new();
        let mut max_area: i32 = 0;
        heights
            .iter()
            .enumerate()
            .for_each(|(curr_idx, curr_height)| {
                // println!("{:?} {:?}", (curr_idx, curr_height), s);
                let mut lowest_idx = curr_idx;
                // pop stack while topstack height >= curr_height
                while s.len() > 0 && s[s.len() - 1].1 >= curr_height {
                    let curr_top_stack = s.pop().unwrap();
                    lowest_idx = curr_top_stack.0;
                    max_area = calc_area(max_area, curr_idx, curr_top_stack);
                    // println!(
                    //     "calculating area for {:?} {:?}, max area: {:?}",
                    //     curr_idx, curr_top_stack, max_area,
                    // );
                }
                s.push((lowest_idx, curr_height));
            });
        s.iter().rev().for_each(|curr_top_stack| {
            max_area = calc_area(max_area, heights.len(), *curr_top_stack);
        });
        max_area
    }
}

pub fn calc_area(max_area: i32, cur_idx: usize, bar: (usize, &i32)) -> i32 {
    let area = (cur_idx - bar.0) as i32 * bar.1;
    if area > max_area {
        return area;
    }
    max_area
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn solution() {
        let sol = Solution {};
        assert_eq!(sol.largest_rectangle_area(vec![]), 0);
        assert_eq!(sol.largest_rectangle_area(vec![4, 2]), 4);
        assert_eq!(sol.largest_rectangle_area(vec![2, 3]), 4);
        assert_eq!(sol.largest_rectangle_area(vec![1]), 1);
        assert_eq!(sol.largest_rectangle_area(vec![1, 1, 1, 1, 1, 1]), 6);
        assert_eq!(sol.largest_rectangle_area(vec![2, 1, 5, 6, 2, 3]), 10);
        assert_eq!(sol.largest_rectangle_area(vec![2, 1, 2]), 3);
        assert_eq!(sol.largest_rectangle_area(vec![2, 0, 3, 2]), 4)
    }
}
