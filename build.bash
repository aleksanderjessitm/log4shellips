# We are going to build for 64 bit systems running macOS, Linux, and Windows
archs=(amd64 arm64)
systems=(darwin linux windows)

for arch in ${archs[@]}
do
    for system in ${systems[@]} 
    do
        if [ $system == 'windows' ]
        then
            env GOOS=${system} GOARCH=${arch} go build -o log4shellips_windows_${arch}.exe    
        else
            env GOOS=${system} GOARCH=${arch} go build -o log4shellips_${system}_${arch}
        fi

    done
done