pub fn hof() {
    fn is_odd(n: u32) -> bool {
        n % 2 == 1
    }

    println!("Find the sum of all the squared odd numbers under 1000");
    let upper = 1000;

    // Imperative approach
    // Declare accumulator variable
    let mut acc = 0;
    // Iterate: 0, 1, 2, ... to infinity
    for n in 0.. {
        // Square the number
        let n_squared = n * n;

        if n_squared >= upper {
            // Break loop if exceeded the upper limit
            break;
        } else if is_odd(n_squared) {
            // Accumulate value, if it's odd
            acc += n_squared;
        }
    }
    println!("imperative style: {}", acc);

    // Functional approach
    let sum_of_squared_odd_numbers: u32 =
        (0..).map(|n| n * n)                             // All natural numbers squared
            .take_while(|&n_squared| n_squared < upper) // Below upper limit
            .filter(|&n_squared| is_odd(n_squared))     // That are odd
            .fold(0, |acc, n_squared| acc + n_squared); // Sum them
    println!("functional style: {}", sum_of_squared_odd_numbers);

    println!("{}", (1..5).map(|x| x * 2).fold(0, |n, sum| n + sum))
}

#[allow(dead_code)]
#[allow(unused_variables)]
pub fn diverging() {
    fn foo() -> ! {
        panic!("This call never returns.");
    }

    fn some_fn() {
        println!("panic");
        ()
    }

    fn main() {
        let a: () = some_fn();
        println!("This function returns and you can see this line.")
    }
    main()
}

#[allow(unused_attributes)]
#[allow(unused_variables)]
#[allow(unreachable_code)]
#[allow(dead_code)]
pub fn diverging2() {
    #![feature(never_type)]
    fn main() {
        let x = panic!("This call never returns.");
        println!("You will never see this line!");
    }
//    main()
}

pub fn diverging3() {
    fn sum_odd_numbers(up_to: u32) -> u32 {
        let mut acc = 0;
        for i in 0..up_to {
            // Notice that the return type of this match expression must be u32
            // because of the type of the "addition" variable.
            let addition: u32 = match i % 2 == 1 {
                // The "i" variable is of type u32, which is perfectly fine.
                true => i,
                // On the other hand, the "continue" expression does not return
                // u32, but it is still fine, because it never returns and therefore
                // does not violate the type requirements of the match expression.
                false => continue,
            };
            acc += addition;
        }
        acc
    }
    println!("Sum of odd numbers up to 9 (excluding): {}", sum_odd_numbers(9));
    println!("Sum of odd numbers up to 9 (excluding): {}",
             (0..9).filter(|&n| n % 2 == 1)     // That are odd
                 .fold(0, |x, y| x + y));
}