pub fn if_else() {
    let n = 5;

    if n < 0 {
        print!("{} is negative", n);
    } else if n > 0 {
        print!("{} is positive", n);
    } else {
        print!("{} is zero", n);
    }

    let big_n =
        if n < 10 && n > -10 {
            println!(", and is a small number, increase ten-fold");

            // This expression returns an `i32`.
            10 * n
        } else {
            println!(", and is a big number, halve the number");

            // This expression must return an `i32` as well.
            n / 2
            // TODO ^ Try suppressing this expression with a semicolon.
        };
    //   ^ Don't forget to put a semicolon here! All `let` bindings need it.

    println!("{} -> {}", n, big_n);
}

pub fn loop_test() {
    let mut count = 0u32;

    println!("Let's count until infinity!");

// Infinite loop
    loop {
        count += 1;

        if count == 3 {
            println!("three");

// Skip the rest of this iteration
            continue;
        }

        println!("{}", count);

        if count == 5 {
            println!("OK, that's enough");

// Exit this loop
            break;
        }
    }
}

#[allow(unreachable_code)]
pub fn loop_nested_label() {
    'outer: loop {
        println!("Entered the outer loop");

        'inner: loop {
            println!("Entered the inner loop");

            // This would break only the inner loop
            //break;

            // This breaks the outer loop
            break 'outer;
        }

        println!("This point will never be reached");
    }

    println!("Exited the outer loop");
}

pub fn loop_return() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("counter is {}", result)
}

pub fn while_test() {
    // A counter variable
    let mut n = 1;

    // Loop while `n` is less than 101
    while n < 21 {
        if n % 15 == 0 {
            println!("n % 15 == 0, n is {}", n);
        } else if n % 3 == 0 {
            println!("n % 3 == 0, n is {}", n);
        } else if n % 5 == 0 {
            println!("n % 5 == 0, n is {}", n);
        } else {
            println!("{}", n);
        }

        // Increment counter
        n += 1;
    }
    println!("n is {}", n)
}

pub fn for_range() {
    fn handle(n: i32) {
        if n % 15 == 0 {
            println!("fizzbuzz");
        } else if n % 3 == 0 {
            println!("fizz");
        } else if n % 5 == 0 {
            println!("buzz");
        } else {
            println!("{}", n);
        }
    }
    // `n` will take the values: 1, 2, ..., 10 in each iteration
    for n in 1..11 {
        handle(n)
    }
    // `n` will take the values: 1, 2, ..., 10 in each iteration
    for n in 1..=10 {
        handle(n)
    }
}

pub fn for_iter() {
    let names = vec!["Bob", "Frank", "Ferris"];

    for name in names.into_iter() {
        match name {
            "Ferris" => println!("There is a rustacean among us!"),
            _ => println!("Hello {}", name),
        }
        if name == "Ferris" {
            println!("There is a rustacean among us!");
        } else {
            println!("Hello {}", name);
        }
    }
}

pub fn for_iter_mut() {
    let mut names = vec!["Bob", "Frank", "Ferris"];

    for name in names.iter_mut() {
        *name = match name {
            &mut "Ferris" => "There is a rustacean among us!",
            _ => "Hello",
        }
    }

    println!("names: {:?}", names);
}
