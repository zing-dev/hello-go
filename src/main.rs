#![allow(unused)]

pub mod learn;


fn types() {
    use learn::types;

    types::literals();
    types::inference();
    types::alias()
}

fn conversion(){
    use learn::conversion;

    conversion::from_into();
    conversion::try_from_into();
    conversion::to_string()
}

fn expression(){
    use learn::expression;

    expression::expression2()
}


fn flow_control(){
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


fn main() {
    flow_control()
}