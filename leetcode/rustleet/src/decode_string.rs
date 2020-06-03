use std::vec::Vec;

struct Solution {}

impl Solution {
    pub fn decode_string(&self, s: String) -> String {
        // println!("{:?}", s);

        let mut stack_n: Vec<u32> = vec![1];
        let mut stack_w: Vec<String> = vec![String::from("")];
        let mut tmp_n: Vec<u32> = Vec::new();
        let mut tmp_w: Vec<char> = Vec::new();

        s.chars().for_each(|c| {
            let r = match c {
                c if c.is_digit(10) => {
                    tmp_n.push(c.to_digit(10).unwrap());
                }
                c if c.is_alphabetic() => {
                    tmp_w.push(c);
                }
                c if c.to_string() == String::from("[") => {
                    if let Some(w) = flatten_w(&mut tmp_w) {
                        let last_idx_w = stack_w.len() - 1;
                        stack_w[last_idx_w].push_str(&w);
                    }
                    let n = flatten_n(&mut tmp_n).expect("no multiplier available");
                    stack_n.push(n);
                    stack_w.push(String::from(""));
                }
                c if c.to_string() == String::from("]") => {
                    if let Some(w) = flatten_w(&mut tmp_w) {
                        let last_idx_w = stack_w.len() - 1;
                        stack_w[last_idx_w].push_str(&w);
                    }
                    if stack_w.len() == stack_n.len() {
                        let curr_n = stack_n.pop();
                        let curr_w = stack_w.pop();
                        let nw = combine_nw(curr_n, curr_w.clone());
                        let last_w_idx = stack_w.len() - 1;
                        stack_w[last_w_idx].push_str(&nw);
                    }
                }
                _ => panic!("wrong char {:?}", c),
            };
            // println!(
            //     "{:?} -> {:?} {:?} || {:?} {:?}",
            //     c, tmp_n, tmp_w, stack_n, stack_w
            // );
            r
        });
        // left over
        if let Some(w) = flatten_w(&mut tmp_w) {
            let last_idx_w = stack_w.len() - 1;
            stack_w[last_idx_w].push_str(&w);
        }
        // println!(
        //     "    -> {:?} {:?} || {:?} {:?}",
        //     tmp_n, tmp_w, stack_n, stack_w
        // );

        while stack_w.len() > 1 {
            let curr_n = stack_n.pop();
            let curr_w = stack_w.pop();
            let nw = combine_nw(curr_n, curr_w.clone());
            let last_w_idx = stack_w.len() - 1;
            stack_w[last_w_idx].push_str(&nw);
        }
        stack_w[0].clone()
    }
}

fn flatten_n(tmp_n: &mut Vec<u32>) -> Option<u32> {
    if tmp_n.len() == 0 {
        return None;
    }
    let n = tmp_n.iter().rev().enumerate().fold(0, |mut acc, d| {
        acc += d.1 * 10u32.pow(d.0 as u32);
        acc
    });
    tmp_n.truncate(0);
    Some(n)
}

fn flatten_w(tmp_w: &mut Vec<char>) -> Option<String> {
    if tmp_w.len() == 0 {
        return None;
    }
    let w = tmp_w.iter().collect::<String>();
    tmp_w.truncate(0);
    Some(w)
}

fn combine_nw(opt_curr_n: Option<u32>, opt_curr_w: Option<String>) -> String {
    let curr_w = opt_curr_w.unwrap_or(String::from(""));
    let mut res = String::from("");
    let mut curr_n = opt_curr_n.unwrap_or(1);

    while curr_n != 0 {
        res.push_str(&curr_w.clone());
        curr_n -= 1;
    }

    res
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn solution_example() {
        let s = Solution {};
        assert_eq!(s.decode_string(String::from("3[a]")), String::from("aaa"));
        assert_eq!(
            s.decode_string(String::from("3[a2[c]]")),
            String::from("accaccacc")
        );
        assert_eq!(
            s.decode_string(String::from("3[a]2[bc]")),
            String::from("aaabcbc")
        );
        assert_eq!(
            s.decode_string(String::from("2[abc]3[cd]ef")),
            String::from("abcabccdcdcdef")
        );
        assert_eq!(
            s.decode_string(String::from("10[a]")),
            String::from("aaaaaaaaaa")
        );
        assert_eq!(
            s.decode_string(String::from("3[a]2[b4[F]c]")),
            String::from("aaabFFFFcbFFFFc")
        );
        assert_eq!(
            s.decode_string(String::from("2[2[y]]")),
            String::from("yyyy")
        );
        assert_eq!(
            s.decode_string(String::from("1[j]e2[f]")),
            String::from("jeff")
        );
        assert_eq!(
            s.decode_string(String::from("2[1[j]e2[f]]")),
            String::from("jeffjeff")
        );
        assert_eq!(
            s.decode_string(String::from("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")),
            String::from("zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef")
        );
    }
}
