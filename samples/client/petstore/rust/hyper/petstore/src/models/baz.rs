/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

/// Baz : Test handling of empty variants
/// Test handling of empty variants
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Baz {
    #[serde(rename = "A")]
    A,
    #[serde(rename = "B")]
    B,
    #[serde(rename = "")]
    Empty,

}

impl ToString for Baz {
    fn to_string(&self) -> String {
        match self {
            Self::A => String::from("A"),
            Self::B => String::from("B"),
            Self::Empty => String::from(""),
        }
    }
}

impl Default for Baz {
    fn default() -> Baz {
        Self::A
    }
}

