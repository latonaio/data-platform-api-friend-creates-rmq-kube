package dpfm_api_processing_formatter

type GeneralUpdates struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	Friend						int		`json:"Friend"`
	CommunityRank				int		`json:"CommunityRank"`
	FriendIsBlocked				bool	`json:"FriendIsBlocked"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
}
