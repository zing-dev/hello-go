#[allow(unused_imports)]
#[allow(dead_code)]
use std::io::{BufWriter, stdin, stdout};

use ferris_says::say;

fn hello() {
    let stdout = stdout();
    let out = b"hello rust!";
    let width = 24;

    let mut writer = BufWriter::new(stdout.lock());
    say(out, width, &mut writer).unwrap();
}

fn print() {
    // `print!` is like `println!` but it doesn't add a newline at the end
    print!("January has ");

    // `{}` are placeholders for arguments that will be stringified
    println!("{} days", 31);
    // The `i` suffix indicates the compiler that this literal has type: signed
    // pointer size integer, see next chapter for more details

    // The positional arguments can be reused along the template
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

    // Named arguments can also be used
    println!("{subject} {verb} {predicate}",
             predicate = "over the lazy dog",
             subject = "the quick brown fox",
             verb = "jumps");

    // Special formatting can be specified in the placeholder after a `:`
    println!("{} of {:b} people know binary, the other half don't", 1, 2);

    // Error! You are missing an argument
    println!("My name is {0}, {1} {0}", "Bond", 's');
    // FIXME ^ Add the missing argument: "James"
}

fn literals() {
    // Integer addition
    println!("1 + 2 = {}", 1u8 + 2);

    // Integer subtraction
    println!("1 - 2 = {}", 1i8 - 2);
    // TODO ^ Try changing `1i` to `1u` to see why the type is important

    // Short-circuiting boolean logic
    println!("true AND false is {}", true && false);
    println!("true OR false is {}", true || false);
    println!("NOT true is {}", !true);

    // Bitwise operations
    println!("0011 AND 0101 is {:04b}", 0b0011u8 & 0b0101);
    println!("0011 OR 0101 is {:04b}", 0b0011u8 | 0b0101);
    println!("0011 XOR 0101 is {:04b}", 0b0011u8 ^ 0b0101);
    println!("1 << 5 is {}", 1u8 << 5);
    println!("0x80 >> 2 is 0x{:x}", 0x80u8 >> 2);

    // Use underscores to improve readability!
    println!("One million is written as {}", 1_000_000u32);
}

fn variables_mut() {
    let _immutable_variable = 1i8;
    let mut mutable_variable = 1i8;

    println!("Before mutation: {}", mutable_variable);

    // Ok
    mutable_variable += 1;

    println!("After mutation: {}", mutable_variable);

    // Error!
//    _immutable_variable += 1;
    // FIXME ^ Comment out this line
}

fn variables_scope(){
    // This variable lives in the main function
    let long_lived_variable = 1i32;
    println!("long_lived_variable {}",long_lived_variable);
    let mut long_lived_variable = 10i32;

    // This is a block, and has a smaller scope than the main function
    {
        long_lived_variable = 2i32;
        // This variable only exists in this block
        let short_lived_variable = 2i32;

        println!("inner short: {}", short_lived_variable);

        // This variable *shadows* the outer one
        let long_lived_variable = 5_f32;

        println!("inner long: {}", long_lived_variable);
    }
    // End of the block

    // Error! `short_lived_variable` doesn't exist in this scope
//    println!("outer short: {}", short_lived_variable);
    // FIXME ^ Comment out this line

    println!("outer long: {}", long_lived_variable);
}
fn main() {
//    hello()
//    print()
//    literals()
//    variables_mut()
    variables_scope()
}
