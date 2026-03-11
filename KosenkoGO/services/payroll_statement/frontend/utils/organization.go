package utils

import (
	"fmt"

	organizationDto "github.com/joyzem/documents/services/organization/dto"
	"github.com/joyzem/documents/services/base"
	"github.com/levigross/grequests"
)

func GetOrganizations() (*organizationDto.GetOrganizationsResponse, error) {
	organizationsUrl := fmt.Sprintf("http://localhost:%s/organizations", base.GetEnv("ORGANIZATION_PORT", "7072"))
	resp, err := grequests.Get(organizationsUrl, nil)
	if err != nil {
		return nil, err
	}
	var organizations organizationDto.GetOrganizationsResponse
	resp.JSON(&organizations)
	return &organizations, nil
}

func GetOrganizationById(id int) (*organizationDto.OrganizationByIdResponse, error) {
	organizationUrl := fmt.Sprintf("http://localhost:%s/organizations/%d", base.GetEnv("ORGANIZATION_PORT", "7072"), id)
	resp, err := grequests.Get(organizationUrl, &grequests.RequestOptions{
		JSON: organizationDto.OrganizationByIdRequest{
			Id: id,
		}})
	if err != nil {
		return nil, err
	}
	var organization organizationDto.OrganizationByIdResponse
	resp.JSON(&organization)
	return &organization, nil
}
