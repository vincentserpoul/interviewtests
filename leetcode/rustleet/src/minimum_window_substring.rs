use std::collections::hash_map::Entry::Occupied;
use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn min_window(&self, s: String, t: String) -> String {
        if s.len() < t.len() {
            return String::from("");
        }
        let target_hm = t.chars().fold(HashMap::new(), |mut acc, c| {
            *acc.entry(c).or_insert(0) += 1;
            acc
        });

        let s_vec: Vec<char> = s.chars().collect();
        let filtered_s: Vec<(usize, char)> = s
            .chars()
            .enumerate()
            .collect::<Vec<(usize, char)>>()
            .into_iter()
            .filter(|(_, c)| target_hm.get(c).is_some())
            .collect();

        let mut min_window_chars: Option<(usize, usize)> = None;
        let mut curr_hm: HashMap<char, u32> = HashMap::new();
        let mut l_idx = 0 as usize;
        let mut r_idx = 0 as usize;

        while r_idx < filtered_s.len() {
            let curr_char = filtered_s[r_idx];
            *curr_hm.entry(curr_char.1).or_insert(0) += 1;

            // we move the l_idx til we still have a complete set of letters with occurences
            while hm_is_valid(&target_hm, &curr_hm) && l_idx <= r_idx {
                // if we don't have a window yet, or the length of the curr window is too small
                match min_window_chars {
                    None => {
                        min_window_chars = Some((filtered_s[l_idx].0, filtered_s[r_idx].0));
                    }
                    Some(mwc) => {
                        if (mwc.1 - mwc.0) > (filtered_s[r_idx].0 - filtered_s[l_idx].0) {
                            min_window_chars = Some((filtered_s[l_idx].0, filtered_s[r_idx].0));
                        }
                    }
                }
                match curr_hm.entry(filtered_s[l_idx].1) {
                    Occupied(mut e) => {
                        let v = e.get_mut();
                        if v > &mut 1 {
                            *v -= 1;
                        } else {
                            e.remove();
                        }
                    }
                    _ => panic!("char vec curr_hm filled the wrong way"),
                };
                // dbg!(
                //     &l_idx,
                //     &s_vec[filtered_s[l_idx].0..=filtered_s[r_idx].0]
                //         .to_vec()
                //         .iter()
                //         .collect::<String>(),
                //     &curr_hm,
                // );
                l_idx += 1;
            }
            r_idx += 1;
        }
        match min_window_chars {
            None => String::from(""),
            Some(mwc) => s_vec[mwc.0..=mwc.1].to_vec().iter().collect(),
        }
    }
}

pub fn hm_is_valid(target_hm: &HashMap<char, u32>, test_hm: &HashMap<char, u32>) -> bool {
    if target_hm == test_hm {
        return true;
    }

    for (c, count) in target_hm.iter() {
        match test_hm.get(c) {
            None => return false,
            Some(test_count) if test_count < count => return false,
            _ => (),
        }
    }

    true
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn solution() {
        let sol = Solution {};
        assert_eq!(
            sol.min_window(String::from("ADOBECODEBANC"), String::from("ABC")),
            "BANC"
        );
        assert_eq!(sol.min_window(String::from(""), String::from("ABC")), "");
        assert_eq!(sol.min_window(String::from("RTY"), String::from("ABC")), "");
        assert_eq!(
            sol.min_window(String::from("AAAAAAAAA"), String::from("A")),
            "A"
        );
        assert_eq!(sol.min_window(String::from("a"), String::from("aa")), "");
        assert_eq!(sol.min_window(String::from("aa"), String::from("aa")), "aa");
        assert_eq!(
            sol.min_window(String::from("bba"), String::from("ab")),
            "ba"
        );
    }
}
