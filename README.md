Build with:

go build extract.go

Run:

./extract /home/user/.geth/keystore/YOUR_KEY_FILE

Enter your password, and it will show the address associated to the key file, the two X, Y curve points in the secp256k1 public key, and the private key number (which is just a bigint)

The password is shown as it is typed, because this is not a tool written for security, just to extract keys in plain format.
