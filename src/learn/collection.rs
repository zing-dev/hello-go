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
    println!("{:?}",v);//[14, 13, 12, 22, 10, 1, 21, 30, 31, 32, 33]
    println!("{:?}",v.as_slice());

    println!("{}",v.contains(&1)); //true
    println!("{}",v.ends_with(&vec![33])); //true

    v.truncate(8);
    println!("{:?}",v);//[14, 13, 12, 22, 10, 1, 21, 30]

    println!("{:?}",v.first());//Some(14)
    println!("{:?}",v.last());//Some(30)
    println!("{}",v.is_empty()); //false

    let mut iter = v.iter();
    println!("{:?}",iter.next());//Some(14)
    println!("{:?}",iter.next());//Some(13)
}

pub fn string(){
    let string = String::new();
    println!("string -> {}",string);
    println!("{}",String::from("hello world"));
    let mut str1 = String::from("hello rust");
    str1.push('.');
    str1.push('哈');
    println!("{}",str1);
    println!("{}",str1.len());
    println!("{}",'哈'.to_string().len());

}