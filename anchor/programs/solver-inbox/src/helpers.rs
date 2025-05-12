use anchor_lang::prelude::Pubkey;
use sha2::Digest;

// orderID returns the order ID (32 byte array Pubkey),
//  by hashing the account (pubkey) and nonce.
pub fn order_id(account: Pubkey, nonce: u64) -> Pubkey {
    let mut hasher = sha2::Sha256::new();
    hasher.update(account.to_bytes());
    hasher.update(nonce.to_le_bytes());

    Pubkey::new_from_array(hasher.finalize().into())
}
