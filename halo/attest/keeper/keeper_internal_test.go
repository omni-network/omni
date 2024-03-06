package keeper

// AttestTable returns the attestations ORM table.
func (k *Keeper) AttestTable() AttestationTable {
	return k.attTable
}

// SignatureTable returns the attestations ORM table.
func (k *Keeper) SignatureTable() SignatureTable {
	return k.sigTable
}
