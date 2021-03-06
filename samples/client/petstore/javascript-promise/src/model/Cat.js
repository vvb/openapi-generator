/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 *
 * OpenAPI Generator version: 4.3.1-SNAPSHOT
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/Animal', 'model/CatAllOf'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Animal'), require('./CatAllOf'));
  } else {
    // Browser globals (root is window)
    if (!root.OpenApiPetstore) {
      root.OpenApiPetstore = {};
    }
    root.OpenApiPetstore.Cat = factory(root.OpenApiPetstore.ApiClient, root.OpenApiPetstore.Animal, root.OpenApiPetstore.CatAllOf);
  }
}(this, function(ApiClient, Animal, CatAllOf) {
  'use strict';



  /**
   * The Cat model module.
   * @module model/Cat
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Cat</code>.
   * @alias module:model/Cat
   * @class
   * @extends module:model/Animal
   * @implements module:model/Animal
   * @implements module:model/CatAllOf
   * @param className {String} 
   */
  var exports = function(className) {
    var _this = this;

    Animal.call(_this, className);
  };

  /**
   * Constructs a <code>Cat</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Cat} obj Optional instance to populate.
   * @return {module:model/Cat} The populated <code>Cat</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      Animal.constructFromObject(data, obj);
      if (data.hasOwnProperty('declawed')) {
        obj['declawed'] = ApiClient.convertToType(data['declawed'], 'Boolean');
      }
    }
    return obj;
  }

  exports.prototype = Object.create(Animal.prototype);
  exports.prototype.constructor = exports;

  /**
   * @member {Boolean} declawed
   */
  exports.prototype['declawed'] = undefined;

  // Implement Animal interface:
  /**
   * @member {String} className
   */
exports.prototype['className'] = undefined;

  /**
   * @member {String} color
   * @default 'red'
   */
exports.prototype['color'] = 'red';

  // Implement CatAllOf interface:
  /**
   * @member {Boolean} declawed
   */
exports.prototype['declawed'] = undefined;



  return exports;
}));


