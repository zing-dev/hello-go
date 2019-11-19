#![allow(unused_variables)]
#![allow(dead_code)]

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

fn main() {
    conversion()
}