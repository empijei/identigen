grep -Po "\"\K[^\-0-9\",]*" italianPIva.go | 
while read line
do 
	echo $line
done |
sort -u | 
sed 's/ /%20/g' | 
while read line
do                         
	curl "https://maps.googleapis.com/maps/api/geocode/json?address=$line&region=it&language=it"
done | 
grep -E '("location"|formatted)' -A3 | 
grep -E "(address|lat|lng)" | 
grep -Po ": \K.*" | 
grep -Eo "([0-9]{1,2}\..*|\"[^,]+)" | 
grep -Po '("[0-9]* ?\K.*|[0-9]{1,2}\..*)' | 
tr '\n' ' ' | 
sed -re 's/([0-9] )([A-Z])/\1\n\2/g' -e 's/[A-Z][A-Z] //g'
