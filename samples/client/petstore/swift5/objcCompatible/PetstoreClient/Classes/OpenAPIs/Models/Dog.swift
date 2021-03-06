//
// Dog.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation

@objc public class Dog: NSObject, Codable {

    public var _className: String
    public var color: String? = "red"
    public var breed: String?

    public init(_className: String, color: String?, breed: String?) {
        self._className = _className
        self.color = color
        self.breed = breed
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case _className = "className"
        case color
        case breed
    }

}
