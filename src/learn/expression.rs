#[allow(path_statements)]
#[allow(unused_must_use)]
pub fn expression1() {
    // variable binding
    let x = 5;

    // expression;
    x;
    x + 1;
    15;
}

pub fn expression2() {
    let x = 5u32;

    let y = {
        let x_squared = x * x;
        let x_cube = x_squared * x;

        // This expression will be assigned to `y`
        x_cube + x_squared + x
    };

    #[allow(unused_must_use)]
    let z = {
        // The semicolon suppresses this expression and `()` is assigned to `z`
        2 * x;
    };

    let z2 = {
        2 * x
    };

    println!("x is {:?}", x);
    println!("y is {:?}", y);
    println!("z is {:#?}", z);
    println!("z is {:#?}", z2);
}
