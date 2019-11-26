//#[attribute = "value"]
//#[attribute(key = "value")]
//#[attribute(value)]

//// This crate is a library
//#![crate_type = "lib"]
//// The library is named "rary"
//#![crate_name = "rary"]
//pub fn public_function() {
//    println!("called rary's `public_function()`");
//}
//
//fn private_function() {
//    println!("called rary's `private_function()`");
//}
//
//pub fn indirect_access() {
//    print!("called rary's `indirect_access()`, that\n> ");
//
//    private_function();
//}

pub fn dead_code() {
    fn used_function() {}

    // `#[allow(dead_code)]` is an attribute that disables the `dead_code` lint
    #[allow(dead_code)]
    fn unused_function() {}

    #[allow(dead_code)]
    fn unused_function2() {
        println!("unused_function2")
    }

    #[allow(dead_code)]
    fn noisy_unused_function() {}
// FIXME ^ Add an attribute to suppress the warning

    used_function();
    unused_function();
    unused_function2();
}

// This function only gets compiled if the target OS is linux
#[cfg(target_os = "linux")]
fn are_you_on_linux() {
    println!("You are running linux!");
}

// And this function only gets compiled if the target OS is *not* linux
#[cfg(not(target_os = "linux"))]
fn are_you_on_linux() {
    println!("You are *not* running linux!");
}

#[cfg(target_os = "windows")]
fn are_you_on_windows() {
    println!("You are running windows!");
}

pub fn cfg() {
    are_you_on_linux();
    println!("Are you sure?");
    if cfg!(target_os = "linux") {
        println!("Yes. It's definitely linux!");
    } else {
        println!("Yes. It's definitely *not* linux!");
    }

    are_you_on_windows()
}
