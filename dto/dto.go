package dto

type Dto struct {
	AgentId string `json:"agentId,omitempty"`
	Amount string `json:"amount,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	Cif string `json:"cif,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Data string `json:"data,omitempty"`
	Flag string `json:"flag,omitempty"`
	GramTransaksi string `json:"gramtransaksi,omitempty"`
	IbuKandung string `json:"ibuKandung,omitempty"`
	IdKelurahan string `json:"idKelurahan,omitempty"`
	Jalan string `json:"jalan,omitempty"`
	JenisKelamin string `json:"jenisKelamin,omitempty"`
	JenisTransaksi string `json:"jenisTransaksi,omitempty"`
	Kewarganegaraan string `json:"kewarganegaraan,omitempty"`
	KodeBankPembayar string `json:"kodeBankPembayar,omitempty"`
	KodeBankTujuan string `json:"kodeBankTujuan,omitempty"`
	KodeCabang string `json:"kodeCabang,omitempty"`
	NamaBankTujuan string `json:"namaBankTujuan,omitempty"`
	NamaNasabah string `json:"namaNasabah,omitempty"`
	NilaiTransaksi string `json:"nilaiTransaksi,omitempty"`
	NoHp string `json:"noHp,omitempty"`
	NoHpAgent string `json:"noHpAgent,omitempty"`
	NoIdentitas string `json:"noIdentitas,omitempty"`
	Norek string `json:"norek,omitempty"`
	NorekBankTujuan string `json:"norekBankTujuan,omitempty"`
	NorekWallet string `json:"norekWallet,omitempty"`
	NorekWalletTujuan string `json:"norekWalletTujuan,omitempty"`
	PaymentMethod string `json:"paymentMethod,omitempty"`
	ReffBiller string `json:"reffBiller,omitempty"`
	ReffSwitching string `json:"reffSwitching,omitempty"`
	RequestType string `json:"requestType,omitempty"`
	ResponseCode string `json:"responseCode,omitempty"`
	ResponseDesc string `json:"responseDesc,omitempty"`
	StatusKawin string `json:"statusKawin,omitempty"`
	TanggalExpiredId string `json:"tanggalExpiredId,omitempty"`
	TanggalLahir string `json:"tanggalLahir,omitempty"`
	TempatLahir string `json:"tempatLahir,omitempty"`
	TipeIdentitas string `json:"tipeIdentitas,omitempty"`
	TipeTransaksi string `json:"tipeTransaksi,omitempty"`
	Token string `json:"token,omitempty"`
	WalletId string `json:"walletId,omitempty"`
}
