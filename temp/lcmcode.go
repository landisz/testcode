package main

import "context"

// func getNSSTListfromGnst(ctx context.Context, gnst model.Gnst) ([]model.Nsst, error) {
// 	const op errors.Op = "lcm.simple.getNSSTListfromGnst"

// 	resource, err := listResource(ctx, "nsst")
// 	if err != nil {
// 		return nil, err
// 	}
// 	res := resource.(*model.NsstList)
// 	nsstList := res.Nsst

// 	// List of nsst to instantiate
// 	nsstListToInstantiate := []model.Nsst{}

// 	//
// 	for _, v1 := range gnst.GnsstList {
// 		// select a NSST for each gnsst based on category
// 		for _, v2 := range nsstList {
// 			if v2.Category == v1.Category {
// 				nsstListToInstantiate = append(nsstListToInstantiate, *v2)
// 				break
// 			}
// 		}
// 	}
// 	if len(nsstListToInstantiate) < len(gnst.GnsstList) {
// 		errmsg := "Error in NSST selection; Could not find suitable nsst(s) for one or more gnssts"
// 		return nil, errors.E(op, errors.ErrorUnknown, errmsg)
// 	}
// 	return nsstListToInstantiate, nil
// }

func createNssProfileFromServiceProfile(ctx context.Context, nssiId string, serviceProfile model.ServiceProfile) model.NssProfile {

log.Ctx(ctx).Debug().Msg("Create NssProfile from ServiceProfile")

NssProfile := model.NssProfile{}
// create SlicePofile
NssProfile.ID = uuid.Must(uuid.NewV4(), nil).String()
NssProfile.NssiID = nssiId
// sliceProfile.CoverageAreaTAList = getTAfromCoverageArea()
NssProfile.Latency = (serviceProfile.Latency) // define unit of performance metrices throughout NOC
NssProfile.MaxNumberofUEs = (serviceProfile.MaxNumberofUEs)
// sliceProfile.Perfreq.
NssProfile.ResourceSharingLevel = serviceProfile.ResourceSharingLevel
// sliceProfile.PlmnIDList = serviceProfile.PlmnIDList
// sliceProfile.SnssaiList = serviceProfile.SnssaiList
NssProfile.UeMobilityLevel = serviceProfile.UEMobilityLevel

return NssProfile
}

// func assignNssicapabilityBasedOnserviceProfile(serviceProfile model.ServiceProfile) model.NssiCapability {
// 	nssiCapability := model.NssiCapability{}
// 	nssiCapability.ActivityFactor = serviceProfile.ActivityFactor
// 	nssiCapability.Availability = serviceProfile.Availability
// 	nssiCapability.CoverageArea = serviceProfile.CoverageArea
// 	nssiCapability.Latency = serviceProfile.Latency
// 	nssiCapability.MaxNumberofUEs = serviceProfile.MaxNumberofUEs
// 	nssiCapability.Sst = serviceProfile.Sst
// 	nssiCapability.PlmnIDList = serviceProfile.PlmnIDList
// 	nssiCapability.SnssaiList = serviceProfile.SnssaiList
// 	nssiCapability.ResourceSharingLevel = serviceProfile.ResourceSharingLevel

// 	return nssiCapability
// }

//create NSSI req for NSSMF


//create TransportSlice req for NSSMF
func (l *nssiLcmLogic) instantiateTransportSlice(ctx context.Context, nssi *model.Nssi) error {
const op errors.Op = "lcm.simple.instantiateTransportSlice"

// Step 1: create TransportSlice profile
l.log.Debug().Msg("Create TransportSliceProfile from ServiceProfile/NssiCapability")
transportSliceProfile := createTransportSliceProfileFromServiceProfile(nssi.ID, nssi.NssiCapability)

// update transport slice info into nssi
nssi.TransportSliceInfo.TransportSliceProfile = transportSliceProfile
l.log.Debug().Msgf("Update Nssi TransportSliceInfo with ProfileId %s ", transportSliceProfile.ID)

// update nssi into database
l.log.Debug().Msgf("Update nssi into database")
l.database.UpdateNssi(ctx, nssi)

l.log.Debug().Msgf("Call allocateTransportSlice API to allocate a transport slice to NSI")
response, err := l.allocateTSI(ctx, transportSliceProfile)

if err != nil {
// publish NssiStateUpdate with err
return err
}

if response.StatusCode != http.StatusAccepted {
errmsg := "Transport slice instantiation request not accepted by the NSSMF; StatusCode from RI" + strconv.FormatInt(int64(response.StatusCode), 10)
return errors.E(op, errors.ErrorUnknown, errmsg)
}
return nil
}

func createTransportSliceProfileFromServiceProfile(nssiId string, nssiCapbility model.NssiCapability) model.TransportSliceProfile {
transportSliceProfile := model.TransportSliceProfile{}
transportSliceProfile.ID = uuid.Must(uuid.NewV4(), nil).String()
transportSliceProfile.NssiID = nssiId
transportSliceProfile.Availability = nssiCapbility.Availability
// transportSliceProfile.DLThptPerSlice = serviceProfile.DLThptPerSlice
// transportSliceProfile.DelayTolerance = serviceProfile.DelayTolerance
transportSliceProfile.Jitter = nssiCapbility.Jitter
// transportSliceProfile.KPIMonitoring = serviceProfile.KPIMonitoring
transportSliceProfile.Latency = nssiCapbility.Latency
transportSliceProfile.ResourceSharingLevel = nssiCapbility.ResourceSharingLevel
transportSliceProfile.Sst = nssiCapbility.Sst
transportSliceProfile.SurvivalTime = nssiCapbility.SurvivalTime
transportSliceProfile.UESpeed = nssiCapbility.UESpeed // Mobility Level
// transportSliceProfile.ULThptPerSlice = serviceProfile.ULThptPerSlice
// 	transportSliceProfile.TransportSliceConnection =

return transportSliceProfile

}

// nssmf response simulation
func (l *nssiLcmLogic) NssiStateUpdateAfterAllocation(ctx context.Context, nssiStateUpdate model.NssiStateUpdate) error {
const op errors.Op = "lcm.simple.NssiStateUpdateAfterAllocation"

// check for err in NSSI instatiation
// if nssiStateUpdate.Error != "" {
// 	l.log.Error().Err(err).Msgf("Transport slice allocation is not successful; Error: %s", nssiStateUpdate.Error)
// 	l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionAllocate, nil, err)
// 	return err
// }

mutex.Lock()
l.log.Debug().Msg("Update NSI using NssiStateUpdate After Allocation")

// Get NssiId from nssiStateUPdate
nssiId := nssiStateUpdate.NssiID
l.log.Debug().Msgf("Update NSSI State with NssiId %s NssiId %s nssProfileID %s", nssiStateUpdate.NssiID, nssiStateUpdate.NssiID, nssiStateUpdate.NssProfileID)

// get nssi from nsmf database and update nssi it
l.log.Debug().Msg("Retrive NSI from nsmf database")
nssi, err := l.database.FindNssiByID(ctx, nssiId)
if err != nil {
fmt.Println("Error in Retriving NSI from NSMF database")
fmt.Println(err)
mutex.Unlock()
return err
}

AllNssiInstantiated := true
l.log.Debug().Msgf("Length of NssIInfoList %s", strconv.Itoa(len(nssi.NssiInfoList)))

// update nssiId in nssi
for i, v := range nssi.NssiInfoList {
if v.NssProfile.ID == nssiStateUpdate.NssProfileID {
nssi.NssiInfoList[i].NssiID = nssiStateUpdate.NssiID
}
}

// update nssi in nsmf database
l.database.UpdateNssi(ctx, nssi)

if len(nssi.NssiInfoList) < len(nssi.Nst.NsstList) {
l.log.Debug().Msg("All NSSI(s) are not instantiated")
AllNssiInstantiated = false
} else {
for _, v := range nssi.NssiInfoList {
l.log.Debug().Msgf("NssProfileId: %s NssiId: %s", v.NssProfile.ID, v.NssiID)
if v.NssiID == "" {
l.log.Debug().Msg("All NSSI(s) are not instantiated")
AllNssiInstantiated = false
break
}
}
}

if AllNssiInstantiated {
// All NSSI is instantiated;
l.log.Debug().Msg("All NSSI(s) are initiated; Now instatiate Transport Slice Instance")
err := l.instantiateTransportSlice(ctx, nssi)
if err != nil {
err = errors.E(op, err)
l.log.Error().Err(err).Msgf("Transport slice allocation is not successful; Error: %s", err)
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionAllocate, nil, err)
mutex.Unlock()
return err
}
}
mutex.Unlock()
return nil
}

func (l *nssiLcmLogic) NssiStateUpdateAfterDeallocation(ctx context.Context, nssiStateUpdate model.NssiStateUpdate) error {
l.log.Debug().Msgf("LCM logic receive a NSSIStateUpdate for nssiId %s after deallocation", nssiStateUpdate.NssiID)
const op errors.Op = "nssiLcmLogic.NssiStateUpdateAfterDeallocation"

// delete nssProfile from the RI database
l.log.Debug().Msg("Delete nssProfile from the RI database")
err := deleteResourceByID(ctx, "nssProfile", nssiStateUpdate.NssProfileID)

if err != nil {
err = errors.E(op, err)
log.Ctx(ctx).Error().Err(err).Msgf("Error in deleting NssProfilefrom from the RI database; Error: %s", err)
return err
}

// get and update nssi form nsmf database
mutex.Lock()
l.log.Debug().Msg("Update NssiInfoList of the nssi in the NSMF database")
l.log.Debug().Msg("Retrive NSI from the NSMF database")
nssi, err := l.database.FindNssiByID(ctx, nssiStateUpdate.NssiID)
if err != nil {
err = errors.E(op, err)
l.log.Error().Err(err).Msgf("Error in retriving Nssi from the NSMF database; Error: %s", err)
mutex.Unlock()
return err
}

// update NssiInfoList
index := 0
for i, v := range nssi.NssiInfoList {
if v.NssiID == nssiStateUpdate.NssiID {
index = i
}
}
nssiInfoList := nssi.NssiInfoList
nssiInfoList = append(nssiInfoList[:index], nssiInfoList[(index+1):]...)

// update nssi
nssi.NssiInfoList = nssiInfoList
err = l.database.UpdateNssi(ctx, nssi)
if err != nil {
err = errors.E(op, err)
l.log.Error().Err(err).Msgf("Error in retriving Nssi from the NSMF database; Error: %s", err)
mutex.Unlock()
return err
}
l.log.Debug().Msgf("NSSI with Id %s is deallocated successfully", nssiStateUpdate.NssiID)
mutex.Unlock()

// check all NSSI deallocated successfully
if len(nssiInfoList) == 0 {
// All the nssi(s) are deallocated;
l.log.Debug().Msg("All the nssi(s) are deallocated successfully; Now deallocate transport slice")
// deallocate TSI and then, delete NSI from the NSMF database
ctx := context.Background()

response, err := l.deallocateTSI(ctx, nssi.TransportSliceInfo.ID, nssi.TransportSliceInfo.TransportSliceProfile)

if err != nil {
err = errors.E(op, err)
l.log.Error().Err(err).Msgf("Error received from NSSMF for NSSI deallocation; Error: %s", err)
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionDeallocate, nil, err)
return err
} else if response.StatusCode != http.StatusAccepted {
errmsg := "Error creating resource in RI from NSMF; StatusCode from RI" + strconv.FormatInt(int64(response.StatusCode), 10)
l.log.Debug().Msg(errmsg)
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionDeallocate, nil, err)
return errors.E(op, errors.ErrorUnknown, errmsg)
}
}
return nil
}

func (l *nssiLcmLogic) TsiStateUpdateAfterDeallocation(ctx context.Context, tsiStateUpdate model.TsiStateUpdate) error {
const op errors.Op = "nssiLcmLogic.simple.TsiStateUpdateAfterDeallocation"
l.log.Debug().Msg("Delete/Update NSI using TsiStateUpdate After Deallocation")

// Get nssi from database
l.log.Debug().Msgf("Retrive NSI with ID %s from the NSMF database", tsiStateUpdate.NssiID)
nssi, err := l.database.FindNssiByID(ctx, tsiStateUpdate.NssiID)
if err != nil {
err = errors.E(op, err)
l.log.Error().Msgf("Error in retriving nssi from the NSMF database; Error: %s", err)
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionDeallocate, nil, err)
return err
}

// Check for error in TSI deallocation
if tsiStateUpdate.Error != "" {
errmsg := "Error in TSI deallocation; Error: " + tsiStateUpdate.Error
l.log.Debug().Msg(errmsg)
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionDeallocate, nil, errors.E(op, errors.ErrorUnknown, errmsg))
return errors.E(op, errors.ErrorUnknown, errmsg)
}

// Delete NSI from the NSMF database
l.log.Debug().Msgf("Delete NSI with ID %s from the NSMF database", tsiStateUpdate.NssiID)
err = l.database.DeleteNssiByID(ctx, tsiStateUpdate.NssiID)
if err != nil {
err = errors.E(op, err)
l.log.Error().Msgf("Error in deleting nssi from the NSMF database; Error: %s", err)
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionDeallocate, nil, err)
return err
}

// Delete NSI from the RI database
l.log.Debug().Msgf("Delete NSI with ID %s from the RI database", tsiStateUpdate.NssiID)
err = deleteResourceByID(ctx, "nssi", nssi.ID)
if err != nil {
err = errors.E(op, err)
log.Ctx(ctx).Error().Err(err).Msgf("Error in deleting Nssi from the RI database; Error: %s", err)
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionDeallocate, nil, err)
return err
}

// publish the delloation msg to csmf
l.log.Debug().Msgf("Publish NssiStateUpdate after NSI deallocation")
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionDeallocate, nssi, nil)
return nil
}

func (l *nssiLcmLogic) TsiStateUpdateAfterAllocation(ctx context.Context, tsiStateUpdate model.TsiStateUpdate) error {
l.log.Debug().Msg("Update NSI using TsiStateUpdate After allocation")
const op errors.Op = "lcm.simple.TsiStateUpdateAfterAllocation"

nssiId := tsiStateUpdate.NssiID

// Retrive NSI from the nsmf database
l.log.Debug().Msg("Retrive NSI from nsmf database")
nssi, err := l.database.FindNssiByID(ctx, nssiId)
if err != nil {
err = errors.E(op, err)
l.log.Debug().Msg("Error in Retriving NSI from NSMF database")
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionAllocate, nil, err)
return err
}

// update tsiId into nssi
l.log.Debug().Msg("Update tsiId into nssi")
nssi.TransportSliceInfo.ID = tsiStateUpdate.TsiID

// set state of NSI to active
l.log.Debug().Msg("Update NSI state to active")
nssi.SliceState = model.SliceStateActive

// update nssi into the NSMF database
err = l.database.UpdateNssi(ctx, nssi)
if err != nil {
err = errors.E(op, err)
fmt.Println("Error in updating NSI from NSMF database")
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionAllocate, nil, err)
return err
}

// update NSI in RI database
l.log.Debug().Msg("Store NSI into RI database")
err = createResource(ctx, nssi, "nssi")
if err != nil {
err = errors.E(op, err)
fmt.Println("Error in adding nssi into the RI database")
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionAllocate, nil, err)
return err
}

// Publish NssiStateUpdate to csmf queue using NOC MB
l.log.Debug().Msg("Publish NssiStateUpdate to csmf queue using NOC MB")
l.PublishNssiStateUpdate(ctx, nssi.CommunicationServiceIDList[0], model.ActionAllocate, nssi, nil)
return nil
}
//
func (l *lcmLogic) allocatedNssi(ctx context.Context, allocateNssi model.AllocateNssi) (*http.Response, error) {
const op errors.Op = "nssi.allocatedNssi"

response := &http.Response{}
response.StatusCode = http.StatusAccepted
nssiId := uuid.Must(uuid.NewV4(), nil).String()
log.Ctx(ctx).Debug().Msg("Publish NssiStateUpdate to NOC MQ")
l.PublishNssiStateUpdate(ctx, nssiId, model.ActionAllocate, allocateNssi.NssProfile, nil)
return response, nil
}

func (l *lcmLogic) deallocateNSSI(ctx context.Context, nssiId string, nssProfile model.NssProfile) (*http.Response, error) {
const op errors.Op = "nssi.allocatedNssi"

response := &http.Response{}
response.StatusCode = http.StatusAccepted
log.Ctx(ctx).Debug().Msg("Publish NssiStateUpdate to NOC MQ")
l.PublishNssiStateUpdate(ctx, nssiId, model.ActionDeallocate, nssProfile, nil)
return response, nil
}

