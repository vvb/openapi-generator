/*
 * Example
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 * Generated by: https://openapi-generator.tech
 */

use std::rc::Rc;
use std::borrow::Borrow;
use std::pin::Pin;
#[allow(unused_imports)]
use std::option::Option;

use hyper;
use hyper_util::client::legacy::connect::Connect;
use futures::Future;

use crate::models;
use super::{Error, configuration};
use super::request as __internal_request;

pub struct DefaultApiClient<C: Connect>
    where C: Clone + std::marker::Send + Sync + 'static {
    configuration: Rc<configuration::Configuration<C>>,
}

impl<C: Connect> DefaultApiClient<C>
    where C: Clone + std::marker::Send + Sync {
    pub fn new(configuration: Rc<configuration::Configuration<C>>) -> DefaultApiClient<C> {
        DefaultApiClient {
            configuration,
        }
    }
}

pub trait DefaultApi {
    fn get_fruit(&self, ) -> Pin<Box<dyn Future<Output = Result<models::Fruit, Error>>>>;
}

impl<C: Connect>DefaultApi for DefaultApiClient<C>
    where C: Clone + std::marker::Send + Sync {
    #[allow(unused_mut)]
    fn get_fruit(&self, ) -> Pin<Box<dyn Future<Output = Result<models::Fruit, Error>>>> {
        let mut req = __internal_request::Request::new(hyper::Method::GET, "/example".to_string())
        ;

        req.execute(self.configuration.borrow())
    }

}
