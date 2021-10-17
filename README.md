# Instructions
The file needs to have the same name and extension defined at ```config.yaml```.

You should set a **KEY** and save it safely, otherwise it will be hard to recover the original file.

**PS**:   Under no circumstances, you should type your key at the ```config.yaml``` file.

## Usage
1. Create a file at the root of the project. **Ex**: ```cmd.txt```
2. In the file, add the content that will be cryptographed.
3. Execute the following command:

        KEY=mystrongerkey go run main.go encrypt

The file ```cmd.txt```  will be removed and a new file will be created with the same name, but with a different extension. Ex: ```cmd.txt.d```

- To reverse the cryptography:

        KEY=mystrongerkey go run main.go decrypt
    
- The ```cmd.txt```  file will be created and the encrypted file (```cmd.txt.d```) will be removed.



## References

- https://pkg.go.dev/crypto
- https://pkg.go.dev/crypto/aes@go1.17.2
- https://pkg.go.dev/crypto/cipher@go1.17.2
- https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
