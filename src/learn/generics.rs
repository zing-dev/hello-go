use std::fmt::{Debug, Error, Formatter};

pub fn gen_fn() {
    struct A;          // Concrete type `A`.
    struct S(A);       // Concrete type `S`.
    struct SGen<T>(T); // Generic type `SGen`.

    impl Debug for A {
        fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), Error> {
            write!(f, "hello A")
        }
    }
// The following functions all take ownership of the variable passed into
// them and immediately go out of scope, freeing the variable.

    // Define a function `reg_fn` that takes an argument `_s` of type `S`.
// This has no `<T>` so this is not a generic function.
    fn reg_fn(_s: S) {}

    // Define a function `gen_spec_t` that takes an argument `_s` of type `SGen<T>`.
// It has been explicitly given the type parameter `A`, but because `A` has not
// been specified as a generic type parameter for `gen_spec_t`, it is not generic.
    fn gen_spec_t(_s: SGen<A>) {}

    // Define a function `gen_spec_i32` that takes an argument `_s` of type `SGen<i32>`.
// It has been explicitly given the type parameter `i32`, which is a specific type.
// Because `i32` is not a generic type, this function is also not generic.
    fn gen_spec_i32(_s: SGen<i32>) {}

    // Define a function `generic` that takes an argument `_s` of type `SGen<T>`.
// Because `SGen<T>` is preceded by `<T>`, this function is generic over `T`.
    fn generic<T>(_s: SGen<T>) {}

    // Using the non-generic functions
    reg_fn(S(A));          // Concrete type.
    gen_spec_t(SGen(A));   // Implicitly specified type parameter `A`.
    gen_spec_i32(SGen(6)); // Implicitly specified type parameter `i32`.

    // Explicitly specified type parameter `char` to `generic()`.
    generic::<char>(SGen('a'));

    // Implicitly specified type parameter `char` to `generic()`.
    generic(SGen('c'));

    let a = A {};
    println!("{:?}", a);
}

pub fn implement() {
    struct Val {
        val: f64,
    }

    struct GenVal<T> {
        gen_val: T,
    }

    struct Rust<T> {
        name: T,
    }

    impl<T> Rust<T> {
        fn value(&self) -> &T {
            &self.name
        }
    }
    // impl of Val
    impl Val {
        fn value(&self) -> &f64 {
            &self.val
        }
    }

    // impl of GenVal for a generic type `T`
    impl<T> GenVal<T> {
        fn value(&self) -> &T {
            &self.gen_val
        }
    }

    let x = Val { val: 3.0 };
    let y = GenVal { gen_val: 3i32 };
    let r = Rust { name: "RUST" };

    println!("{}, {}, {}", x.value(), y.value(), r.value().to_lowercase());
}


pub fn gen_trait() {
    // Non-copyable types.
    struct Empty;
    struct Null;

    // A trait generic over `T`.
    trait DoubleDrop<T> {
        // Define a method on the caller type which takes an
        // additional single parameter `T` and does nothing with it.
        fn double_drop(self, _: T);
    }

    // Implement `DoubleDrop<T>` for any generic parameter `T` and
// caller `U`.
    impl<T, U> DoubleDrop<T> for U {
        // This method takes ownership of both passed arguments,
        // deallocating both.
        fn double_drop(self, _: T) {}
    }

    let empty = Empty;
    let null = Null;

    // Deallocate `empty` and `null`.
    empty.double_drop(null);

    //empty;
    //null;
    // ^ TODO: Try uncommenting these lines.
}

pub fn bounds() {
    // A trait which implements the print marker: `{:?}`.
    trait HasArea {
        fn area(&self) -> f64;
    }

    impl HasArea for Rectangle {
        fn area(&self) -> f64 { self.length * self.height }
    }

    #[derive(Debug)]
    struct Rectangle { length: f64, height: f64 }

    #[derive(Debug)]
    #[allow(dead_code)]
    struct Triangle { length: f64, height: f64 }

    // The generic `T` must implement `Debug`. Regardless
// of the type, this will work properly.
    fn print_debug<T: Debug>(t: &T) {
        println!("{:?}", t);
    }

    // `T` must implement `HasArea`. Any type which meets
// the bound can access `HasArea`'s function `area`.
    fn area<T: HasArea>(t: &T) -> f64 { t.area() }
    fn area2<T>(t: &T) -> f64 where T: HasArea { t.area() }

    let rectangle = Rectangle { length: 3.0, height: 4.0 };
    let _triangle = Triangle { length: 3.0, height: 4.0 };

    print_debug(&rectangle);
    println!("Area: {}", area(&rectangle));
    println!("Area: {}", area2(&rectangle));

    print_debug(&_triangle);
//    println!("Area: {}", area(&_triangle));
    // ^ TODO: Try uncommenting these.
    // | Error: Does not implement either `Debug` or `HasArea`.
}

pub fn case_empty() {
    struct Cardinal;
    struct BlueJay;
    struct Turkey;

    trait Red {}
    trait Blue {}

    impl Red for Cardinal {}
    impl Blue for BlueJay {}
    impl Blue for Turkey {}

    // These functions are only valid for types which implement these
// traits. The fact that the traits are empty is irrelevant.
    fn red<T: Red>(_: &T) -> &'static str { "red" }
    fn blue<T: Blue>(_: &T) -> &'static str { "blue" }

    let cardinal = Cardinal;
    let blue_jay = BlueJay;
    let _turkey = Turkey;

    // `red()` won't work on a blue jay nor vice versa
    // because of the bounds.
    println!("A cardinal is {}", red(&cardinal));
    println!("A blue jay is {}", blue(&blue_jay));
    println!("A turkey is {}", blue(&_turkey));
    // ^ TODO: Try uncommenting this line.
}

#[allow(unused_imports)]
pub fn multi_bounds() {
    use std::fmt::{Debug, Display};

    fn compare_prints<T: Debug + Display>(t: &T) {
        println!("Debug: `{:?}`", t);
        println!("Display: `{}`", t);
    }

    fn compare_types<T: Debug, U: Debug>(t: &T, u: &U) {
        println!("t: `{:?}`", t);
        println!("u: `{:?}`", u);
    }

    let string = "words";
    let array = [1, 2, 3];
    let vec = vec![1, 2, 3];

    compare_prints(&string);
    //compare_prints(&array);
    // TODO ^ Try uncommenting this.

    compare_types(&array, &vec);
}

pub fn where_test() {
    trait PrintInOption {
        fn print_in_option(self);
    }

    // Because we would otherwise have to express this as `T: Debug` or
// use another method of indirect approach, this requires a `where` clause:
    impl<T> PrintInOption for T where
        Option<T>: Debug {
        // We want `Option<T>: Debug` as our bound because that is what's
        // being printed. Doing otherwise would be using the wrong bound.
        fn print_in_option(self) {
            println!("{:?}", Some(self));
        }
    }

    let vec = vec![1, 20, 3];
    vec.print_in_option();
}

pub fn new_types() {
    struct Years(i64);

    struct Days(i64);

    impl Years {
        pub fn to_days(&self) -> Days {
            Days(self.0 * 365)
        }
    }


    impl Days {
        /// truncates partial years
        pub fn to_years(&self) -> Years {
            Years(self.0 / 365)
        }
    }

    fn old_enough(age: &Years) -> bool {
        age.0 >= 18
    }

    let age = Years(5);
    let age_days = age.to_days();
    println!("Old enough {}", old_enough(&age));
    println!("Old enough {}", old_enough(&age_days.to_years()));
    // println!("Old enough {}", old_enough(&age_days));
}