/*******************************************************************************
 * IBM Confidential
 * OCO Source Materials
 * IBM Cloud Container Service, 5737-D43
 * (C) Copyright IBM Corp. 2018 All Rights Reserved.
 * The source code for this program is not  published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with
 * the U.S. Copyright Office.
 ******************************************************************************/

package backend

//Session interface describes methods requiring implementation
type Session interface {
	GetAccountService() AccountService
	GetBillingItemService() BillingItemService
	GetBillingOrderService() BillingOrderService
	GetNetworkStorageIscsiService() NetworkStorageIscsiService
	GetProductOrderService() ProductOrderService
	GetProductPackageService() ProductPackageService
	GetNetworkStorageService() NetworkStorageService
	GetNetworkSubnetService() NetworkSubnetService
	GetNetworkSubnetIpAddressService() NetworkSubnetIpAddressService
	GetResourceMetadataService() ResourceMetadataService
	GetLocationService() LocationService
}
