#![allow(ellipsis_inclusive_range_patterns)]
#![allow(dead_code)]

pub mod learn;

fn types() {
    use learn::types;

    types::literals();
    types::inference();
    types::alias()
}

fn conversion() {
    use learn::conversion;

    conversion::from_into();
    conversion::try_from_into();
    conversion::to_string()
}

fn expression() {
    use learn::expression;

    expression::expression2()
}


fn flow_control() {
    use learn::flow_control;

    flow_control::if_else();
    flow_control::loop_test();
    flow_control::loop_nested_label();
    flow_control::loop_return();
    flow_control::while_test();
    flow_control::for_range();
    flow_control::for_iter();
    flow_control::for_iter_mut();
}

fn match_test() {
    use learn::match_test;
    match_test::match_test();
    match_test::match_tuples();
    match_test::match_enums();
    match_test::match_pointers();
    match_test::match_structs();
    match_test::match_guards();
    match_test::match_bindings();
    match_test::if_let();
    match_test::while_let();
}

fn function() {
    use learn::function;

    function::function();
    function::methods();
    function::closures();
    function::capture();
    function::input_parameters();
    function::input_functions();
    function::anonymity();
    function::output_parameters();
//    function::iter_any();
}

fn function2() {
    use learn::function2;

    function2::hof();
    function2::diverging();
    function2::diverging2();
    function2::diverging3();
}

fn mod_test() {
    use learn::mod_test;

    mod_test::mod_test();
    mod_test::struct_test();
    mod_test::use_test();
    mod_test::super_test();
}

fn attribute() {
    use learn::attribute;

    attribute::dead_code();
    attribute::cfg();
}

fn generics() {
    use learn::generics;
    generics::gen_fn();
    generics::implement();
    generics::gen_trait();
    generics::bounds();
    generics::case_empty();
    generics::multi_bounds();
    generics::where_test();
    generics::new_types();
}

fn main() {
    generics()
}