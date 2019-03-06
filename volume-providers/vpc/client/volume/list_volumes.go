/*******************************************************************************
 * IBM Confidential
 * OCO Source Materials
 * IBM Cloud Container Service, 5737-D43
 * (C) Copyright IBM Corp. 2018 All Rights Reserved.
 * The source code for this program is not  published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with
 * the U.S. Copyright Office.
 ******************************************************************************/

package volume

import (
	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/client"
	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/client/models"
	"strconv"
)

// ListVolumes GETs /volumes
func (vs *VolumeService) ListVolumes(limit int, filters *models.ListVolumeFilters) (*models.VolumeList, error) {
	operation := &client.Operation{
		Name:        "ListVolumes",
		Method:      "GET",
		PathPattern: volumesPath,
	}

	var volumes models.VolumeList
	var apiErr models.Error

	req := vs.client.NewRequest(operation).JSONSuccess(&volumes).JSONError(&apiErr)

	if limit > 0 {
		req.AddQueryValue("limit", strconv.Itoa(limit))
	}

	if filters != nil {
		if filters.ResourceGroupID != "" {
			req.AddQueryValue("resource_group.id", filters.ResourceGroupID)
		}
		if filters.Tag != "" {
			req.AddQueryValue("tag", filters.Tag)
		}
		if filters.ZoneName != "" {
			req.AddQueryValue("zone.name", filters.ZoneName)
		}

	}

	_, err := req.Invoke()
	if err != nil {
		return nil, err
	}

	return &volumes, nil
}
