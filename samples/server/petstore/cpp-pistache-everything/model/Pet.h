/**
* OpenAPI Petstore
* This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
*
* The version of the OpenAPI document: 1.0.0
* 
*
* NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
* https://openapi-generator.tech
* Do not edit the class manually.
*/
/*
 * Pet.h
 *
 * A pet for sale in the pet store
 */

#ifndef Pet_H_
#define Pet_H_


#include "Tag.h"
#include <nlohmann/json.hpp>
#include "Pet_bestFriends.h"
#include <string>
#include "Category.h"
#include <vector>
#include <nlohmann/json.hpp>

namespace org::openapitools::server::model
{

/// <summary>
/// A pet for sale in the pet store
/// </summary>
class  Pet
{
public:
    Pet();
    virtual ~Pet() = default;


    /// <summary>
    /// Validate the current data in the model. Throws a ValidationException on failure.
    /// </summary>
    void validate() const;

    /// <summary>
    /// Validate the current data in the model. Returns false on error and writes an error
    /// message into the given stringstream.
    /// </summary>
    bool validate(std::stringstream& msg) const;

    /// <summary>
    /// Helper overload for validate. Used when one model stores another model and calls it's validate.
    /// Not meant to be called outside that case.
    /// </summary>
    bool validate(std::stringstream& msg, const std::string& pathPrefix) const;

    bool operator==(const Pet& rhs) const;
    bool operator!=(const Pet& rhs) const;

    /////////////////////////////////////////////
    /// Pet members

    /// <summary>
    /// 
    /// </summary>
    int64_t getId() const;
    void setId(int64_t const value);
    bool idIsSet() const;
    void unsetId();
    /// <summary>
    /// 
    /// </summary>
    org::openapitools::server::model::Category getCategory() const;
    void setCategory(org::openapitools::server::model::Category const& value);
    bool categoryIsSet() const;
    void unsetCategory();
    /// <summary>
    /// 
    /// </summary>
    std::string getName() const;
    void setName(std::string const& value);
    /// <summary>
    /// 
    /// </summary>
    std::vector<std::string> getPhotoUrls() const;
    void setPhotoUrls(std::vector<std::string> const& value);
    /// <summary>
    /// 
    /// </summary>
    std::vector<org::openapitools::server::model::Tag> getTags() const;
    void setTags(std::vector<org::openapitools::server::model::Tag> const& value);
    bool tagsIsSet() const;
    void unsetTags();
    /// <summary>
    /// pet status in the store
    /// </summary>
    std::string getStatus() const;
    void setStatus(std::string const& value);
    bool statusIsSet() const;
    void unsetStatus();
    /// <summary>
    /// last veterinarian visit advice
    /// </summary>
    nlohmann::json getVeterinarianVisit() const;
    void setVeterinarianVisit(nlohmann::json const& value);
    bool veterinarianVisitIsSet() const;
    void unsetVeterinarianVisit();
    /// <summary>
    /// to help you installing your pet at home
    /// </summary>
    std::vector<nlohmann::json> getGoodies() const;
    void setGoodies(std::vector<nlohmann::json> const& value);
    bool goodiesIsSet() const;
    void unsetGoodies();
    /// <summary>
    /// 
    /// </summary>
    org::openapitools::server::model::Pet_bestFriends getBestFriends() const;
    void setBestFriends(org::openapitools::server::model::Pet_bestFriends const& value);
    bool bestFriendsIsSet() const;
    void unsetBestFriends();

    friend  void to_json(nlohmann::json& j, const Pet& o);
    friend  void from_json(const nlohmann::json& j, Pet& o);
protected:
    int64_t m_Id;
    bool m_IdIsSet;
    org::openapitools::server::model::Category m_Category;
    bool m_CategoryIsSet;
    std::string m_Name;

    std::vector<std::string> m_PhotoUrls;

    std::vector<org::openapitools::server::model::Tag> m_Tags;
    bool m_TagsIsSet;
    std::string m_Status;
    bool m_StatusIsSet;
    nlohmann::json m_VeterinarianVisit;
    bool m_VeterinarianVisitIsSet;
    std::vector<nlohmann::json> m_Goodies;
    bool m_GoodiesIsSet;
    org::openapitools::server::model::Pet_bestFriends m_BestFriends;
    bool m_BestFriendsIsSet;
    
};

} // namespace org::openapitools::server::model

#endif /* Pet_H_ */
