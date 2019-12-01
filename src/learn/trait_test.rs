pub fn derive() {
    // `Centimeters`, a tuple struct that can be compared
    #[derive(PartialEq, PartialOrd)]
    struct Centimeters(f64);

    // `Inches`, a tuple struct that can be printed
    #[derive(Debug)]
    struct Inches(i32);

    impl Inches {
        fn to_centimeters(&self) -> Centimeters {
            let &Inches(inches) = self;

            Centimeters(inches as f64 * 2.54)
        }
    }

    // `Seconds`, a tuple struct with no additional attributes
    struct Seconds(i32);

    let _one_second = Seconds(1);

    // Error: `Seconds` can't be printed; it doesn't implement the `Debug` trait
    //println!("One second looks like: {:?}", _one_second);
    // TODO ^ Try uncommenting this line

    // Error: `Seconds` can't be compared; it doesn't implement the `PartialEq` trait
    //let _this_is_true = (_one_second == _one_second);
    // TODO ^ Try uncommenting this line

    let foot = Inches(12);

    println!("One foot equals {:?}", foot);

    let meter = Centimeters(100.0);

    let cmp =
        if foot.to_centimeters() < meter {
            "smaller"
        } else {
            "bigger"
        };

    println!("One foot is {} than one meter.", cmp);
}

pub fn dyn_test() {
    struct Sheep {}
    struct Cow {}

    trait Animal {
        // Instance method signature
        fn noise(&self) -> &'static str;
    }

    // Implement the `Animal` trait for `Sheep`.
    impl Animal for Sheep {
        fn noise(&self) -> &'static str {
            "baaaaah!"
        }
    }

    // Implement the `Animal` trait for `Cow`.
    impl Animal for Cow {
        fn noise(&self) -> &'static str {
            "moooooo!"
        }
    }

    // Returns some struct that implements Animal, but we don't know which one at compile time.
    fn random_animal(random_number: f64) -> Box<dyn Animal> {
        if random_number < 0.5 {
            Box::new(Sheep {})
        } else {
            Box::new(Cow {})
        }
    }

    let random_number = 0.234;
    let animal = random_animal(random_number);
    println!("You've randomly chosen an animal, and it says {}", animal.noise());
}

pub fn overload() {
    use std::ops;

    struct Foo;
    struct Bar;

    #[derive(Debug)]
    struct FooBar;

    #[derive(Debug)]
    struct BarFoo;

    // The `std::ops::Add` trait is used to specify the functionality of `+`.
// Here, we make `Add<Bar>` - the trait for addition with a RHS of type `Bar`.
// The following block implements the operation: Foo + Bar = FooBar
    impl ops::Add<Bar> for Foo {
        type Output = FooBar;

        fn add(self, _rhs: Bar) -> FooBar {
            println!("> Foo.add(Bar) was called");

            FooBar
        }
    }

    // By reversing the types, we end up implementing non-commutative addition.
// Here, we make `Add<Foo>` - the trait for addition with a RHS of type `Foo`.
// This block implements the operation: Bar + Foo = BarFoo
    impl ops::Add<Foo> for Bar {
        type Output = BarFoo;

        fn add(self, _rhs: Foo) -> BarFoo {
            println!("> Bar.add(Foo) was called");

            BarFoo
        }
    }

    println!("Foo + Bar = {:?}", Foo + Bar);
    println!("Bar + Foo = {:?}", Bar + Foo);
}

pub fn drop_test() {
    struct DropTest {
        name: &'static str,
    }

    // This trivial implementation of `drop` adds a print to console.
    impl Drop for DropTest {
        fn drop(&mut self) {
            println!("> Dropping {}", self.name);
        }
    }

    let _a = DropTest { name: "a" };

    // block A
    {
        let _b = DropTest { name: "b" };

        // block B
        {
            let _c = DropTest { name: "c" };
            let _d = DropTest { name: "d" };

            println!("Exiting block B");
        }
        println!("Just exited block B");

        println!("Exiting block A");
    }
    println!("Just exited block A");

    // Variable can be manually dropped using the `drop` function
    drop(_a);
    // TODO ^ Try commenting this line
    println!("end of the main function");

    // `_a` *won't* be `drop`ed again here, because it already has been
    // (manually) `drop`ed
}

pub fn iter() {
    struct Fibonacci {
        curr: u32,
        next: u32,
    }

    // Implement `Iterator` for `Fibonacci`.
// The `Iterator` trait only requires a method to be defined for the `next` element.
    impl Iterator for Fibonacci {
        type Item = u32;

        // Here, we define the sequence using `.curr` and `.next`.
        // The return type is `Option<T>`:
        //     * When the `Iterator` is finished, `None` is returned.
        //     * Otherwise, the next value is wrapped in `Some` and returned.
        fn next(&mut self) -> Option<u32> {
            let new_next = self.curr + self.next;

            self.curr = self.next;
            self.next = new_next;

            // Since there's no endpoint to a Fibonacci sequence, the `Iterator`
            // will never return `None`, and `Some` is always returned.
            Some(self.curr)
        }
    }

    // Returns a Fibonacci sequence generator
    fn fibonacci() -> Fibonacci {
        Fibonacci { curr: 0, next: 1 }
    }

    // `0..3` is an `Iterator` that generates: 0, 1, and 2.
    let mut sequence = 0..3;

    println!("Four consecutive `next` calls on 0..3");
    println!("> {:?}", sequence.next());
    println!("> {:?}", sequence.next());
    println!("> {:?}", sequence.next());
    println!("> {:?}", sequence.next());

    // `for` works through an `Iterator` until it returns `None`.
    // Each `Some` value is unwrapped and bound to a variable (here, `i`).
    println!("Iterate through 0..3 using `for`");
    for i in 0..3 {
        println!("> {}", i);
    }

    // The `take(n)` method reduces an `Iterator` to its first `n` terms.
    println!("The first four terms of the Fibonacci sequence are: ");
    for i in fibonacci().take(4) {
        println!("> {}", i);
    }

    // The `skip(n)` method shortens an `Iterator` by dropping its first `n` terms.
    println!("The next four terms of the Fibonacci sequence are: ");
    for i in fibonacci().skip(4).take(4) {
        println!("> {}", i);
    }

    let array = [1u32, 3, 3, 7];

    // The `iter` method produces an `Iterator` over an array/slice.
    println!("Iterate the following array {:?}", &array);
    for i in array.iter() {
        println!("> {}", i);
    }
}

pub fn impl_trait() {
    use std::iter;
    use std::vec::IntoIter;

    // This function combines two `Vec<i32>` and returns an iterator over it.
// Look how complicated its return type is!
    fn combine_vecs_explicit_return_type<'a>(
        v: Vec<i32>,
        u: Vec<i32>,
    ) -> iter::Cycle<iter::Chain<IntoIter<i32>, IntoIter<i32>>> {
        v.into_iter().chain(u.into_iter()).cycle()
    }

    // This is the exact same function, but its return type uses `impl Trait`.
// Look how much simpler it is!
    fn combine_vecs<'a>(
        v: Vec<i32>,
        u: Vec<i32>,
    ) -> impl Iterator<Item=i32> {
        v.into_iter().chain(u.into_iter()).cycle()
    }

    fn double_positives<'a>(numbers: &'a Vec<i32>) -> impl Iterator<Item=i32> + 'a {
        numbers
            .iter()
            .filter(|x| x > &&0)
            .map(|x| x * 2)
    }

    // Returns a function that adds `y` to its input
    fn make_adder_function(y: i32) -> impl Fn(i32) -> i32 {
        let closure = move |x: i32| { x + y };
        closure
    }

    let plus_one = make_adder_function(1);
    assert_eq!(plus_one(2), 3);
}

pub fn clone() {
    // A unit struct without resources
    #[derive(Debug, Clone, Copy)]
    struct Nil;

    // A tuple struct with resources that implements the `Clone` trait
    #[derive(Clone, Debug)]
    struct Pair(Box<i32>, Box<i32>);

    // Instantiate `Nil`
    let nil = Nil;
    // Copy `Nil`, there are no resources to move
    let copied_nil = nil;

    // Both `Nil`s can be used independently
    println!("original: {:?}", nil);
    println!("copy: {:?}", copied_nil);

    // Instantiate `Pair`
    let pair = Pair(Box::new(1), Box::new(2));
    println!("original: {:?}", pair);

    // Copy `pair` into `moved_pair`, moves resources
    let moved_pair = pair;
    println!("copy: {:?}", moved_pair);

    // Error! `pair` has lost its resources
    //println!("original: {:?}", pair);
    // TODO ^ Try uncommenting this line

    // Clone `moved_pair` into `cloned_pair` (resources are included)
    let cloned_pair = moved_pair.clone();
    // Drop the original pair using std::mem::drop
    drop(moved_pair);

    // Error! `moved_pair` has been dropped
    //println!("copy: {:?}", moved_pair);
    // TODO ^ Try uncommenting this line

    // The result from .clone() can still be used!
    println!("clone: {:?}", cloned_pair);
}

pub fn supertraits() {
    trait Person {
        fn name(&self) -> String;
    }

    // Student is a supertrait of Person.
// Implementing Student requires you to also impl Person.
    trait Student: Person {
        fn university(&self) -> String;
    }

    trait Programmer {
        fn fav_language(&self) -> String;
    }

    // CompSciStudent (computer science student) is a supertrait of both Programmer
// and Student. Implementing CompSciStudent requires you to impl both subtraits.
    trait CompSciStudent: Programmer + Student {
        fn git_username(&self) -> String;
    }

    fn comp_sci_student_greeting(student: &dyn CompSciStudent) -> String {
        format!(
            "My name is {} and I attend {}. My Git username is {}",
            student.name(),
            student.university(),
            student.git_username()
        )
    }

    fn main() {}
}

pub fn disambiguating() {
    trait UsernameWidget {
        // Get the selected username out of this widget
        fn get(&self) -> String;
    }

    trait AgeWidget {
        // Get the selected age out of this widget
        fn get(&self) -> u8;
    }

    // A form with both a UsernameWidget and an AgeWidget
    struct Form {
        username: String,
        age: u8,
    }

    impl UsernameWidget for Form {
        fn get(&self) -> String {
            self.username.clone()
        }
    }

    impl AgeWidget for Form {
        fn get(&self) -> u8 {
            self.age
        }
    }

    let form = Form {
        username: "rustacean".to_owned(),
        age: 28,
    };

    // If you uncomment this line, you'll get an error saying
    // "multiple `get` found". Because, after all, there are multiple methods
    // named `get`.
    // println!("{}", form.get());

    let username = <Form as UsernameWidget>::get(&form);
    assert_eq!("rustacean".to_owned(), username);
    let age = <Form as AgeWidget>::get(&form);
    assert_eq!(28, age);
}