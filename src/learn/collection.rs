pub fn vector() {
    let mut v: Vec<i32> = Vec::new();
    println!("{:?}", v);
    v.push(1);
    println!("{:?}", v);
    v.push(2);
    println!("{:?}", v);//[1, 2]
    v.pop();
    println!("{:?}", v);//[1]
    v.extend(&[10, 12, 13, 14]);
    println!("{:?}", v); //[1, 10, 12, 13, 14]

    println!("{} {} {}", v[0], v[1], v[2]);

    for i in &v {
        print!("{}\t", i);
    }
    println!();
    println!("{:?}", v.get(1));
    println!("{:?}", v);
    v.insert(0, 21);
    v.insert(3, 22);
    println!("{:?}", v);//[21, 1, 10, 22, 12, 13, 14]
    println!("{}", v.len()); //7
    v.reverse();
    println!("{:?}", v);//[14, 13, 12, 22, 10, 1, 21]

    let mut v2 = vec![30, 31, 32, 33];
    v.append(&mut v2);
    println!("{:?}", v);//[14, 13, 12, 22, 10, 1, 21, 30, 31, 32, 33]
    println!("{:?}", v.as_slice());

    println!("{}", v.contains(&1)); //true
    println!("{}", v.ends_with(&vec![33])); //true

    v.truncate(8);
    println!("{:?}", v);//[14, 13, 12, 22, 10, 1, 21, 30]

    println!("{:?}", v.first());//Some(14)
    println!("{:?}", v.last());//Some(30)
    println!("{}", v.is_empty()); //false

    let mut iter = v.iter();
    println!("{:?}", iter.next());//Some(14)
    println!("{:?}", iter.next());//Some(13)
}

pub fn string() {
    let string = String::new();
    println!("string -> {}", string);
    println!("{}", String::from("hello world"));
    let mut str1 = String::from("hello rust");
    str1.push('.');
    str1.push('哈');
    println!("{}", str1);
    println!("{}", str1.len());
    println!("{}", '哈'.to_string().len());
}

#[allow(unused_variables)]
pub fn string2() {

    //There are multiple ways to create a new [`String`] from a string literal:
    let s = "Hello".to_string();
    let s = String::from("world");
    let s: String = "also this".into();

    let message = s + " world!";

    let sparkle_heart = vec![240, 159, 146, 150];
    // We know these bytes are valid, so we'll use `unwrap()`.
    let sparkle_heart = String::from_utf8(sparkle_heart).unwrap();
    assert_eq!("💖", sparkle_heart);
    let bytes = sparkle_heart.into_bytes();
    assert_eq!(bytes, [240, 159, 146, 150]);

    let mut hello = String::from("Hello, ");
    hello.push('w');
    hello.push_str("orld!");
}

pub fn string3() {
    use std::mem;
    let story = String::from("Once upon a time...");
    let ptr = story.as_ptr();
    let len = story.len();
    let capacity = story.capacity();
// story has nineteen bytes
    assert_eq!(19, len);
// Now that we have our parts, we throw the story away.
    mem::forget(story);
// We can re-build a String out of ptr, len, and capacity. This is all
// unsafe because we are responsible for making sure the components are
// valid:
    let s = unsafe {
        String::from_raw_parts(ptr as *mut _, len, capacity)
    };
    assert_eq!(String::from("Once upon a time..."), s);

    let mut s = String::new();
    println!("{}", s.capacity());
    for _ in 0..5 {
        s.push_str("hello");
        println!("{}", s.capacity());
    }

    let mut s = String::with_capacity(25);
    println!("{}", s.capacity());
    for _ in 0..10 {
        s.push_str("hello");
        println!("{}", s.capacity());
    }

    /////////////////////////////////////////////
    // some invalid bytes, in a vector
    let bytes = vec![0, 159];
    let value = String::from_utf8(bytes);
    assert!(value.is_err());
    assert_eq!(vec![0, 159], value.unwrap_err().into_bytes());

    let v = &[0xD834, 0xDD1E, 0x006d, 0x0075,
        0xD800, 0x0069, 0x0063];
    assert!(String::from_utf16(v).is_err());
    println!("{}", String::from_utf16(v).is_ok());
    println!("{:?}", String::from("哈哈").into_bytes());
    println!("{}", String::from_utf8(String::from("哈哈").into_bytes()).unwrap());


    let s = String::from("hello");
    let bytes = s.into_bytes();
    println!("{:?}", bytes);
    assert_eq!(&[104, 101, 108, 108, 111][..], &bytes[..]);

    let mut s = String::from("foobar");
    let s_mut_str = s.as_mut_str();
    s_mut_str.make_ascii_uppercase();
    assert_eq!("FOOBAR", s_mut_str);


    //reserve
    let mut s = String::with_capacity(10);
    s.push('a');
    s.push('b');
// s now has a length of 2 and a capacity of 10
    assert_eq!(2, s.len());
    assert_eq!(10, s.capacity());
// Since we already have an extra 8 capacity, calling this...
    s.reserve(8);
    println!("{}", s);
// ... doesn't actually increase.
    assert_eq!(10, s.capacity());


    //push
    let mut s = String::from("abc");
    s.push('1');
    s.push('2');
    s.push('3');
    assert_eq!("abc123", s);

    //as_bytes
    let s = String::from("hello");
    assert_eq!(&[104, 101, 108, 108, 111], s.as_bytes());

    //truncate
    let mut s = String::from("hello");
    s.truncate(2);
    assert_eq!("he", s);

    //pop
    let mut s = String::from("foo");
    assert_eq!(s.pop(), Some('o'));
    assert_eq!(s.pop(), Some('o'));
    assert_eq!(s.pop(), Some('f'));
    assert_eq!(s.pop(), None);

    //remove
    let mut s = String::from("foo");
    assert_eq!(s.remove(0), 'f');
    assert_eq!(s.remove(1), 'o');
    assert_eq!(s.remove(0), 'o');
    assert_eq!(s, "");


    //retain
    let mut s = String::from("f_o_ob_ar");
    s.retain(|c| c != '_');
    assert_eq!(s, "foobar");

    let mut s = String::from("abcde");
    let keep = [false, true, true, false, true];
    let mut i = 0;
    s.retain(|_| (keep[i], i += 1).0);
    assert_eq!(s, "bce");

    //insert
    let mut s = String::with_capacity(3);
    s.insert(0, 'f');
    s.insert(1, 'o');
    s.insert(2, 'o');
    assert_eq!("foo", s);

    //insert_str
    let mut s = String::from("bar");
    s.insert_str(0, "foo");
    assert_eq!("foobar", s);

    //len
    let a = String::from("foo");
    assert_eq!(a.len(), 3);

    //is_empty
    let mut v = String::new();
    assert!(v.is_empty());
    v.push('a');
    assert!(!v.is_empty());

    //split_off
    let mut hello = String::from("Hello, World!");
    let world = hello.split_off(7);
    assert_eq!(hello, "Hello, ");
    assert_eq!(world, "World!");

    //clear
    let mut s = String::from("foo");
    s.clear();
    assert!(s.is_empty());
    assert_eq!(0, s.len());
    assert_eq!(3, s.capacity());

    //drain
    let mut s = String::from("α is alpha, β is beta");
    let beta_offset = s.find('β').unwrap_or(s.len());
    // Remove the range up until the β from the string
    let t: String = s.drain(..beta_offset).collect();
    assert_eq!(t, "α is alpha, ");
    assert_eq!(s, "β is beta");
    // A full range clears the string
    s.drain(..);
    assert_eq!(s, "");


    //replace_range
    let mut s = String::from("α is alpha, β is beta");
    let beta_offset = s.find('β').unwrap_or(s.len());
    // Replace the range up until the β from the string
    s.replace_range(0..beta_offset, "Α is capital alpha; ");
    assert_eq!(s, "Α is capital alpha; β is beta");

    //as_bytes
    let bytes = vec![0, 159];
    let value = String::from_utf8(bytes);
    assert_eq!(&[0, 159], value.unwrap_err().as_bytes());

    //into_bytes
    // some invalid bytes, in a vector
    let bytes = vec![0, 159];
    let value = String::from_utf8(bytes);
    assert_eq!(vec![0, 159], value.unwrap_err().into_bytes());

    //utf8_error
    // some invalid bytes, in a vector
    let bytes = vec![0, 159];
    let error = String::from_utf8(bytes).unwrap_err().utf8_error();
    // the first byte is invalid here
    assert_eq!(1, error.valid_up_to());
}