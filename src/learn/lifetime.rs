// Lifetimes are annotated below with lines denoting the creation
// and destruction of each variable.
// `i` has the longest lifetime because its scope entirely encloses
// both `borrow1` and `borrow2`. The duration of `borrow1` compared
// to `borrow2` is irrelevant since they are disjoint.
pub fn lifetimes() {
    let i = 3; // Lifetime for `i` starts. ────────────────┐
    //                                                     │
    { //                                                   │
        let borrow1 = &i; // `borrow1` lifetime starts. ──┐│
        //                                                ││
        println!("borrow1: {}", borrow1); //              ││
    } // `borrow1 ends. ──────────────────────────────────┘│
    //                                                     │
    //                                                     │
    { //                                                   │
        let borrow2 = &i; // `borrow2` lifetime starts. ──┐│
        //                                                ││
        println!("borrow2: {}", borrow2); //              ││
    } // `borrow2` ends. ─────────────────────────────────┘│
    //                                                     │
}   // Lifetime ends. ─────────────────────────────────────┘


pub fn explicit() {
    // `print_refs` takes two references to `i32` which have different
// lifetimes `'a` and `'b`. These two lifetimes must both be at
// least as long as the function `print_refs`.
    fn print_refs(x: &i32, y: &i32) {
        println!("x is {} and y is {}", x, y);
    }

    // A function which takes no arguments, but has a lifetime parameter `'a`.
    fn failed_borrow<'a>() {
        let _x = 12;

        // ERROR: `_x` does not live long enough
//        let y: &'a i32 = &_x;
        // Attempting to use the lifetime `'a` as an explicit type annotation
        // inside the function will fail because the lifetime of `&_x` is shorter
        // than that of `y`. A short lifetime cannot be coerced into a longer one.
    }

    // Create variables to be borrowed below.
    let (four, nine) = (4, 9);

    // Borrows (`&`) of both variables are passed into the function.
    print_refs(&four, &nine);
    // Any input which is borrowed must outlive the borrower.
    // In other words, the lifetime of `four` and `nine` must
    // be longer than that of `print_refs`.

    failed_borrow();
    // `failed_borrow` contains no references to force `'a` to be
    // longer than the lifetime of the function, but `'a` is longer.
    // Because the lifetime is never constrained, it defaults to `'static`.
}

pub fn function() {
    // One input reference with lifetime `'a` which must live
// at least as long as the function.
    fn print_one(x: &i32) {
        println!("`print_one`: x is {}", x);
    }

    // Mutable references are possible with lifetimes as well.
    fn add_one(x: &mut i32) {
        *x += 1;
    }

    // Multiple elements with different lifetimes. In this case, it
// would be fine for both to have the same lifetime `'a`, but
// in more complex cases, different lifetimes may be required.
    fn print_multi(x: &i32, y: &i32) {
        println!("`print_multi`: x is {}, y is {}", x, y);
    }

    // Returning references that have been passed in is acceptable.
// However, the correct lifetime must be returned.
    fn pass_x<'a, 'b>(x: &'a i32, _: &'b i32) -> &'a i32 { x }

//fn invalid_output<'a>() -> &'a String { &String::from("foo") }
// The above is invalid: `'a` must live longer than the function.
// Here, `&String::from("foo")` would create a `String`, followed by a
// reference. Then the data is dropped upon exiting the scope, leaving
// a reference to invalid data to be returned.

    let x = 7;
    let y = 9;

    print_one(&x);
    print_multi(&x, &y);

    let z = pass_x(&x, &y);
    print_one(z);

    let mut t = 3;
    add_one(&mut t);
    print_one(&t);
}

pub fn methods() {
    struct Owner(i32);

    impl Owner {
        // Annotate lifetimes as in a standalone function.
        fn add_one(&mut self) { self.0 += 1; }
        fn print(&self) {
            println!("`print`: {}", self.0);
        }
    }

    let mut owner = Owner(18);

    owner.add_one();
    owner.print();
}

pub fn struct_test() {
    // A type `Borrowed` which houses a reference to an
// `i32`. The reference to `i32` must outlive `Borrowed`.
    #[derive(Debug)]
    struct Borrowed<'a>(&'a i32);

    // Similarly, both references here must outlive this structure.
    #[derive(Debug)]
    struct NamedBorrowed<'a> {
        x: &'a i32,
        y: &'a i32,
    }

    // An enum which is either an `i32` or a reference to one.
    #[derive(Debug)]
    enum Either<'a> {
        Num(i32),
        Ref(&'a i32),
    }

    let x = 18;
    let y = 15;

    let single = Borrowed(&x);
    let double = NamedBorrowed { x: &x, y: &y };
    let reference = Either::Ref(&x);
    let number = Either::Num(y);

    println!("x is borrowed in {:?}", single);
    println!("x and y are borrowed in {:?}", double);
    println!("x is borrowed in {:?}", reference);
    println!("y is *not* borrowed in {:?}", number);
}

pub fn trait_test() {
    // A struct with annotation of lifetimes.
    #[derive(Debug)]
    struct Borrowed<'a> {
        x: &'a i32,
    }

    #[derive(Debug)]
    struct Borrowed2 {
        x: i32,
    }

    // Annotate lifetimes to impl.
    impl<'a> Default for Borrowed<'a> {
        fn default() -> Self {
            Self {
                x: &10,
            }
        }
    }
    // Annotate lifetimes to impl.
    impl Default for Borrowed2 {
        fn default() -> Self {
            Self {
                x: 11,
            }
        }
    }

    let b: Borrowed = Default::default();
    let b2: Borrowed2 = Default::default();
    println!("b is {:?}", b);
    println!("b is {:?}", b2);
}

pub fn bounds() {
    use std::fmt::Debug; // Trait to bound with.

    #[derive(Debug)]
    struct Ref<'a, T: 'a>(&'a T);
// `Ref` contains a reference to a generic type `T` that has
// an unknown lifetime `'a`. `T` is bounded such that any
// *references* in `T` must outlive `'a`. Additionally, the lifetime
// of `Ref` may not exceed `'a`.

    // A generic function which prints using the `Debug` trait.
    fn print<T>(t: T) where
        T: Debug {
        println!("`print`: t is {:?}", t);
    }

    // Here a reference to `T` is taken where `T` implements
// `Debug` and all *references* in `T` outlive `'a`. In
// addition, `'a` must outlive the function.
    fn print_ref<'a, T>(t: &'a T) where
        T: Debug + 'a {
        println!("`print_ref`: t is {:?}", t);
    }

    let x = 7;
    let ref_x = Ref(&x);

    print_ref(&ref_x);
    print(ref_x);
}

pub fn coercion() {
    // Here, Rust infers a lifetime that is as short as possible.
// The two references are then coerced to that lifetime.
    fn multiply<'a>(first: &'a i32, second: &'a i32) -> i32 {
        first * second
    }

    // `<'a: 'b, 'b>` reads as lifetime `'a` is at least as long as `'b`.
// Here, we take in an `&'a i32` and return a `&'b i32` as a result of coercion.
    fn choose_first<'a: 'b, 'b>(first: &'a i32, _: &'b i32) -> &'b i32 {
        first
    }

    let first = 2; // Longer lifetime

    {
        let second = 3; // Shorter lifetime

        println!("The product is {}", multiply(&first, &second));
        println!("The product is {}", multiply(&3, &4));
        println!("{} is the first", choose_first(&first, &second));
    };
}

pub fn static_test() {
    // Make a constant with `'static` lifetime.
    static NUM: i32 = 18;

    // Returns a reference to `NUM` where its `'static`
// lifetime is coerced to that of the input argument.
    fn coerce_static(_: &i32) -> &i32 {
        &NUM
    }

    {
        // Make a `string` literal and print it:
        let static_string = "I'm in read-only memory";
        println!("static_string: {}", static_string);

        // When `static_string` goes out of scope, the reference
        // can no longer be used, but the data remains in the binary.
    }

    {
        // Make an integer to use for `coerce_static`:
        let lifetime_num = 9;

        // Coerce `NUM` to lifetime of `lifetime_num`:
        let coerced_static = coerce_static(&lifetime_num);

        println!("coerced_static: {}", coerced_static);
    }

    println!("NUM: {} stays accessible!", NUM);
}