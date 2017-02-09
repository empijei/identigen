for os in darwin freebsd
do 
	for arch in amd64 386
	do 
		mkdir -p "bin/$os/$arch/"
		GOOS=$os GOARCH=$arch go build -o "bin/$os/$arch/identigen"
	done
done

for arch in amd64 386
do 
	mkdir -p "bin/windows/$arch/"
	GOOS=windows GOARCH=$arch go build -o "bin/windows/$arch/identigen.exe"
done
