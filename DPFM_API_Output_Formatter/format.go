package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-friend-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-friend-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-friend-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToGeneralCreates(subfuncSDC *sub_func_complementer.SDC) (*General, error) {
	data := subfuncSDC.Message.General

	general, err := TypeConverter[*General](data)
	if err != nil {
		return nil, err
	}

	return general, nil
}

func ConvertToGeneralUpdates(generalData dpfm_api_input_reader.General) (*General, error) {
	data := generalData

	general, err := TypeConverter[*General](data)
	if err != nil {
		return nil, err
	}

	return general, nil
}

func ConvertToGeneral(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.General = &sub_func_complementer.General{
		BusinessPartner:			input.General.BusinessPartner,
		Friend:						input.General.Friend,
		BPBusinessPartnerType:		input.General.BPBusinessPartnerType,
		FriendBusinessPartnerType:	input.General.FriendBusinessPartnerType,
		CommunityRank:				input.General.CommunityRank,
		FriendIsBlocked:			input.General.FriendIsBlocked,
		CreationDate:				input.General.CreationDate,
		CreationTime:				input.General.CreationTime,
		LastChangeDate:				input.General.LastChangeDate,
		LastChangeTime:				input.General.LastChangeTime,
		IsMarkedForDeletion:		input.General.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
