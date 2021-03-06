# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    OpenAPI spec version: 1.0.0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import petstore_api
from petstore_api.api.pet_api import PetApi  # noqa: E501
from petstore_api.rest import ApiException


class TestPetApi(unittest.TestCase):
    """PetApi unit test stubs"""

    def setUp(self):
        self.api = petstore_api.api.pet_api.PetApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_pet(self):
        """Test case for add_pet

        Add a new pet to the store  # noqa: E501
        """
        pass

    def test_delete_pet(self):
        """Test case for delete_pet

        Deletes a pet  # noqa: E501
        """
        pass

    def test_find_pets_by_status(self):
        """Test case for find_pets_by_status

        Finds Pets by status  # noqa: E501
        """
        pass

    def test_find_pets_by_tags(self):
        """Test case for find_pets_by_tags

        Finds Pets by tags  # noqa: E501
        """
        pass

    def test_get_pet_by_id(self):
        """Test case for get_pet_by_id

        Find pet by ID  # noqa: E501
        """
        pass

    def test_update_pet(self):
        """Test case for update_pet

        Update an existing pet  # noqa: E501
        """
        pass

    def test_update_pet_with_form(self):
        """Test case for update_pet_with_form

        Updates a pet in the store with form data  # noqa: E501
        """
        pass

    def test_upload_file(self):
        """Test case for upload_file

        uploads an image  # noqa: E501
        """
        pass

    def test_upload_file_with_required_file(self):
        """Test case for upload_file_with_required_file

        uploads an image (required)  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
