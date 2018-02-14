package dto

type Dto struct {
	AgentId string `json:"agentId"`
	Amount string `json:"amount"`
	ChannelId string `json:"channelId"`
	Cif string `json:"cif"`
	ClientId string `json:"clientId"`
	Data string `json:"data"`
	Flag string `json:"flag"`
	GramTransaksi string `json:"gramtransaksi"`
	IbuKandung string `json:"ibuKandung"`
	IdKelurahan string `json:"idKelurahan"`
	Jalan string `json:"jalan"`
	JenisKelamin string `json:"jenisKelamin"`
	JenisTransaksi string `json:"jenisTransaksi"`
	Kewarganegaraan string `json:"kewarganegaraan"`
	KodeBankPembayar string `json:"kodeBankPembayar"`
	KodeBankTujuan string `json:"kodeBankTujuan"`
	KodeCabang string `json:"kodeCabang"`
	NamaBankTujuan string `json:"namaBankTujuan"`
	NamaNasabah string `json:"namaNasabah"`
	NilaiTransaksi string `json:"nilaiTransaksi"`
	NoHp string `json:"noHp"`
	NoHpAgent string `json:"noHpAgent"`
	NoIdentitas string `json:"noIdentitas"`
	Norek string `json:"norek"`
	NorekBankTujuan string `json:"norekBankTujuan"`
	NorekWallet string `json:"norekWallet"`
	NorekWalletTujuan string `json:"norekWalletTujuan"`
	PaymentMethod string `json:"paymentMethod"`
	ReffBiller string `json:"reffBiller"`
	ReffSwitching string `json:"reffSwitching"`
	ResponseCode string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
	StatusKawin string `json:"statusKawin"`
	TanggalExpiredId string `json:"tanggalExpiredId"`
	TanggalLahir string `json:"tanggalLahir"`
	TempatLahir string `json:"tempatLahir"`
	TipeIdentitas string `json:"tipeIdentitas"`
	TipeTransaksi string `json:"tipeTransaksi"`
	Token string `json:"token"`
	WalletId string `json:"walletId"`
}
