/*
 * Copyright (c) 2018.
 */

package dto

type ResponseServices struct {
	ResponseCode, ResponseDesc, Data string
}

type TabunganRequest struct {
	Amount, ChannelId, Cif, ClientId, Flag, GramTransaksi, IbuKandung,
	IdKelurahan, Jalan, JenisKelamin, JenisTransaksi, Kewarganegaraan,
	KodeBankPembayar, KodeBankTujuan, KodeCabang, NamaBankTujuan, NamaNasabah,
	NilaiTransaksi, NoHp, NoHpAgent, NoIdentitas, Norek, NorekBankTujuan,
	NorekWallet, NorekWalletTujuan,	PaymentMethod, ReffBiller, ReffSwitching,
	StatusKawin, TanggalExpiredId, TanggalLahir, TempatLahir, TipeIdentitas,
	TipeTransaksi, Token, WalletId string
}
