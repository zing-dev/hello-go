pub mod variable;

fn main() {
    {
        use variable::bindings;
//        bindings::bindings();
//        bindings::mut_test();
//        bindings::scope();
        bindings::declare();
    }
}