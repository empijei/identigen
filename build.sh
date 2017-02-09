for os in darwin freebsd windows
do 
	for arch in amd64 386
	do 
		mkdir -p "bin/$os/$arch/"
		GOOS=$os GOARCH=$arch go build -o "bin/$os/$arch/identigen"
	done
done
