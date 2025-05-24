# SFTPGo SecretBox Decrypt

A simple script to decrypt secrets encrypted with SFTPGo's secretbox encryption.

## Setup

Since this script uses SFTPGo's internal packages, it needs to be run from within the SFTPGo source tree.

1. Clone the SFTPGo repository:
   ```bash
   git clone https://github.com/drakkan/sftpgo.git
   cd sftpgo
   ```

2. Copy the decrypt script into the SFTPGo directory:
   ```bash
   cp /path/to/this/repo/local_run.go ./decrypt.go
   ```

3. Build and run the script:
   ```bash
   go run decrypt.go <password_file> <contents_file> <additional_data>
   ```

## Usage

The script takes three arguments:
- `password_file`: Path to a file containing the decryption password
- `contents_file`: Path to a file containing the encrypted secretbox data
- `additional_data`: Additional data string used during encryption

Example:
```bash
go run decrypt.go password.txt encrypted_data.txt "some_folder_name"
```

## What it does

The script:
1. Reads the password from the specified file
2. Reads the encrypted data from the specified file
3. Uses SFTPGo's internal KMS functionality to decrypt the secretbox
4. Outputs the decrypted content

The `additional_data` parameter must match the value used when the secret was originally encrypted.