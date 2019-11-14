fn main() {
    let x = 5;
    println!("The value of x is: {}", x);

    let x = x + 1;
    println!("The value of x is: {}", x);

    let x = x * 2;

    println!("The value of x is: {}", x);

    let spaces = "   ";
    println!("The value of x is: {}", spaces);
    let spaces = spaces.len();
    println!("The value of x is: {}", spaces);

    let guess: u32 = "42".parse().expect("Not a number!");
    println!("{0}",guess);

    let x = 2.0; // f64
    let y: f32 = 3.0; // f32
    let t = true;
    let f: bool = false; // with explicit type annotation

    let c = 'z';
    let z = 'ℤ';
    let heart_eyed_cat = '😻';
    println!("The value of x is: {}", heart_eyed_cat);

    let tup: (i32, f64, u8) = (500, 6.4, 1);
    println!("The value of x is: {}", tup.0);
    println!("The value of x is: {}", tup.1);
    println!("The value of x is: {}", tup.2);

    let tup = (500, 6.4, 1);

    let (x, y, z) = tup;

    println!("The value of y is: {}", y);


    let a = [1, 2, 3, 4, 5];
    let months = ["January", "February", "March", "April", "May", "June", "July",
        "August", "September", "October", "November", "December"];
    let a: [i32; 5] = [1, 2, 3, 4, 5];
    println!("The value of y is: {}", a[0]);
    println!("The value of y is: {}", months[0]);
    println!("The value of y is: {}", months.len());

    another_function(123456789);


    let x = 5;

    let y = {
        let x = 3;
        x + 1
    };

    println!("The value of y is: {}", y);

    println!("The value of y is: {}", five());


    let condition = true;
    let number = if condition {
        5
    } else {
        6
    };

    println!("The value of number is: {}", number);

    let a = [10, 20, 30, 40, 50];

    for element in a.iter() {
        println!("the value is: {}", element);
    }

    for number in (1..4).rev() {
        println!("{}!", number);
    }
    println!("LIFTOFF!!!");

}


fn another_function(x: i32) {
    println!("The value of x is: {}", x);
}
fn five() -> i32 {
    5
}
