# Run bash application
- Change array of dirs (`dirs=("dir1" "dir2" "dir3")`).
- Run `./count.sh`

# Run go application
- Change array of dirs (`dirs := []string{"dir1", "dir2", "dir3"}`).
- Build application 
    ```bash
    docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app golang:1.13 go build -v
    ```
- Run `./app`
