fn main() {
    let s1 = String::from("hello");

    let (s2, len) = calculate_length(s1);

    println!("The length of '{}' is {}.", s2, len);


    let s = "hello";

    {                      // s is not valid here, it’s not yet declared
        println!("{}",s);
        let s = "hello rust";   // s is valid from this point forward
        println!("{}",s);

        // do stuff with s
    }


    {
        let s = String::from("hello");
        let mut s = String::from("hello");

        s.push_str(", world!"); // push_str() appends a literal to a String

        println!("{}", s); // This will print `hello, world!`

    }

    let s1 = String::from("hello");

    let len = calculate_length2(&s1);

    println!("The length of '{}' is {}.", s1, len);

    {
        let mut s = String::from("hello");

        change(&mut s);

        println!("The length of '{}' ", s);
    }
}
fn change(some_string: &mut String) {
    some_string.push_str(", world");
}


fn calculate_length2(s: &String) -> usize {
    s.len()
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len(); // len() returns the length of a String

    (s, length)
}